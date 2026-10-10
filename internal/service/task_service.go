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

func (s *TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	tasks, err := s.repo.GetAllTasks(ctx)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) GetTaskById(ctx context.Context, id string) (*model.Task, error) {
	task, err := s.repo.GetTaskById(ctx, id)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) DeleteTaskById(ctx context.Context, id string) (string, error) {
	status, err := s.repo.DeleteTaskById(ctx, id)

	if err != nil {
		return "", err
	}

	return status, nil
}

func (s *TaskService) UpdateTaskById(ctx context.Context, task model.Task) (string, error) {
	status, err := s.repo.UpdateTaskById(ctx, task)

	if err != nil {
		return "", err
	}

	return status, nil
}

