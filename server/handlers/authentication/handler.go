package authentication

import (
	"mentatfoundation/stock-journal/server/logger"
	"net/http"

	authService "mentatfoundation/stock-journal/server/services"

	"github.com/labstack/echo"
)

type Handler struct {
	AuthService *authService.AuthService
	Logger      *logger.Logger
}

func New(as *authService.AuthService, l *logger.Logger) *Handler {
	return &Handler{
		AuthService: as,
		Logger:      l,
	}
}

func (a *Handler) Login(c echo.Context) error {
	a.Logger.Info("hello")
	return c.String(http.StatusOK, "login")
}
