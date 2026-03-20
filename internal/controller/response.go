package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func sendSuccessResponse(ctx *gin.Context, opr string, code int, data any) {
	ctx.JSON(code, Response{
		Success: true,
		Message: fmt.Sprintf("operação '%s' concluída com sucesso!", opr),
		Data:    data,
	})
}

func sendErrorResponse(ctx *gin.Context, opr string, code int, msg string) {
	ctx.JSON(code, Response{
		Success: false,
		Message: fmt.Sprintf("operação '%s' falhou, motivo: %s", opr, msg),
	})
}
