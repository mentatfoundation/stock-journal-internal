package server

import (
	authHandlers "mentatfoundation/stock-journal/server/handlers/authentication"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/api/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "client/build",
		HTML5: true,
	}))

	authHandler := authHandlers.AuthHandler{}

	e.GET("/api/auth/login", authHandler.Login)

	e.Logger.Fatal(e.Start(":5000"))
}
