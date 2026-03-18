package controller

import (
	"github.com/H0wZy/user-api/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Initialize_Controller() {
	logger = config.GetLogger("controller")
	db = config.GetSQLite()
}

func CreateUser(ctx *gin.Context) {

}
