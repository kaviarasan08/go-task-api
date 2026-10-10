package main

import (
	"fmt"
	"log"
	"net/http"

	"go_task_api/internal/database"
	"go_task_api/internal/handler"
	"go_task_api/internal/repository"
	"go_task_api/internal/service"
)

func main() {

	fmt.Println("Task Api Building.....")

	db, err := database.NewPostgresDb()

	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println("DB Connected Sucessfully..")

	taskRepository := repository.NewTaskRepository(db)

	taskService := service.NewTaskService(taskRepository)

	taskHandler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /tasks", taskHandler.CreateTask)

	mux.HandleFunc("GET /tasks", taskHandler.GetAllTasks)

	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetTaskById)

	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.DeleteTaskById)

	mux.HandleFunc("PUT /tasks/{id}", taskHandler.UpdateTaskById)

	log.Println("Server is RUnning on : 8080")

	err = http.ListenAndServe( ":8080", mux)

	if err!= nil {
		log.Fatal(err)
	}


}