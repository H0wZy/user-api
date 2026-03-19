package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	StatusCode int    `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

func sendSuccessResponse(ctx *gin.Context, opr string, data any) {
	ctx.JSON(http.StatusOK, SuccessResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    fmt.Sprintf("operação do controller '%s' concluída com sucesso!", opr),
		Data:       data,
	})
}

func sendErrorResponse(ctx *gin.Context, opr string, code int, msg string) {
	ctx.JSON(code, ErrorResponse{
		StatusCode: code,
		Success:    false,
		Message:    fmt.Sprintf("operação do controller '%s' falhou, motivo: %s", opr, msg),
	})
}
