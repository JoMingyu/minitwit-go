package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"minitwit-go/config"
	"minitwit-go/controllers"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	return port
}

func main() {
	e := echo.New()

	config.Setup(e)
	controllers.Setup(e.Router())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Nothing to see here.")
	})

	err := e.Start(":" + getPort())
	if err != nil {
		panic(err)
	}
}

// vi:syntax=go
