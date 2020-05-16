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
	e.File("/", "static/index.html")	
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		HTML5: true,
	  }))

	// Setup routes
	setupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
