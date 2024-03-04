package repository

import (
	"context"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryItf interface {
	Create(ctx context.Context, user model.StoreUser) error
	GetIDByEmail(ctx context.Context, email string) (string, error)
	UpdatePassword(ctx context.Context, id, hashedPassword string) error
	CreateResetAttempt(ctx context.Context, attempt model.StoreResetAttempt) error
	DeleteOldResetAttempt(ctx context.Context, id string) error
	GetResetAttemptID(ctx context.Context, id, token string) (string, error)
	UpdateResetAttemptStatus(ctx context.Context, id string) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepositoryItf {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user model.StoreUser) error {
	namedQuery, args, err := sqlx.Named(queryCreateUser, user)
	if err != nil {
		return err
	}
	query := sqlx.Rebind(sqlx.QUESTION, namedQuery)

	tx, err := r.db.Begin()
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

func (r *UserRepository) GetIDByEmail(ctx context.Context, email string) (string, error) {
	return "", nil
}
func (r *UserRepository) UpdatePassword(ctx context.Context, id, hashedPassword string) error {
	return nil
}
func (r *UserRepository) CreateResetAttempt(ctx context.Context, attempt model.StoreResetAttempt) error {
	return nil
}
func (r *UserRepository) DeleteOldResetAttempt(ctx context.Context, id string) error {
	return nil
}
func (r *UserRepository) GetResetAttemptID(ctx context.Context, id, token string) (string, error) {
	return "", nil
}
func (r *UserRepository) UpdateResetAttemptStatus(ctx context.Context, id string) error {
	return nil
}
