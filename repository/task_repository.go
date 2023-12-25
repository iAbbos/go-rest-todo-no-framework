package repository

import (
	"context"
	"golang-crud/models"
)

type TaskRepository interface {
	Save(ctx context.Context, task models.Task)
	Update(ctx context.Context, task models.Task)
	Delete(ctx context.Context, taskId int)
	FindById(ctx context.Context, taskId int) (models.Task, error)
	FindAll(ctx context.Context) []models.Task
}
