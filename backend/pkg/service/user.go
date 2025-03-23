package service

import (
	"backend/gen/user"
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"errors"
	"time"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) Create(ctx context.Context, u *user.User) (*user.UserResult, error) {
	userModel := &model.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	createdUser, err := s.userRepository.Create(userModel)
	if err != nil {
		return nil, err
	}

	return &user.UserResult{
		ID:        int(createdUser.ID),
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: createdUser.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) Login(ctx context.Context, u *user.LoginPayload) (*user.UserResult, error) {
	foundUser, err := s.userRepository.FindByEmail(u.Email)
	if err != nil {
		return nil, err
	}

	if foundUser.Password != u.Password {
		return nil, errors.New("invalid password")
	}

	return &user.UserResult{
		ID:        int(foundUser.ID),
		Name:      foundUser.Name,
		Email:     foundUser.Email,
		CreatedAt: foundUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: foundUser.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) Get(ctx context.Context, u *user.GetPayload) (*user.UserResult, error) {
	foundUser, err := s.userRepository.FindById(uint(u.UserID))
	if err != nil {
		return nil, err
	}

	return &user.UserResult{
		ID:        int(foundUser.ID),
		Name:      foundUser.Name,
		Email:     foundUser.Email,
		CreatedAt: foundUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: foundUser.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) Update(ctx context.Context, u *user.UpdatePayload) (*user.UserResult, error) {
	foundUser, err := s.userRepository.FindById(uint(u.UserID))
	if err != nil {
		return nil, err
	}

	foundUser.Name = *u.Name
	foundUser.Email = *u.Email
	updatedUser, err := s.userRepository.Update(foundUser)
	if err != nil {
		return nil, err
	}

	return &user.UserResult{
		ID:        int(updatedUser.ID),
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedUser.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) Delete(ctx context.Context, u *user.DeletePayload) error {
	foundUser, err := s.userRepository.FindById(uint(u.UserID))
	if err != nil {
		return err
	}
	err = s.userRepository.Delete(foundUser)
	if err != nil {
		return err
	}
	return nil
}
