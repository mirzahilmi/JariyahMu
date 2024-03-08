package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryItf interface {
	CreateUser(ctx context.Context, user model.StoreUser) error
	GetUserByParam(ctx context.Context, param string, value string) (model.StoreUser, error)
	UpdateUserPassword(ctx context.Context, id, hashedPassword string) error
	CreateResetAttempt(ctx context.Context, attempt model.StoreResetAttempt) error
	DeleteOldResetAttempt(ctx context.Context, id string) error
	GetResetAttemptExpiration(ctx context.Context, id, token string) (time.Time, error)
	UpdateResetAttemptStatus(ctx context.Context, id string) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepositoryItf {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.StoreUser) error {
	query, args, err := sqlx.Named(queryCreateUser, user)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByParam(ctx context.Context, param string, value string) (model.StoreUser, error) {
	var user model.StoreUser
	if err := r.db.GetContext(ctx, &user, fmt.Sprintf(queryGetUserByParam, param), value); err != nil {
		return model.StoreUser{}, err
	}

	return user, nil
}
func (r *UserRepository) UpdateUserPassword(ctx context.Context, id, hashedPassword string) error {
	query, args, err := sqlx.Named(queryUpdatePassword, map[string]string{"ID": id, "Hashed": hashedPassword})
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) CreateResetAttempt(ctx context.Context, attempt model.StoreResetAttempt) error {
	query, args, err := sqlx.Named(queryCreateAttempt, attempt)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteOldResetAttempt(ctx context.Context, id string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, queryDeleteOldAttempt, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetResetAttemptExpiration(ctx context.Context, id, token string) (time.Time, error) {
	var result = map[string]time.Time{"Expiration": {}}
	if err := r.db.GetContext(ctx, &result, queryGetAttemptExpiration, id, token); err != nil {
		return time.Time{}, err
	}

	return result["Expiration"], nil
}

func (r *UserRepository) UpdateResetAttemptStatus(ctx context.Context, id string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, queryUpdateAttemptStatus, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
