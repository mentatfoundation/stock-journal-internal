package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"mentatfoundation/stock-journal/server/config"
	"os"
)

type App struct {
	Server *echo.Echo
	Config config.ConfigurationSettings
}

func New(c config.ConfigurationSettings) App {
	e := echo.New()
	return App{
		Server: e,
		Config: c,
	}
}

func (a App) Configure() {
	a.Config.DisplaySettings()
	a.ConfigureMiddleware()
	a.ConfigureRoutes()
}

func (a App) ConfigureMiddleware() {
	a.Server.Use(middleware.Recover())

	a.Server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "client/build",
		HTML5: true,
	}))

	if a.Config.IsDev() {
		a.Server.Use(middleware.CORS())
	} else {
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName("stock-journal-internal"),
			newrelic.ConfigLicense(a.Config.NewRelicKey),
		)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		a.Server.Use(nrecho.Middleware(app))
	}
}

func (a App) Run() {
	a.Configure()
	a.Server.Logger.Fatal(a.Server.Start(":" + a.Config.Port))
}
