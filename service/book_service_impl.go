package service

import (
	"context"

	"golang-crud/helpers"
	"golang-crud/models"
	"golang-crud/repository"
	"golang-crud/request"
	"golang-crud/response"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

// Create implements BookService.
func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) {
	book := models.Book{
		Name: request.Name,
	}
	b.BookRepository.Save(ctx, book)
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helpers.PanicIfError(err)
	b.BookRepository.Delete(ctx, book.ID)
}

// FindAll implements BookService.
func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	var bookResponses []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{
			ID:   value.ID,
			Name: value.Name,
		}
		bookResponses = append(bookResponses, book)
	}
	return bookResponses
}

// FindById implements BookService.
func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helpers.PanicIfError(err)

	return response.BookResponse(book)
}

// Update implements BookService.
func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.ID)
	helpers.PanicIfError(err)

	book.Name = request.Name
	b.BookRepository.Update(ctx, book)
}
