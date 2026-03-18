package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()
	initialize_routes(r)
	_ = r.Run(":8080")
}
