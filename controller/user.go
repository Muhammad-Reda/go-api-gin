package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/domain/service"
)

type UserController struct {
	userService service.UserService
	ctx         *gin.Context
}

func NewUserController(userService service.UserService, ctx *gin.Context) UserController {
	return UserController{
		userService: userService,
		ctx:         ctx,
	}
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.userService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (uc *UserController) GetByID(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errStrConv := strconv.ParseInt(param, 10, 64)
	if errStrConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errStrConv.Error(),
		})
	}

	user, err := uc.userService.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uc *UserController) Create(ctx *gin.Context) {
	user, err := uc.userService.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

func (uc *UserController) Update(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errStrConv := strconv.ParseInt(param, 10, 64)
	if errStrConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errStrConv.Error(),
		})
	}

	user, err := uc.userService.Update(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uc *UserController) Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	id, errStrConv := strconv.ParseInt(param, 10, 64)
	if errStrConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errStrConv.Error(),
		})
	}

	err := uc.userService.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
