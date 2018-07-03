package controllers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Setup sets up all controllers.
func Setup(router *echo.Router, group *echo.Group) {
	router.Add(echo.POST, "/signup", signup)
	router.Add(echo.POST, "/login", login)

	group.Use(middleware.JWT([]byte("secret")))
	group.Add(echo.POST, "/:username/follow", follow)
	group.Add(echo.POST, "/:username/unfollow", unfollow)
}
