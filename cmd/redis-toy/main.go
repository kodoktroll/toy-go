package main

import (
	"context"
	"log"
	"net/http"

	"github.com/kodoktroll/toy-go/pkg/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ctx = context.Background()

func main() {
	exampleClient()
}

var (
	RedisAddr = "localhost:6379"
)

func dbHandlePointsPost(database *db.Database) func(echo.Context) error {
	return func(c echo.Context) error {
		var userJson db.User
		if err := c.Bind(&userJson); err != nil {
			return err
		}
		log.Println(&userJson)
		err := database.SaveUser(&userJson)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, userJson)
	}
}

func dbHandlePointsGet(database *db.Database) func(echo.Context) error {
	return func(c echo.Context) error {
		username := c.Param("username")
		user, err := database.GetUser(username)
		if err != nil {
			if err == db.ErrNil {
				return c.NoContent(http.StatusNotFound)
			}
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, user)
	}
}

func dbHandleLeaderboardsGet(database *db.Database) func(echo.Context) error {
	return func(c echo.Context) error {
		leaderboard, err := database.GetLeaderboard()
		if err != nil {
			if err == db.ErrNil {
				return c.NoContent(http.StatusNotFound)
			}
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, leaderboard)
	}
}

func exampleClient() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database, err := db.NewDatabase(RedisAddr)

	e.POST("/points", dbHandlePointsPost(database))
	e.GET("/points/:username", dbHandlePointsGet(database))
	e.GET("/leaderboards", dbHandleLeaderboardsGet(database))

	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error)
	}
	e.Logger.Fatal(e.Start(":8000"))
}
