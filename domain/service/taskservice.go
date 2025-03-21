package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammad-reda/go-api-gin/domain/entity"
	"github.com/muhammad-reda/go-api-gin/domain/repository"
	dto "github.com/muhammad-reda/go-api-gin/dto/task"
)

type Taskservice interface {
	GetAll(ctx *gin.Context) ([]entity.Task, error)
	GetById(ctx *gin.Context, id int64) (entity.Task, error)
	Create(ctx *gin.Context) (*entity.Task, error)
	Update(ctx *gin.Context, id int64) (*entity.Task, error)
	Delete(ctx *gin.Context, id int64) error
}

type TaskServiceImplementation struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskrepo repository.TaskRepository) Taskservice {
	return &TaskServiceImplementation{
		taskRepo: taskrepo,
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

	errBindJson := ctx.ShouldBindJSON(&input)
	if errBindJson != nil {
		return nil, errBindJson
	}

	validator := validator.New()
	errValidate := validator.Struct(input)
	if errValidate != nil {
		return nil, errValidate
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

	errBindJson := ctx.ShouldBindJSON(&input)
	if errBindJson != nil {
		return nil, errBindJson
	}

	validator := validator.New()
	errValidate := validator.Struct(input)
	if errValidate != nil {
		return nil, errValidate
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
	errDelete := ts.taskRepo.Delete(ctx, id)
	return errDelete
}
