package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	e.Logger.Fatal(e.Start(":5000"))
}
