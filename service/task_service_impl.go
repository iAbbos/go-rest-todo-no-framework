package service

import (
	"context"
	"golang-crud/helpers"
	"golang-crud/models"
	"golang-crud/repository"
	"golang-crud/request"
	"golang-crud/response"
)

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
}

func NewTaskServiceImpl(taskRepository repository.TaskRepository) TaskService {
	return &TaskServiceImpl{TaskRepository: taskRepository}
}

func (taskService *TaskServiceImpl) Create(ctx context.Context, request request.TaskCreateRequest) {
	task := models.Task{
		Name:   request.Name,
		Note:   request.Note,
		Status: models.Status(request.Status),
		Date:   request.Date,
	}
	taskService.TaskRepository.Save(ctx, task)
}

func (taskService *TaskServiceImpl) Delete(ctx context.Context, taskId int) {
	task, err := taskService.TaskRepository.FindById(ctx, taskId)
	helpers.PanicIfError(err)
	taskService.TaskRepository.Delete(ctx, task.ID)
}

func (taskService *TaskServiceImpl) FindAll(ctx context.Context) []response.TaskResponse {

	tasks := taskService.TaskRepository.FindAll(ctx)
	var taskResponses []response.TaskResponse

	for _, task := range tasks {
		task := response.TaskResponse{
			ID:     task.ID,
			Name:   task.Name,
			Note:   task.Note,
			Status: string(task.Status),
			Date:   task.Date,
		}
		taskResponses = append(taskResponses, task)
	}
	return taskResponses
}

func (taskService *TaskServiceImpl) FindById(ctx context.Context, taskID int) response.TaskResponse {
	task, err := taskService.TaskRepository.FindById(ctx, taskID)
	helpers.PanicIfError(err)

	taskResponse := response.TaskResponse{
		ID:     task.ID,
		Name:   task.Name,
		Note:   task.Note,
		Status: string(task.Status),
		Date:   task.Date,
	}

	return taskResponse
}

func (taskService *TaskServiceImpl) Update(ctx context.Context, request request.TaskUpdateRequest) {
	task, err := taskService.TaskRepository.FindById(ctx, request.ID)
	helpers.PanicIfError(err)

	task.Name = request.Name
	task.Note = request.Note
	task.Status = models.Status(request.Status)
	task.Date = request.Date

	taskService.TaskRepository.Update(ctx, task)
}
