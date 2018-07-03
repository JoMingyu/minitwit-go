package controllers

import (
	"minitwit-go/config"
	"minitwit-go/models"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

func follow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	followerUsername := claims["username"]
	var follower models.User
	if err := config.UserCol.Find(bson.M{"username": followerUsername}).One(&follower); err != nil {
		return c.String(http.StatusForbidden, "")
	}

	followeeUsername := c.Param("username")
	var followee models.User
	if err := config.UserCol.Find(bson.M{"username": followeeUsername}).One(&followee); err != nil {
		return c.String(http.StatusNoContent, "")
	}

	if followerUsername == followeeUsername {
		return c.String(http.StatusBadRequest, "")
	}

	if count, _ := config.FollowCol.Find(bson.M{"follower": follower, "followee": followee}).Count(); count == 0 {
		config.FollowCol.Insert(models.Follow{follower, followee})
	}

	return c.String(http.StatusCreated, "")
}

func unfollow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	followerUsername := claims["username"]
	var follower models.User
	if err := config.UserCol.Find(bson.M{"username": followerUsername}).One(&follower); err != nil {
		return c.String(http.StatusForbidden, "")
	}

	followeeUsername := c.Param("username")
	var followee models.User
	if err := config.UserCol.Find(bson.M{"username": followeeUsername}).One(&followee); err != nil {
		return c.String(http.StatusNoContent, "")
	}

	if followerUsername == followeeUsername {
		return c.String(http.StatusBadRequest, "")
	}

	config.FollowCol.RemoveAll(bson.M{"follower": follower, "followee": followee})

	return c.String(http.StatusCreated, "")
}
