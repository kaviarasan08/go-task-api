package repository

import (
	"context"
	"database/sql"
	"go_task_api/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, title string) (*model.Task, error){
	var task model.Task

	query := `
		INSERT INTO tasks(title) 
		VALUES($1)
		RETURNING id, title, completed
	`

	err := r.db.QueryRowContext(ctx, query, title).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
	)

	if err != nil {
		return nil, err
	}

	return &task, err

}