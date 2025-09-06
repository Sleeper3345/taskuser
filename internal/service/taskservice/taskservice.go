package taskservice

import (
    "taskuser/internal/repository/inmemory"
)

type TaskService struct {
    repo *inmemory.TaskRepository
}

func NewTaskService(repo *inmemory.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) Repository() *inmemory.TaskRepository {
    return s.repo
}
