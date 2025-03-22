package service

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/domain/entity"
	"github.com/muhammad-reda/go-api-gin/domain/repository"
	dto "github.com/muhammad-reda/go-api-gin/dto/user"
	validation "github.com/muhammad-reda/go-api-gin/validation/body"
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

	if veer := validation.BodyValidation(ctx, &input); veer != nil {
		return nil, veer
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
	foundedUser, errorFound := us.userRepo.FindById(ctx, id)
	if errorFound != nil {
		return nil, errorFound
	}
	if input.Email == "" {
		input.Email = foundedUser.Email
	}
	if input.Password == "" {
		input.Password = foundedUser.Password
	}
	if input.Username == "" {
		input.Username = foundedUser.Username
	}

	if veer := validation.BodyValidation(ctx, &input); veer != nil {
		return nil, veer
	}

	user := entity.User{
		Id:        foundedUser.Id,
		Email:     input.Email,
		Username:  input.Username,
		Password:  input.Password,
		CreatedAt: foundedUser.CreatedAt,
		UpdatedAt: foundedUser.UpdatedAt,
		DeletedAt: foundedUser.DeletedAt,
	}

	result, errUpdate := us.userRepo.Update(ctx, user, id)
	if errUpdate != nil {
		return nil, errUpdate
	}

	return result, nil
}

func (us *UserServiceImplementation) Delete(ctx *gin.Context, id int64) error {
	_, errorFound := us.userRepo.FindById(ctx, id)
	if errorFound != nil {
		return errorFound
	}
	errDelete := us.userRepo.Delete(ctx, id)
	return errDelete
}
