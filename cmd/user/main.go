package main

import (
	"github.com/H0wZy/user-api/internal/config"
	"github.com/H0wZy/user-api/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.InitDb()

	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}
	router.Initialize()
}
