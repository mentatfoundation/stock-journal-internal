package authentication

import (
	"github.com/labstack/echo/v4"
)

type baseResponse struct {
	Code    int
	Message string
}

func BindErrorResponse(ctx echo.Context) interface{} {
	return baseResponse{
		Code:    400,
		Message: "Unable to process request",
	}
}
