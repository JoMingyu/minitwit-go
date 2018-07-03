package main

import (
	"os"

	"minitwit-go/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	return port
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	controllers.Setup(e.Router(), e.Group(""))

	err := e.Start(":" + port())
	if err != nil {
		panic(err)
	}
}
