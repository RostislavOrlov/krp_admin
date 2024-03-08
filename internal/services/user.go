package services

import (
	"krp_admin/internal/dto"
	"krp_admin/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) (*UserService, error) {
	return &UserService{
		repo: repo,
	}, nil
}

func (srv *UserService) EditUser(req dto.EditUserRequest) (*dto.EditUserResponse, error) {
	return srv.repo.EditUser(req)
}

func (srv *UserService) DeleteUser(req dto.DeleteUserRequest) (*dto.DeleteUserRequest, error) {
	return srv.repo.DeleteUser(req)
}

func (srv *UserService) ListUsers() (*dto.ListUserResponse, error) {
	return srv.repo.ListUsers()
}
