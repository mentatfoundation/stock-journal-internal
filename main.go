package main

import (
	"mentatfoundation/stock-journal/server"
	"mentatfoundation/stock-journal/server/config"
)

func main() {
	configuration := config.New()
	app := server.New(configuration)
	app.Run()
}
