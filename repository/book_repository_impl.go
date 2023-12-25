package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helpers"
	"golang-crud/models"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

// Delete implements BookRepository.
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := "DELETE FROM book WHERE id = $1"
	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helpers.PanicIfError(errExec)
}

// FindAll implements BookRepository.
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []models.Book {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	SQL := "SELECT id, name FROM book"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(errQuery)
	defer result.Close()

	var books []models.Book

	for result.Next() {
		var book models.Book

		errScan := result.Scan(&book.ID, &book.Name)
		helpers.PanicIfError(errScan)

		books = append(books, book)
	}

	return books
}

// FindById implements BookRepository.
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (models.Book, error) {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	SQL := "SELECT id, name FROM book WHERE id = $1"
	result, errQuery := tx.QueryContext(ctx, SQL, bookId)

	helpers.PanicIfError(errQuery)
	defer result.Close()

	var book models.Book

	if result.Next() {
		errScan := result.Scan(&book.ID, &book.Name)
		helpers.PanicIfError(errScan)
		return book, nil
	} else {
		return book, errors.New("Book is not found")
	}
}

// Save implements BookRepository.
func (b *BookRepositoryImpl) Save(ctx context.Context, book models.Book) {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := "INSERT INTO book (name) VALUES ($1)"
	_, errExec := tx.ExecContext(ctx, SQL, book.Name)
	helpers.PanicIfError(errExec)
}

// Update implements BookRepository.
func (b *BookRepositoryImpl) Update(ctx context.Context, book models.Book) {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := "UPDATE book SET name = $1 WHERE id = $2"
	_, errExec := tx.ExecContext(ctx, SQL, book.Name, book.ID)
	helpers.PanicIfError(errExec)
}
