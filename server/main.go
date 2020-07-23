package main

import (
	"fmt"

	"github.com/jak103/uno/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Database db.UnoDB

func main() {
	fmt.Println("USU - UNO v0.0.0")

	Database = db.GetDb()

	// New Echo server
	e := echo.New()

	// Setup middleware
	//e.File("/", "/client/dist/index.html")

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "/client/dist/",
		HTML5: true,
	}))

	// Setup routes
	setupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
