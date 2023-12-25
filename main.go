package main

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helpers"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")

	// database
	db := config.DatabaseConnection()

	// repository
	taskRepository := repository.NewTaskRepository(db)

	// service
	taskService := service.NewTaskServiceImpl(taskRepository)

	// controller
	taskController := controller.NewTaskController(taskService)

	// routes
	routes := router.NewRouter(taskController)

	server := http.Server{Addr: "localhost:8881", Handler: routes}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
