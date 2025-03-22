package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/domain/service"
)

var (
	ErrInternalServerError  = "Internal Server Error"
	ErrFailedToGetAllTask   = "Failed to get tasks"
	ErrFailedToGetTask      = "Failed to get task"
	ErrFailedToCrateTask    = "Failed to create task"
	ErrFailedToUpdateTask   = "Failed to update task"
	ErrFailedToDeleteTask   = "Failed to delete task"
	ErrFailedToProcessParam = "Failed to prcess param"
)

type TaskHandler struct {
	taskService service.Taskservice
	ctx         *gin.Context
}

func NewTaskHandler(taskService service.Taskservice, ctx *gin.Context) TaskHandler {
	return TaskHandler{
		taskService: taskService,
		ctx:         ctx,
	}
}

func (th *TaskHandler) GetAllTask(ctx *gin.Context) {
	data, err := th.taskService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": ErrFailedToGetAllTask,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    data,
	})
}

func (th *TaskHandler) GetTaskById(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errConv := strconv.ParseInt(param, 10, 64)
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToProcessParam,
			"error":   "id param should be a number",
		})
		return
	}

	data, err := th.taskService.GetById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "task not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": ErrFailedToGetTask,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    data,
	})
}

func (th *TaskHandler) Create(ctx *gin.Context) {
	_, err := th.taskService.Create(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": ErrFailedToCrateTask,
				"error":   err,
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToCrateTask,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task created successfully",
	})
}

func (th *TaskHandler) Update(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errConv := strconv.ParseInt(param, 10, 64)
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToProcessParam,
			"error":   "id param should be a number",
		})
		return
	}

	data, err := th.taskService.Update(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToUpdateTask,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"data":    data,
	})
}

func (th *TaskHandler) Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errConv := strconv.ParseInt(param, 10, 64)
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToProcessParam,
			"error":   "id param should be a number",
		})
		return
	}

	err := th.taskService.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": ErrFailedToDeleteTask,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
