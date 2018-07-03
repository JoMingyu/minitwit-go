package controllers

import (
	"github.com/labstack/echo"
)

// Setup sets up all controllers.
func Setup(router *echo.Router) {
	router.Add(echo.POST, "/signup", Signup)
	router.Add(echo.POST, "/login", Login)
}

// vi:syntax=go
