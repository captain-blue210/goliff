package main

import (
	"github.com/captainblue210/goliff/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	routes.Init(e)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
