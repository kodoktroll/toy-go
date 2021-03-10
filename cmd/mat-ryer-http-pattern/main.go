package main

import (
	"log"
	"net/http"

	"github.com/kodoktroll/toy-go/pkg/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	db *db.Database
}

var (
	RedisAddr = "localhost:6379"
)

func (s *server) handleLeaderboardsGet() func(echo.Context) error {
	return func(c echo.Context) error {
		leaderboard, err := s.db.GetLeaderboard()
		if err != nil {
			if err == db.ErrNil {
				return c.NoContent(http.StatusNotFound)
			}
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, leaderboard)
	}
}

func newServer() *server {
	s := &server{}
	return s
}

func exampleClient2() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s := newServer()

	database, err := db.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalf(err.Error())
	}

	s.db = database
	e.GET("/leaderboards", s.handleLeaderboardsGet())
}
