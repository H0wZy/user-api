package router

import (
	"github.com/H0wZy/user-api/internal/controller"
	c "github.com/H0wZy/user-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func initialize_routes(router *gin.Engine) {
	controller.Initialize_Controller()

	base_path := "/api/v1"

	v1 := router.Group(base_path)
	{
		v1.GET("/user", c.CreateUser)
		v1.GET("/users", c.CreateUser)
		v1.POST("/user", c.CreateUser)
		v1.PATCH("/user", c.CreateUser)
		v1.DELETE("/user", c.CreateUser)
	}

}
