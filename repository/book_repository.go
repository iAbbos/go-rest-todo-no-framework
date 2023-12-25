package repository

import (
	"context"
	"golang-crud/models"
)

type BookRepository interface {
	Save(ctx context.Context, book models.Book)
	Update(ctx context.Context, book models.Book)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) (models.Book, error)
	FindAll(ctx context.Context) []models.Book
}
