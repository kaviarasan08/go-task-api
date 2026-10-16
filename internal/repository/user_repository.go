package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go_task_api/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, name string, email string, passwordHash string) (*model.User, error) {
	var user model.User

	query := `
	INSERT INTO USERS(name, email, password_hash)
	VALUES ($1, $2, $3)
	RETURNING id, name, email, password_hash, created_at, updated_at
	`

	err := u.db.QueryRowContext(ctx, query, name, email, passwordHash).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err!= nil {
		return nil, fmt.Errorf("create user : %w", err)
	}

	return &user, nil

}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error){
	
	var user model.User

	query := `
	SELECT name, email, password_hash, created_at, updated_at
	FROM USERS
	WHERE email=$1
	`
	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&user.Name, &user.Email, 
		&user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil

}
