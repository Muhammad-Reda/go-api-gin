package test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/muhammad-reda/go-api-gin/domain/entity"
	"github.com/muhammad-reda/go-api-gin/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestTaskFindAll(t *testing.T) {
	// Setup database mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	// Buat repository dengan DB mock
	repo := repository.NewTaskRepository(db)
	ctx := context.Background()

	// Definisikan test cases
	tests := []struct {
		name        string
		setupMock   func()
		wantTasks   []entity.Task
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Success - tasks found",
			setupMock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "status", "user_id", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "Task 1", "Desc 1", "pending", 1, time.Now(), time.Now(), nil).
					AddRow(2, "Task 2", "Desc 2", "done", 2, time.Now(), time.Now(), nil)

				mock.ExpectQuery("SELECT id, name, description, status, user_id, created_at, updated_at, deleted_at FROM tasks WHERE deleted_at IS NULL").
					WillReturnRows(rows)
			},
			wantTasks: []entity.Task{
				{Id: 1, Name: "Task 1", Description: "Desc 1", Status: "pending", UserId: 1},
				{Id: 2, Name: "Task 2", Description: "Desc 2", Status: "done", UserId: 2},
			},
			wantErr: false,
		},
	}

	// Jalankan test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			tt.setupMock()

			// Jalankan fungsi yang di-test
			tasks, err := repo.FindAll(ctx)

			// Cetak data yang didapat
			t.Logf("Data yang didapat dari FindAll: %+v", tasks)
			if err != nil {
				t.Logf("Error yang didapat: %v", err)
			}

			// Verifikasi semua ekspektasi mock terpenuhi
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %v", err)
				return
			}

			// Verifikasi hasil
			if tt.wantErr {
				assert.Error(t, err, "expected an error but got none")
				if tt.expectedErr != nil {
					assert.Equal(t, tt.expectedErr, err, "expected error did not match")
				}
			} else {
				assert.NoError(t, err, "unexpected error occurred")
				assert.Equal(t, len(tt.wantTasks), len(tasks), "length of returned tasks does not match expected")

				if len(tasks) != len(tt.wantTasks) {
					return
				}

				for i, wantTask := range tt.wantTasks {
					assert.Equal(t, wantTask.Id, tasks[i].Id, "task ID mismatch at index %d", i)
					assert.Equal(t, wantTask.Name, tasks[i].Name, "task Name mismatch at index %d", i)
					assert.Equal(t, wantTask.Description, tasks[i].Description, "task Description mismatch at index %d", i)
					assert.Equal(t, wantTask.Status, tasks[i].Status, "task Status mismatch at index %d", i)
					assert.Equal(t, wantTask.UserId, tasks[i].UserId, "task UserId mismatch at index %d", i)
				}
			}
		})
	}
}

func TestTaskFindById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Error creating mock: %v", err)
	}
	defer db.Close()

	repo := repository.NewTaskRepository(db)
	ctx := context.Background()

	tests := []struct {
		name        string
		setupMock   func()
		taskId      int64
		wantTask    entity.Task
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Success - tasks found",
			setupMock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "status", "user_id", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "Task 1", "Desc 1", "pending", 1, time.Now(), time.Now(), nil)

				// Gunakan regexp untuk query matching
				mock.ExpectQuery(`SELECT id, name, description, status, user_id, created_at, updated_at, deleted_at FROM tasks WHERE id = \? AND deleted_at IS NULL`).
					WithArgs(1).
					WillReturnRows(rows)
			},
			taskId: 1,
			wantTask: entity.Task{
				Id: 1, Name: "Task 1", Description: "Desc 1", Status: "pending", UserId: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			got, err := repo.FindById(ctx, tt.taskId)

			t.Logf("Data yang didapat dari FindById: %+v", got)
			if err != nil {
				t.Logf("Error yang didapat: %v", err)
			}

			// Periksa ekspektasi mock
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					assert.Equal(t, tt.expectedErr, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantTask.Id, got.Id)
				assert.Equal(t, tt.wantTask.Name, got.Name)
				assert.Equal(t, tt.wantTask.Description, got.Description)
				assert.Equal(t, tt.wantTask.Status, got.Status)
				assert.Equal(t, tt.wantTask.UserId, got.UserId)
			}
		})
	}
}
