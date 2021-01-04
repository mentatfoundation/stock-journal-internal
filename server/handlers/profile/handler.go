package profile

import (
	"mentatfoundation/stock-journal/server/logger"
	authservice "mentatfoundation/stock-journal/server/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	AuthService authservice.AuthService
	Logger      logger.Logger
}

func NewProfileHandler(l logger.Logger, as authservice.AuthService) *Handler {
	return &Handler{
		AuthService: as,
		Logger:      l,
	}
}

func (p *Handler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "profile")
}
