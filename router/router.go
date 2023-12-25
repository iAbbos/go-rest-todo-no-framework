package router

import (
	"fmt"
	"golang-crud/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(taskController *controller.TaskController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "I'm Working!")
	})

	router.GET("/api/tasks", taskController.FindAll)
	router.GET("/api/tasks/:taskId", taskController.FindById)
	router.POST("/api/tasks", taskController.Create)
	router.PUT("/api/tasks/:taskId", taskController.Update)
	router.DELETE("/api/tasks/:taskId", taskController.Delete)

	return router
}
