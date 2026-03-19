package controller

import (
	"net/http"

	"github.com/H0wZy/user-api/internal/config"
	"github.com/H0wZy/user-api/internal/model"
	"github.com/H0wZy/user-api/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) Create(ctx *gin.Context) {
	var user model.User
	var request CreateUserRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		sendErrorResponse(ctx, "Create", http.StatusBadRequest, err.Error())
		return
	}

	request.ValidateGender()

	if err := ctrl.service.Create(ctx.Request.Context(), &user); err != nil {
		sendErrorResponse(ctx, "Create", http.StatusBadRequest, err.Error())
		return
	}

	sendSuccessResponse(ctx, "Create", user)

}

var (
	logger *config.Logger
	db     *gorm.DB
)

func Initialize_Controller() {
	logger = config.GetLogger("controller")
	db = config.GetSQLite()
}

func CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
