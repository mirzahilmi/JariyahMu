package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/MirzaHilmi/JariyahMu/internal/app/repository"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/auth"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/email"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/response"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type UserUsecaseItf interface {
	RegisterUser(ctx context.Context, user model.CreateUserRequest) error
	LogUserIn(ctx context.Context, attempt model.UserLoginAttemptRequest) (string, error)
	VerifyUser(ctx context.Context, id, token string) error
	AttemptResetPassword(ctx context.Context, query model.QueryUserByEmailRequest) (string, error)
	ResetPassword(ctx context.Context, attempt model.AttemptResetPasswordRequest) error
}

type UserUsecase struct {
	userRepository             repository.UserRepositoryItf
	userVerificationRepository repository.UserVerificationRepositoryItf
	resetAttemptRepository     repository.ResetAttemptRepositoryItf
	mailer                     email.VerificationMailer
	pasetoInstance             *auth.Paseto
	viper                      *viper.Viper
}

func NewUserUsecase(
	userRepository repository.UserRepositoryItf,
	userVerificationRepository repository.UserVerificationRepositoryItf,
	resetAttemptRepository repository.ResetAttemptRepositoryItf,
	mailer email.VerificationMailer,
	pasetoInstance *auth.Paseto,
	viper *viper.Viper,
) UserUsecaseItf {
	return &UserUsecase{
		userRepository,
		userVerificationRepository,
		resetAttemptRepository,
		mailer,
		pasetoInstance,
		viper,
	}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, userRequest model.CreateUserRequest) error {
	userID, err := helper.ULID()
	if err != nil {
		return err
	}

	hashed, err := helper.BcryptHash(userRequest.Password)
	if err != nil {
		return err
	}

	user := model.UserResource{
		ID:             userID,
		FullName:       userRequest.FullName,
		Email:          userRequest.Email,
		HashedPassword: hashed,
	}
	mysqlErr := mysqlErrPool.Get().(*mysql.MySQLError)
	defer mysqlErrPool.Put(mysqlErr)
	if err := u.userRepository.Create(ctx, user); err != nil {
		switch {
		case errors.As(err, &mysqlErr) && mysqlErr.Number == 1062:
			return response.NewError(fiber.StatusConflict, ErrEmailExist)
		default:
			return err
		}
	}

	verificationID, err := helper.ULID()
	if err != nil {
		return err
	}

	attempt := model.UserVerificationResource{
		ID:     verificationID,
		UserID: userID,
		Token:  helper.RandString(32),
	}
	if err := u.userVerificationRepository.Create(ctx, attempt); err != nil {
		return err
	}

	emailProps := map[string]string{
		"Name": user.FullName,
		"URL":  fmt.Sprintf("%s/api/v1/auth/verify/%s?t=%s", u.viper.GetString("APP_HOST"), verificationID, attempt.Token),
		"Day":  time.Now().Weekday().String(),
	}

	if err := u.mailer.SendMail(user.Email, "Account Verification", email.TemplateURLVerification, emailProps); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) VerifyUser(ctx context.Context, id, token string) error {
	verificationAttempt, err := u.userVerificationRepository.GetByIDAndToken(ctx, id, token)
	if err != nil {
		return response.NewError(fiber.StatusBadRequest, ErrVerificationNotExist)
	}

	if err := u.userVerificationRepository.UpdateSucceedStatus(ctx, verificationAttempt.ID); err != nil {
		return err
	}

	if err := u.userRepository.UpdateActiveStatus(ctx, verificationAttempt.UserID); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) LogUserIn(ctx context.Context, attempt model.UserLoginAttemptRequest) (string, error) {
	user, err := u.userRepository.GetByParam(ctx, "Email", attempt.Email)
	if err != nil {
		return "", response.NewError(fiber.StatusNotFound, ErrUserNotExist)
	}

	if !user.Active {
		return "", response.NewError(fiber.StatusUnauthorized, ErrUserNotActive)
	}

	if err := helper.BcryptCompare(user.HashedPassword, attempt.Password); err != nil {
		return "", response.NewError(fiber.StatusUnauthorized, ErrWrongPassword)
	}

	token := paseto.NewToken()

	token.SetAudience("*")
	token.SetIssuer(u.viper.GetString("APP_HOST"))
	token.SetSubject(user.ID)
	token.SetExpiration(time.Now().Add(viper.GetDuration("PASETO_TTL") * time.Minute))
	token.SetNotBefore(time.Now())
	token.SetIssuedAt(time.Now())

	signed, err := u.pasetoInstance.Encode(token)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (u *UserUsecase) AttemptResetPassword(ctx context.Context, query model.QueryUserByEmailRequest) (string, error) {
	user, err := u.userRepository.GetByParam(ctx, "Email", query.Email)
	if err != nil {
		return "", err
	}

	if err := u.resetAttemptRepository.DeleteOld(ctx, user.ID); err != nil {
		return "", err
	}

	id, err := helper.ULID()
	if err != nil {
		return "", err
	}

	token, err := helper.RandNumber(6)
	if err != nil {
		return "", err
	}

	attempt := model.StoreResetAttempt{
		ID:         id,
		UserID:     user.ID,
		Token:      token,
		Expiration: time.Now().Add(10 * time.Minute),
	}

	if err := u.resetAttemptRepository.Create(ctx, attempt); err != nil {
		return "", err
	}

	emailProps := map[string]string{
		"Name": user.FullName,
		"Code": token,
		"Day":  time.Now().Weekday().String(),
	}

	if err := u.mailer.SendMail(user.Email, "Reset Password Attempt", email.TemplateCodeVerification, emailProps); err != nil {
		return "", err
	}

	return id, nil
}

func (u *UserUsecase) ResetPassword(ctx context.Context, attempt model.AttemptResetPasswordRequest) error {
	return nil
}
