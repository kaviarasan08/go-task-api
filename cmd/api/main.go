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

	// task dependency injection

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)
	
	// user dependency injection

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	// task api

	mux.HandleFunc("POST /tasks", taskHandler.CreateTask)
	mux.HandleFunc("GET /tasks", taskHandler.GetAllTasks)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetTaskById)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.DeleteTaskById)
	mux.HandleFunc("PUT /tasks/{id}", taskHandler.UpdateTaskById)

	// user api

	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)

	log.Println("Server is Running on : 8080")

	err = http.ListenAndServe( ":8080", mux)

	if err!= nil {
		log.Fatal(err)
	}


}