package services

import (
	"errors"
	"github.com/RostislavOrlov/krp_admin/internal/dto"
	"github.com/RostislavOrlov/krp_admin/internal/repositories"
	"github.com/RostislavOrlov/krp_admin/internal/utils"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) (*UserService, error) {
	return &UserService{
		repo: repo,
	}, nil
}

func (srv *UserService) AddUser(req dto.AddUserRequest) (*dto.AddUserResponse, error) {
	password := utils.GeneratePassword(req)
	req.Password = password
	usr, err := srv.repo.AddUser(req)
	if err != nil {
		return nil, errors.New("failed user registration")
	}

	return usr, nil
}

func (srv *UserService) EditUser(req dto.EditUserRequest) (*dto.EditUserResponse, error) {
	return srv.repo.EditUser(req)
}

func (srv *UserService) DeleteUser(req dto.DeleteUserRequest) (*dto.DeleteUserResponse, error) {
	return srv.repo.DeleteUser(req)
}

func (srv *UserService) ListUsers() (*dto.ListUserResponse, error) {
	return srv.repo.ListUsers()
}
