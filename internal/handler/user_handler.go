package handler

import (
	"encoding/json"
	"errors"
	"go_task_api/internal/service"
	"net/http"
	"time"
)

type UserHandler struct {
	userService *service.UserService
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writJSONEror(w, "invalid JSON request", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(
		r.Context(),
		req.Name,
		req.Email,
		req.Password,
	)
	if err != nil {
		if errors.Is(err, service.ErrInvalidUserInput) {
			writJSONEror(w, err.Error(), http.StatusBadRequest)
			return
		}

		// writJSONEror(w, "registration failed", http.StatusInternalServerError)
		writJSONEror(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		return
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writJSONEror(w, "invalid JSON request", http.StatusBadRequest)
		return
	}

	user, token, err := h.userService.Login(
		r.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		if errors.Is(err, service.ErrInvalidUserInput) {
			writJSONEror(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err.Error() == "invalid email or password" {
			writJSONEror(w, "invalid email or password", http.StatusUnauthorized)
			return
		}

		writJSONEror(w, "login failed", http.StatusInternalServerError)
		return
	}

	response := struct {
		Token     string    `json:"token"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		Token:     token,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
