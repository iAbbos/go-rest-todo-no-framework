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
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// routes
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8881", Handler: routes}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
