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

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

// Create
func (controller *BookController) Create(write http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helpers.ReadRequestBody(requests, &bookCreateRequest)

	controller.BookService.Create(requests.Context(), bookCreateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helpers.WriteResponseBody(write, webResponse)
}

// Update
func (controller *BookController) Update(write http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helpers.ReadRequestBody(requests, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helpers.PanicIfError(err)
	bookUpdateRequest.ID = id
	controller.BookService.Update(requests.Context(), bookUpdateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helpers.WriteResponseBody(write, webResponse)
}

// Delete
func (controller *BookController) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helpers.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helpers.WriteResponseBody(write, webResponse)
}

// FindAll
func (controller *BookController) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	result := controller.BookService.FindAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helpers.WriteResponseBody(write, webResponse)
}

// FindById
func (controller *BookController) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helpers.PanicIfError(err)

	result := controller.BookService.FindById(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helpers.WriteResponseBody(write, webResponse)
}
