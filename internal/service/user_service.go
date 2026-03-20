package service

import (
	"context"
	"errors"
	"strings"

	"github.com/H0wZy/user-api/internal/model"
	"github.com/H0wZy/user-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uint) (*model.User, error)
	List(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, id uint, input *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func (s *userService) Create(ctx context.Context, user *model.User) error {
	existingEmail, _ := s.repo.GetByEmail(ctx, user.Email)

	if existingEmail != nil {
		return errors.New("E-mail já cadastrado")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hash)

	return s.repo.Create(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	_, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}

func (s *userService) GetByID(ctx context.Context, id uint) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) List(ctx context.Context) ([]model.User, error) {
	return s.repo.List(ctx)
}

func (s *userService) Update(ctx context.Context, id uint, input *model.User) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(input.FirstName) != "" {
		user.FirstName = input.FirstName
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
