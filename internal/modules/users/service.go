package users

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

func (s *Service) Create(ctx context.Context, user *User) error {
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}

	log.FromContext(ctx).Info("user created", "userID", user.ID)
	return nil
}

func (s *Service) Delete(ctx context.Context, userID int) error {
	if err := s.repo.Delete(ctx, userID); err != nil {
		return err
	}

	log.FromContext(ctx).Info("user deleted", "userID", userID)
	return nil
}
