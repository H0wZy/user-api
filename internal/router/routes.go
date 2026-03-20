package router

import (
	"github.com/H0wZy/user-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func initialize_routes(router *gin.Engine, ctrl *controller.UserController) {
	base_path := "/api/v1"

	v1 := router.Group(base_path)
	{
		v1.POST("/user", ctrl.Create)
	}
}
