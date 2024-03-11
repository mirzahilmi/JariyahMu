package repository

import (
	"context"
	"time"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/jmoiron/sqlx"
)

type ResetAttemptRepositoryItf interface {
	CreateResetAttempt(ctx context.Context, attempt model.StoreResetAttempt) error
	DeleteOldResetAttempt(ctx context.Context, userID string) error
	GetResetAttemptExpiration(ctx context.Context, id, token string) (time.Time, error)
	UpdateResetAttemptStatus(ctx context.Context, id string) error
}

type ResetAttemptRepository struct {
	db *sqlx.DB
}

func NewResetAttemptRepository(db *sqlx.DB) ResetAttemptRepositoryItf {
	return &ResetAttemptRepository{db}
}

func (r *ResetAttemptRepository) CreateResetAttempt(ctx context.Context, attempt model.StoreResetAttempt) error {
	query, args, err := sqlx.Named(queryCreateResetAttempt, attempt)
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

func (r *ResetAttemptRepository) DeleteOldResetAttempt(ctx context.Context, userID string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, queryDeleteOldResetAttempt, userID); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *ResetAttemptRepository) GetResetAttemptExpiration(ctx context.Context, id, token string) (time.Time, error) {
	var result = map[string]time.Time{"Expiration": {}}
	if err := r.db.GetContext(ctx, &result, queryGetResetAttemptExpiration, id, token); err != nil {
		return time.Time{}, err
	}

	return result["Expiration"], nil
}

func (r *ResetAttemptRepository) UpdateResetAttemptStatus(ctx context.Context, id string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, queryUpdateResetAttemptStatus, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
