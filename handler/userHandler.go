package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/domain/service"
)

var (
	ErrFailedToGetAllUsers = "Failed to get tasks"
	ErrFailedToGetUser     = "Failed to get task"
	ErrFailedToCrateUser   = "Failed to create task"
	ErrFailedToUpdateUser  = "Failed to update task"
	ErrFailedToDeleteUser  = "Failed to delete task"
)

type UserHandler struct {
	userService service.UserService
	ctx         *gin.Context
}

func NewUserHandler(userService service.UserService, ctx *gin.Context) UserHandler {
	return UserHandler{
		userService: userService,
		ctx:         ctx,
	}
}

func (uc *UserHandler) GetAll(ctx *gin.Context) {
	users, err := uc.userService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": ErrInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (uc *UserHandler) GetByID(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errStrConv := strconv.ParseInt(param, 10, 64)
	if errStrConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToProcessParam,
			"error":   "id param should be a number",
		})
	}

	user, err := uc.userService.GetByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": ErrFailedToGetUser,
			"error":   ErrInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uc *UserHandler) Create(ctx *gin.Context) {
	user, err := uc.userService.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToCrateUser,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"data": gin.H{
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (uc *UserHandler) Update(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errStrConv := strconv.ParseInt(param, 10, 64)
	if errStrConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToProcessParam,
			"error":   "id param should be a number",
		})
		return
	}

	user, err := uc.userService.Update(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToUpdateUser,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uc *UserHandler) Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errStrConv := strconv.ParseInt(param, 10, 64)
	if errStrConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToProcessParam,
			"error":   "id param should be a number",
		})
		return
	}

	err := uc.userService.Delete(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrFailedToDeleteUser,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
