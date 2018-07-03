package controllers

import (
	"net/http"
	"time"

	"minitwit-go/config"
	"minitwit-go/models"

	jwt "github.com/dgrijalva/jwt-go"
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

func Login(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusInternalServerError, "")
	} else {
		if count, _ := config.UserCol.Find(bson.M{"username": u.Username, "pw": u.Pw}).Count(); count == 0 {
			return c.String(http.StatusUnauthorized, "")
		} else {
			token := jwt.New(jwt.SigningMethodHS256)

			claims := token.Claims.(jwt.MapClaims)
			claims["username"] = u.Username
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			t, _ := token.SignedString([]byte("secret"))

			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}
	}
}
