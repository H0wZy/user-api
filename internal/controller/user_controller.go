package controller

import (
	"net/http"

	"github.com/H0wZy/user-api/internal/model"
	"github.com/H0wZy/user-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) Create(ctx *gin.Context) {
	var request CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		sendErrorResponse(ctx, "Create", http.StatusBadRequest, err.Error())
		return
	}

	if err := request.ValidateGender(); err != nil {
		sendErrorResponse(ctx, "Create", http.StatusBadRequest, err.Error())
		return
	}

	user := model.User{
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  request.Password,
		Phone:     request.Phone,
		Gender:    request.Gender,
		BirthDate: request.BirthDate,
	}

	if err := ctrl.service.Create(ctx.Request.Context(), &user); err != nil {
		sendErrorResponse(ctx, "Create", http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(ctx, "Create", http.StatusCreated, user)
}

func (ctrl *UserController) GetByID(ctx *gin.Context, id uint) {
	userId := model.User{

		// TODO
		// gorm.Model{ID}: id,
	}

	if err := ctrl.service.GetByID(ctx.Request.Context(), &user); err != nil {
		sendErrorResponse(ctx, "GetByID", http.StatusOK, err.Error())
		return
	}

	sendSuccessResponse(ctx, "GetByID", http.StatusOK, userId)
}
