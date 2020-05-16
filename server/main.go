package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("USU - UNO v0.0.0")

	// New Echo server
	e := echo.New()

	// Setup middleware
	// TODO: figure out how to serve the app html for the frontend
	e.File("/", "static/index.html")
	//e.Use(middleware.Static("/static")) TODO: Figure out how to serve the statis js files
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	// Setup routes
	setupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
