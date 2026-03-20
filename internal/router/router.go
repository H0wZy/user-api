package router

import (
	"github.com/H0wZy/user-api/internal/config"
	"github.com/H0wZy/user-api/internal/controller"
	"github.com/H0wZy/user-api/internal/repository"
	"github.com/H0wZy/user-api/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger
)

func Initialize() {
	r := gin.Default()

	db := config.GetSQLite()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	initialize_routes(r, userController)

	if err := r.Run(":8080"); err != nil {
		logger.Errorf("Failed to start server: %v", err)
	}
}
