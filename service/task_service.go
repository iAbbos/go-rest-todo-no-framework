package service

import (
	"context"
	"golang-crud/request"
	"golang-crud/response"
)

type TaskService interface {
	Create(ctx context.Context, request request.TaskCreateRequest)
	Update(ctx context.Context, request request.TaskUpdateRequest)
	Delete(ctx context.Context, taskId int)
	FindById(ctx context.Context, taskID int) response.TaskResponse
	FindAll(ctx context.Context) []response.TaskResponse
}
