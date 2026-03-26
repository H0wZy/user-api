package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/H0wZy/user-api/internal/model"
	"github.com/H0wZy/user-api/internal/service"
	"github.com/H0wZy/user-api/internal/utils"
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

func (ctrl *UserController) GetByID(ctx *gin.Context) {
	const opr = "GetByID"

	id, err := utils.ParseUintParam(ctx, "id")

	if err != nil {
		sendErrorResponse(ctx, opr, http.StatusBadRequest, err.Error())
		return
	}

	user, err := ctrl.service.GetByID(ctx.Request.Context(), id)

	if err != nil {
		switch {
		case errors.Is(err, utils.UserNotFound):
			sendErrorResponse(ctx, opr, http.StatusNotFound, utils.UserNotFound.Error())
		default:
			sendErrorResponse(ctx, opr, http.StatusInternalServerError, err.Error())
		}
		return
	}

	sendSuccessResponse(ctx, opr, http.StatusOK, user)
}

func (ctrl *UserController) List(ctx *gin.Context) {
	users, err := ctrl.service.List(ctx.Request.Context())

	if err != nil {
		sendErrorResponse(ctx, "List", http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(ctx, "List", http.StatusOK, users)

}

func (ctrl *UserController) Delete(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendErrorResponse(ctx, "Delete", http.StatusBadRequest, "id na query param é obrigatório")
		return
	}

	id64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		sendErrorResponse(ctx, "Delete", http.StatusBadRequest, "id inválido")
		return
	}

	if err := ctrl.service.Delete(ctx.Request.Context(), uint(id64)); err != nil {
		if errors.Is(err, utils.UserNotFound) {
			sendErrorResponse(ctx, "Delete", http.StatusNotFound, "usuário não encontrado")
			return
		}
		sendErrorResponse(ctx, "Delete", http.StatusInternalServerError, "erro ao deletar usuário")
		return
	}

	sendSuccessResponse(ctx, "Delete", http.StatusOK, gin.H{
		"deleted_id": uint(id64),
	})

}
