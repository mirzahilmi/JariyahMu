package repository

import (
	"context"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserVerificationRepositoryItf interface {
	Create(ctx context.Context, attempt model.UserVerificationResource) error
	GetByIDAndToken(ctx context.Context, id, token string) (model.UserVerificationResource, error)
	UpdateSucceedStatus(ctx context.Context, id string) error
}

type UserVerificationRepository struct {
	db *sqlx.DB
}

func NewUserVerificationRepository(db *sqlx.DB) UserVerificationRepositoryItf {
	return &UserVerificationRepository{db}
}

func (r *UserVerificationRepository) Create(ctx context.Context, attempt model.UserVerificationResource) error {
	query, args, err := sqlx.Named(queryCreateUserVerification, attempt)
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

func (r *UserVerificationRepository) GetByIDAndToken(ctx context.Context, id string, token string) (model.UserVerificationResource, error) {
	var attempt model.UserVerificationResource
	if err := r.db.GetContext(ctx, &attempt, queryGetUserVerificationByIDAndToken, id, token); err != nil {
		return model.UserVerificationResource{}, err
	}

	return attempt, nil
}

func (r *UserVerificationRepository) UpdateSucceedStatus(ctx context.Context, id string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, queryUpdateUserVerificationStatus, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
