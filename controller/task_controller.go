package controller

import (
	"golang-crud/helpers"
	"golang-crud/request"
	"golang-crud/response"
	"golang-crud/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TaskController struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

// Create
func (controller *TaskController) Create(write http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	taskCreateRequest := request.TaskCreateRequest{}
	helpers.ReadRequestBody(requests, &taskCreateRequest)

	controller.TaskService.Create(requests.Context(), taskCreateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helpers.WriteResponseBody(write, webResponse)
}

// Update
func (controller *TaskController) Update(write http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	taskUpdateRequest := request.TaskUpdateRequest{}
	helpers.ReadRequestBody(requests, &taskUpdateRequest)

	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helpers.PanicIfError(err)
	taskUpdateRequest.ID = id
	controller.TaskService.Update(requests.Context(), taskUpdateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helpers.WriteResponseBody(write, webResponse)
}

// Delete
func (controller *TaskController) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helpers.PanicIfError(err)

	controller.TaskService.Delete(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helpers.WriteResponseBody(write, webResponse)
}

// FindAll
func (controller *TaskController) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	result := controller.TaskService.FindAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helpers.WriteResponseBody(write, webResponse)
}

// FindById
func (controller *TaskController) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helpers.PanicIfError(err)

	result := controller.TaskService.FindById(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helpers.WriteResponseBody(write, webResponse)
}
