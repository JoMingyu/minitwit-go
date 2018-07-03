package controllers

import (
	"net/http"
	"time"

	"minitwit-go/config"
	"minitwit-go/models"

	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusInternalServerError, "")
	} else {
		if count, _ := config.UserCol.Find(bson.M{"username": u.Username}).Count(); count != 0 {
			// username이 이미 존재하는 경우
			return c.String(http.StatusConflict, "")
		} else {
			config.UserCol.Insert(u)

			return c.String(http.StatusOK, "")
		}
	}
}

