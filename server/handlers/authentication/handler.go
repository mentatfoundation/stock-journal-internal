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

func (a *Handler) Logout(c echo.Context) error {
	a.Logger.Info("logout")
	return c.String(http.StatusOK, "logout")
}

func (a *Handler) User(c echo.Context) error {
	a.Logger.Info(c.Param("id"))
	return c.String(http.StatusOK, c.Param("id"))
}
