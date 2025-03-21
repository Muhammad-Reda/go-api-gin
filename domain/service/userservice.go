package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammad-reda/go-api-gin/domain/entity"
	"github.com/muhammad-reda/go-api-gin/domain/repository"
	dto "github.com/muhammad-reda/go-api-gin/dto/user"
)

type UserService interface {
	GetAll(ctx *gin.Context) ([]entity.User, error)
	GetByID(ctx *gin.Context, id int64) (entity.User, error)
	Create(ctx *gin.Context) (*entity.User, error)
	Update(ctx *gin.Context, id int64) (*entity.User, error)
	Delete(ctx *gin.Context, id int64) error
}

type UserServiceImplementation struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImplementation{
		userRepo: userRepository,
	}
}

func (us *UserServiceImplementation) GetAll(ctx *gin.Context) ([]entity.User, error) {
	return us.userRepo.FindAll(ctx)
}

func (us *UserServiceImplementation) GetByID(ctx *gin.Context, id int64) (entity.User, error) {
	return us.userRepo.FindById(ctx, id)
}

func (us *UserServiceImplementation) Create(ctx *gin.Context) (*entity.User, error) {
	var input dto.UserCrate

	errBindJson := ctx.ShouldBindJSON(&input)
	if errBindJson != nil {
		return nil, errBindJson
	}

	validator := validator.New()
	errValidate := validator.Struct(input)
	if errValidate != nil {
		return nil, errValidate
	}

	user := entity.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	result, errSave := us.userRepo.Save(ctx, user)
	if errSave != nil {
		return nil, errSave
	}

	return result, nil
}

func (us *UserServiceImplementation) Update(ctx *gin.Context, id int64) (*entity.User, error) {
	var input dto.UserUpdate

	errBindJson := ctx.ShouldBindJSON(&input)
	if errBindJson != nil {
		return nil, errBindJson
	}

	validator := validator.New()
	errValidate := validator.Struct(input)
	if errValidate != nil {
		return nil, errValidate
	}

	user := entity.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	result, errUpdate := us.userRepo.Update(ctx, user, id)
	if errUpdate != nil {
		return nil, errUpdate
	}

	return result, nil
}

func (us *UserServiceImplementation) Delete(ctx *gin.Context, id int64) error {
	errDelete := us.userRepo.Delete(ctx, id)
	return errDelete
}
