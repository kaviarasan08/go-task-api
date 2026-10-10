package handler

import (
	"encoding/json"
	"fmt"
	"go_task_api/internal/model"
	"go_task_api/internal/service"
	"log"
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

type updateTaskRequest struct {
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

type deleteTaskRepsonse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var req createTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Reuest", http.StatusBadRequest) // sends error repsonse
		return
	}

	task, err := h.service.CreateTask(r.Context(), req.Title)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	task, err := h.service.GetTaskById(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTaskById(w http.ResponseWriter, r *http.Request) {
    var response deleteTaskRepsonse

    id := r.PathValue("id")

    status, err := h.service.DeleteTaskById(r.Context(), id)
	str := err.Error()
    if err != nil {
        http.Error(w, str, http.StatusInternalServerError)
		fmt.Println(err)
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

	json.NewDecoder(r.Body).Decode(&request)

    id := r.PathValue("id")

	task.ID = id
	task.Title = request.Title
	task.Completed = request.Completed


    status, err := h.service.UpdateTaskById(r.Context(), task)
	// str := err.Error()
    if err != nil {
        http.Error(w, "Failed to update task", http.StatusInternalServerError)
		fmt.Println(err)
        return
    }

    response.Message = status
    response.Status = http.StatusOK

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(response)
}