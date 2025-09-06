package userservice

import (
    "taskuser/internal/repository/inmemory"
)

type UserService struct {
    repo *inmemory.UserRepository
}

func NewUserService(repo *inmemory.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Repository() *inmemory.UserRepository {
    return s.repo
}
