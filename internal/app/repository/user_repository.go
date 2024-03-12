package repository

import (
	"context"
	"fmt"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryItf interface {
	Create(ctx context.Context, user model.UserResource) error
	GetByParam(ctx context.Context, param string, value string) (model.UserResource, error)
	UpdatePassword(ctx context.Context, id, hashedPassword string) error
	UpdateActiveStatus(ctx context.Context, id string) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepositoryItf {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user model.UserResource) error {
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

func (r *UserRepository) GetByParam(ctx context.Context, param string, value string) (model.UserResource, error) {
	var user model.UserResource
	if err := r.db.GetContext(ctx, &user, fmt.Sprintf(queryGetUserByParam, param), value); err != nil {
		return model.UserResource{}, err
	}

	return user, nil
}
func (r *UserRepository) UpdatePassword(ctx context.Context, id, hashedPassword string) error {
	query, args, err := sqlx.Named(queryUpdateUserPassword, map[string]string{"ID": id, "HashedPassword": hashedPassword})
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

func (r *UserRepository) UpdateActiveStatus(ctx context.Context, id string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, queryUpdateUserStatus, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
