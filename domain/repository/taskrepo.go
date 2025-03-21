package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/muhammad-reda/go-api-gin/domain/entity"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]entity.Task, error)
	FindById(ctx context.Context, id int64) (entity.Task, error)
	Save(tctx context.Context, task entity.Task) (*entity.Task, error)
	Update(tctx context.Context, task entity.Task, id int64) (*entity.Task, error)
	Delete(tctx context.Context, id int64) error
}

type TaskRepositoryImplementation struct {
	DB *sql.DB
}

func NewTaskRepository(DB *sql.DB) *TaskRepositoryImplementation {
	return &TaskRepositoryImplementation{
		DB,
	}
}

func (ts *TaskRepositoryImplementation) FindAll(ctx context.Context) ([]entity.Task, error) {
	var tasks []entity.Task

	query := "SELECT id, name, description, status, user_id, created_at, updated_at, deleted_at  FROM tasks WHERE deleted_at IS NULL"
	rows, err := ts.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Status, &task.UserId, &task.CreatedAt, &task.UpdatedAt, &task.DeletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskRepositoryImplementation) FindById(ctx context.Context, id int64) (entity.Task, error) {
	var task entity.Task

	query := "SELECT id, name, description, status, user_id, created_at, updated_at, deleted_at FROM tasks WHERE id = ? AND deleted_at IS NULL"
	db := ts.DB

	err := db.QueryRowContext(ctx, query, id).Scan(&task.Id, &task.Name, &task.Description, &task.Status, &task.UserId, &task.CreatedAt, &task.UpdatedAt, &task.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return task, fmt.Errorf("task with id %d not found", id)
		}
		return task, err
	}

	return task, nil
}

func (ts *TaskRepositoryImplementation) Save(ctx context.Context, task entity.Task) (*entity.Task, error) {
	query := "INSERT INTO tasks (name, description, status, user_id) VALUES (?, ?, ?, ?)"
	db := ts.DB

	_, errQuery := db.ExecContext(ctx, query, task.Name, task.Description, task.Status, task.UserId)
	if errQuery != nil {
		return nil, errQuery
	}

	return &task, nil
}

func (ts *TaskRepositoryImplementation) Update(ctx context.Context, task entity.Task, id int64) (*entity.Task, error) {
	query := "UPDATE tasks SET name = ?, description = ?, status = ?, user_id = ? WHERE id = ? AND deleted_at IS NULL"

	_, errQuery := ts.DB.ExecContext(ctx, query, task.Name, task.Description, task.Status, task.UserId, id)
	if errQuery != nil {
		return nil, errQuery
	}

	return &task, nil
}

func (ts *TaskRepositoryImplementation) Delete(ctx context.Context, id int64) error {
	query := "INSERT INTO tasks (deleted_at) VALUES (?) WHERE id = ? WHERE deleted_at IS NULL"

	_, errQuery := ts.DB.ExecContext(ctx, query, time.Now(), id)
	if errQuery != nil {
		return errQuery
	}
	return nil
}
