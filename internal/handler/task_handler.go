package handler

import (
	"encoding/json"
	"go_task_api/internal/service"
	"net/http"
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

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var req createTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err!= nil {
		http.Error(w, "Invalid Reuest", http.StatusBadRequest) // sends error repsonse
		return
	}

	task, err := h.service.CreateTask(r.Context(), req.Title)

	if err != nil {
		http.Error(w, "failed to create task", http.StatusInternalServerError)
		return
	}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		// Header → status code → body  remmeber this order
	json.NewEncoder(w).Encode(task)

}