package service

import (
	"context"
	"errors"
	"strings"

	"github.com/H0wZy/user-api/internal/model"
	"github.com/H0wZy/user-api/internal/repository"
	"github.com/H0wZy/user-api/internal/utils"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	existingEmail, errEmail := s.repo.GetByEmail(ctx, user.Email)
	existingUsername, errUsername := s.repo.GetByUsername(ctx, user.Username)

	if errEmail != nil && !errors.Is(errEmail, utils.EmailAlreadyExists) && !errors.Is(errEmail, gorm.ErrRecordNotFound) {
		return errEmail
	}
	if errUsername != nil && !errors.Is(errUsername, utils.UsernameAlreadyExists) && !errors.Is(errUsername, gorm.ErrRecordNotFound) {
		return errUsername
	}

	if existingEmail != nil {
		return utils.EmailAlreadyExists
	}

	if existingUsername != nil {
		return utils.UsernameAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	if err := s.repo.Create(ctx, user); err != nil {
		if strings.Contains(err.Error(), "username") {
			return utils.UsernameAlreadyExists
		}
		if strings.Contains(err.Error(), "email") {
			return utils.EmailAlreadyExists
		}
		return err
	}

	return nil
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) GetByID(ctx context.Context, id uint) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.UserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) List(ctx context.Context) ([]model.User, error) {
	users, err := s.repo.List(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) Update(ctx context.Context, id uint, input *model.User) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(user, input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	if strings.TrimSpace(input.Password) != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hash)
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
