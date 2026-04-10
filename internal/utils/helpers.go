package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/H0wZy/user-api/internal/config"
	"github.com/H0wZy/user-api/internal/types"
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

func isValidGender(gender types.Gender) bool {
	validGenders := types.GetValidGenders()
	return slices.Contains(validGenders, gender)
}

func ValidateGender(genders []types.Gender) error {
	if len(genders) == 0 {
		return EmptyGender
	}

	for _, gender := range genders {
		if gender == "" {
			return EmptyGenderString
		}
		if !isValidGender(gender) {
			validGenders := types.GetValidGenders()
			validGendersStr := make([]string, len(validGenders))
			for i, g := range validGenders {
				validGendersStr[i] = string(g)
			}
			return fmt.Errorf("%w: %s. Valores permitidos: %s", InvalidGender, gender, strings.Join(validGendersStr, ", "))
		}
	}

	return nil
}
