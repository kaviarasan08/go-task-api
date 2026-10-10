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

func (r *TaskRepository) CreateTask(ctx context.Context, title string) (*model.Task, error) {
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

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task

	query := `
	SELECT * FROM tasks
	`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task model.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt, &task.UpdatedAt)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}

func (r *TaskRepository) GetTaskById(ctx context.Context, id string) (*model.Task, error) {
	var task model.Task

	query := `
	SELECT * FROM TASKS WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt, &task.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &task, nil

}

func (r *TaskRepository) DeleteTaskById(ctx context.Context, id string) (string, error) {

	query := `
		DELETE FROM tasks WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", sql.ErrNoRows
	}

	return "Deleted successfully", nil
}

func (r *TaskRepository) UpdateTaskById(ctx context.Context, task model.Task) (string, error) {

	query := `
		UPDATE tasks 
		SET title = $1,
		completed = $2,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, task.Title, task.Completed, task.ID)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", sql.ErrNoRows
	}

	return "Updated successfully", nil
}
