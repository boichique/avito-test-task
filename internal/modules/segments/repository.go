package segments

import (
	"context"
	"errors"

	"github.com/boichique/avito-test-task/internal/apperrors"
	"github.com/boichique/avito-test-task/internal/dbx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, segment *Segment) error {
	err := r.db.
		QueryRow(
			ctx,
			`INSERT INTO segments (slug)
			VALUES ($1)
			RETURNING created_at`,
			segment.Slug,
		).Scan(&segment.CreatedAt)

	switch {
	case dbx.IsUniqueViolation(err):
		return apperrors.AlreadyExists("segment", "slug", segment.Slug)
	case err != nil:
		return apperrors.Internal(err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, slug string) error {
	n, err := r.db.
		Exec(
			ctx,
			`UPDATE segments 
			SET deleted_at = NOW()
			WHERE slug = $1
			AND deleted_at IS NULL;`,
			slug,
		)
	if err != nil {
		return apperrors.Internal(err)
	}

	if n.RowsAffected() == 0 {
		return apperrors.NotFound("segment", "slug", slug)
	}

	return nil
}

func (r *Repository) GetByUserID(ctx context.Context, userID int) ([]Segment, error) {
	rows, err := r.db.
		Query(
			ctx,
			`SELECT s.slug, s.created_at
			FROM segments s
			JOIN users_segments us ON s.slug = us.slug
			JOIN users u ON us.user_id = u.id
			WHERE u.id = $1
			AND s.deleted_at IS NULL
			AND u.deleted_at IS NULL`,
			userID,
		)

	switch {
	case dbx.IsNoRows(err):
		return nil, apperrors.NotFound("segments for user", "user_id", userID)
	case err != nil:
		return nil, apperrors.Internal(err)
	}

	var segments []Segment
	for rows.Next() {
		var segment Segment
		if err := rows.Scan(
			&segment.Slug,
			&segment.CreatedAt,
		); err != nil {
			return nil, apperrors.Internal(err)
		}

		segments = append(segments, segment)
	}

	if len(segments) == 0 {
		return nil, apperrors.NotFound("segments for user", "user_id", userID)
	}

	return segments, nil
}

func (r *Repository) UpdateSegmentsByUserID(ctx context.Context, userID int, addSegments []string, deleteSegments []string) error {
	if len(deleteSegments) == 0 && len(addSegments) == 0 {
		return apperrors.BadRequest(errors.New("need to specify at least one segment to delete or add"))
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	// nolint: errcheck
	defer tx.Rollback(ctx)

	if len(deleteSegments) > 0 {
		for _, segment := range deleteSegments {
			_, err = tx.Exec(
				ctx,
				`DELETE FROM users_segments 
				WHERE user_id = $1
				AND slug = $2`,
				userID,
				segment,
			)
			if err != nil {
				return err
			}
		}
	}

	if len(addSegments) > 0 {
		for _, segment := range addSegments {
			_, err = tx.Exec(
				ctx,
				`INSERT INTO users_segments (user_id, slug) 
				VALUES ($1, $2)`,
				userID,
				segment,
			)
			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
