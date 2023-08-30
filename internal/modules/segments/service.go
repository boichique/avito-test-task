package segments

import (
	"context"

	"github.com/boichique/avito-test-task/internal/log"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, segment *Segment) error {
	if err := s.repo.Create(ctx, segment); err != nil {
		return err
	}

	log.FromContext(ctx).Info(
		"segment created",
		"slug", segment.Slug,
		"created_at", segment.CreatedAt,
	)
	return nil
}

func (s *Service) Delete(ctx context.Context, slug string) error {
	if err := s.repo.Delete(ctx, slug); err != nil {
		return err
	}

	log.FromContext(ctx).Info(
		"segment deleted",
		"slug", slug,
	)
	return nil
}

func (s *Service) GetByUserID(ctx context.Context, userID int) ([]Segment, error) {
	segments, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return segments, nil
}

func (s *Service) UpdateSegmentsByUserID(ctx context.Context, userID int, addSegments []string, deleteSegments []string) error {
	err := s.repo.UpdateSegmentsByUserID(ctx, userID, addSegments, deleteSegments)
	if err != nil {
		return err
	}

	log.FromContext(ctx).Info(
		"segments updated for user",
		"userID", userID,
	)
	return nil
}
