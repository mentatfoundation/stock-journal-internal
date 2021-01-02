package authentication

import (
	"mentatfoundation/stock-journal/server/logger"
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	//AuthService *authService.AuthService
	Logger logger.Logger
}

func New(l logger.Logger) *Handler {
	return &Handler{
		Logger: l,
	}
}

func (a *Handler) Login(c echo.Context) error {
	a.Logger.Info("hello")
	return c.String(http.StatusOK, "login")
}
