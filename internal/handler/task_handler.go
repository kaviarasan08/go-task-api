package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"go_task_api/internal/model"
	"go_task_api/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

type createTaskRequest struct {
	Title string `json:"title"`
}

type updateTaskRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type deleteTaskRepsonse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var req createTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		// writJSONEror(w, "Invalid Request", http.StatusBadRequest) // sends error repsonse
		writJSONEror(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		// writJSONEror(w, "title cannot be Empty", http.StatusBadRequest)
		writJSONEror(w, "title cannot be Empty", http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(r.Context(), req.Title)

	if err != nil {
		writJSONEror(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Header → status code → body  remmeber this order
	json.NewEncoder(w).Encode(task)

}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := h.service.GetAllTasks(r.Context())

	if err != nil {
		writJSONEror(w, "Internal server error", http.StatusInternalServerError)
		// log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {

	rawId := r.PathValue("id")

	id, err := strconv.Atoi(rawId)

	if err != nil || id <= 0 {
		writJSONEror(w, "Invalid Task Id", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTaskById(r.Context(), id)

	if errors.Is(err, sql.ErrNoRows) {
		writJSONEror(w, "Task not found", http.StatusNotFound)
		return
	}

	if err != nil {
		writJSONEror(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	var response deleteTaskRepsonse

	rawId := r.PathValue("id")

	id, err := strconv.Atoi(rawId)

	if err != nil || id <= 0 {
		writJSONEror(w, "Invalid Task Id", http.StatusBadRequest)
		return
	}

	status, err := h.service.DeleteTaskById(r.Context(), id)

	if errors.Is(err, sql.ErrNoRows) {
		writJSONEror(w, "Task not found", http.StatusNotFound)
		return
	}

	if err != nil {
		writJSONEror(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	response.Message = status
	response.Status = http.StatusOK

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) UpdateTaskById(w http.ResponseWriter, r *http.Request) {
	var request updateTaskRequest
	var response deleteTaskRepsonse
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		// writJSONEror(w, "Invalid Request", http.StatusBadRequest) // sends error repsonse
		writJSONEror(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	rawId := r.PathValue("id")

	id, err := strconv.Atoi(rawId)

	if err != nil || id <= 0 {
		writJSONEror(w, "Invalid Task Id", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(request.Title) == "" {
		writJSONEror(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	task.ID = id
	task.Title = request.Title
	task.Completed = request.Completed

	status, err := h.service.UpdateTaskById(r.Context(), task)

	if errors.Is(err, sql.ErrNoRows) {
		writJSONEror(w, "Task not found", http.StatusNotFound)
		return
	}
	if err != nil {
		writJSONEror(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	response.Message = status
	response.Status = http.StatusOK

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
