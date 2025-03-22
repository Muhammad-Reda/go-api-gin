package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/domain/entity"
	"github.com/muhammad-reda/go-api-gin/domain/repository"
	dto "github.com/muhammad-reda/go-api-gin/dto/task"
	validation "github.com/muhammad-reda/go-api-gin/validation/body"
)

type Taskservice interface {
	GetAll(ctx *gin.Context) ([]entity.Task, error)
	GetById(ctx *gin.Context, id int64) (entity.Task, error)
	Create(ctx *gin.Context) (*entity.Task, error)
	Update(ctx *gin.Context, id int64) (*entity.Task, error)
	Delete(ctx *gin.Context, id int64) error
}

type ErrTaskService struct {
	Reason string `json:"reason"`
}

func (et *ErrTaskService) Error() string {
	return fmt.Sprintf("Reason: %s", et.Reason)
}

type TaskServiceImplementation struct {
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
}

func NewTaskService(taskrepo repository.TaskRepository, userrepo repository.UserRepository) Taskservice {
	return &TaskServiceImplementation{
		taskRepo: taskrepo,
		userRepo: userrepo,
	}
}

func (ts *TaskServiceImplementation) GetAll(ctx *gin.Context) ([]entity.Task, error) {
	return ts.taskRepo.FindAll(ctx)
}

func (ts *TaskServiceImplementation) GetById(ctx *gin.Context, id int64) (entity.Task, error) {
	return ts.taskRepo.FindById(ctx, id)
}

func (ts *TaskServiceImplementation) Create(ctx *gin.Context) (*entity.Task, error) {
	var input dto.TaskCreate

	if veer := validation.BodyValidation(ctx, &input); veer != nil {
		return nil, veer
	}

	_, errFindUser := ts.userRepo.FindById(ctx, input.UserId)
	if errFindUser != nil {
		if errFindUser == sql.ErrNoRows {
			return nil, &ErrTaskService{
				Reason: "user not found",
			}
		}
		return nil, errFindUser
	}

	task := entity.Task{
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
		UserId:      input.UserId,
	}

	result, errSave := ts.taskRepo.Save(ctx, task)
	if errSave != nil {
		return nil, errSave
	}

	return result, nil
}

func (ts *TaskServiceImplementation) Update(ctx *gin.Context, id int64) (*entity.Task, error) {
	var input dto.TaskUpdate
	FoundTask, errFindTask := ts.taskRepo.FindById(ctx, id)
	if errFindTask != nil {
		if errFindTask == sql.ErrNoRows {
			return nil, &ErrTaskService{
				Reason: "task not found",
			}
		}
		return nil, errFindTask
	}

	if input.Name == "" {
		input.Name = FoundTask.Name
	}
	if input.Description == "" {
		input.Description = FoundTask.Description
	}
	if input.Status == "" {
		input.Status = FoundTask.Status
	}
	if input.UserId == 0 {
		input.UserId = FoundTask.UserId
	}

	if veer := validation.BodyValidation(ctx, &input); veer != nil {
		return nil, veer
	}

	_, errFindUser := ts.userRepo.FindById(ctx, input.UserId)
	if errFindUser != nil {
		if errFindUser == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errFindUser
	}

	task := entity.Task{
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
		UserId:      input.UserId,
	}

	result, errUpdate := ts.taskRepo.Update(ctx, task, id)
	if errUpdate != nil {
		return nil, errUpdate
	}

	return result, nil
}

func (ts *TaskServiceImplementation) Delete(ctx *gin.Context, id int64) error {
	_, errFindTask := ts.taskRepo.FindById(ctx, id)
	if errFindTask != nil {
		if errFindTask == sql.ErrNoRows {
			return &ErrTaskService{
				Reason: "task not found",
			}
		}
		return errFindTask
	}
	errDelete := ts.taskRepo.Delete(ctx, id)
	return errDelete
}
