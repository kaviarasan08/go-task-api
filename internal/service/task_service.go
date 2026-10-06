package service

import (
	"context"
	"go_task_api/internal/model"
	"go_task_api/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo : repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, title string) (*model.Task, error) {

	task, err := s.repo.CreateTask(ctx, title)

	if err!= nil {
		return nil, err
	}

	return task, nil

}
