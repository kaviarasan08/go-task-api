package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go_task_api/internal/auth"
	"go_task_api/internal/model"
	"go_task_api/internal/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

var ErrInvalidUserInput = errors.New("name, email, and password are required")

func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Register(ctx context.Context, name, email, password string) (*model.User, error){
	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))
	password = strings.TrimSpace(password)

	if name == "" || password == "" || email == "" {
		return nil, ErrInvalidUserInput
	}

	if len(password) < 8 {
		return nil, errors.New("password must contain at least 8 characters")
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, fmt.Errorf("hash : %w", err)
	}

	user, err := u.repo.CreateUser(ctx, name, email, string(passwordHash))

	if err != nil {
		return nil, fmt.Errorf("register user: %w", err)
	}

	return user, nil

}

func (u *UserService) Login(ctx context.Context, email string, password string) (*model.User, string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return nil, "", ErrInvalidUserInput
	}

	user, err := u.repo.GetUserByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, "", errors.New("invalid email or password")
		}
		return nil, "", fmt.Errorf("find user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateToken(user.ID)

	if err!= nil {
		return nil, "", fmt.Errorf("generate token : %w", err)
	}

	return user, token, nil

}
