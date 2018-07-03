package controllers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Setup sets up all controllers.
func Setup(router *echo.Router, group *echo.Group) {
	router.Add(echo.POST, "/signup", Signup)
	router.Add(echo.POST, "/login", Login)

	group.Use(middleware.JWT([]byte("secret")))
	group.Add(echo.POST, "/:username/follow", follow)
}

// vi:syntax=go
