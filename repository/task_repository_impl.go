package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-crud/config"
	"golang-crud/helpers"
	"golang-crud/models"
)

const DB_NAME = config.DB_NAME

type TaskRepositoryImpl struct {
	Db *sql.DB
}

func NewTaskRepository(Db *sql.DB) TaskRepository {
	return &TaskRepositoryImpl{Db: Db}
}

func (taskRepo *TaskRepositoryImpl) Delete(ctx context.Context, taskId int) {
	tx, err := taskRepo.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := fmt.Sprintf("DELETE FROM %s WHERE id = $1", DB_NAME)
	_, errExec := tx.ExecContext(ctx, SQL, taskId)
	helpers.PanicIfError(errExec)
}

func (taskRepo *TaskRepositoryImpl) FindAll(ctx context.Context) []models.Task {
	tx, err := taskRepo.Db.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	SQL := fmt.Sprintf("SELECT * FROM %s", DB_NAME)
	result, errQuery := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(errQuery)

	defer result.Close()

	var tasks []models.Task

	for result.Next() {
		var task models.Task

		errScan := result.Scan(&task.ID, &task.Name)
		helpers.PanicIfError(errScan)

		tasks = append(tasks, task)
	}

	return tasks
}

func (taskRepo *TaskRepositoryImpl) FindById(ctx context.Context, taskId int) (models.Task, error) {
	tx, err := taskRepo.Db.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	SQL := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", DB_NAME)
	result, errQuery := tx.QueryContext(ctx, SQL, taskId)

	helpers.PanicIfError(errQuery)
	defer result.Close()

	var task models.Task

	if result.Next() {
		errScan := result.Scan(&task.ID, &task.Name)
		helpers.PanicIfError(errScan)
		return task, nil
	} else {
		return task, errors.New("Task is not found")
	}
}

func (taskRepo *TaskRepositoryImpl) Save(ctx context.Context, task models.Task) {
	tx, err := taskRepo.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := fmt.Sprintf("INSERT INTO %s (name, note, status, date) VALUES ($1, $2, $3, $4)", DB_NAME)

	_, errExec := tx.ExecContext(
		ctx,
		SQL,
		task.Name,
		task.Note,
		task.Status,
		task.Date,
	)
	helpers.PanicIfError(errExec)
}

func (taskRepo *TaskRepositoryImpl) Update(ctx context.Context, task models.Task) {
	tx, err := taskRepo.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := fmt.Sprintf("UPDATE %s SET name = $2, note = $3, status = $4, task_date = $5 WHERE id = $1", DB_NAME)

	_, errExec := tx.ExecContext(
		ctx,
		SQL,
		task.ID,
		task.Name,
		task.Note,
		task.Status,
		task.Date,
	)
	helpers.PanicIfError(errExec)
}
