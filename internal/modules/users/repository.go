package users

import (
	"context"

	"github.com/boichique/avito-test-task/internal/apperrors"
	"github.com/boichique/avito-test-task/internal/dbx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, user *User) error {
	err := r.db.
		QueryRow(
			ctx,
			`INSERT INTO users (id)
			VALUES ($1)
			RETURNING created_at`,
			user.ID,
		).Scan(
		&user.CreatedAt,
	)

	switch {
	case dbx.IsUniqueViolation(err):
		return apperrors.AlreadyExists("user", "user_id", user.ID)
	case err != nil:
		return apperrors.Internal(err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, userID int) error {
	n, err := r.db.
		Exec(
			ctx,
			`UPDATE users 
			SET deleted_at = NOW()
			WHERE id = $1
			AND deleted_at IS NULL;`,
			userID,
		)
	if err != nil {
		return apperrors.Internal(err)
	}

	if n.RowsAffected() == 0 {
		return apperrors.NotFound("user", "id", userID)
	}

	return nil
}
