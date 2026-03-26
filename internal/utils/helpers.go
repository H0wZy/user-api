package utils

import (
	"fmt"
	"strconv"

	"github.com/H0wZy/user-api/internal/config"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger
)

func ParseUintParam(ctx *gin.Context, key string) (uint, error) {
	logger = config.GetLogger("helpers")
	value := ctx.Query(key)

	if value == "" {
		return 0, fmt.Errorf("param %s é obrigatório", key)
	}

	v, err := strconv.ParseUint(value, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("param '%s' inválido: %w", key, err)
	}

	return uint(v), nil
}
