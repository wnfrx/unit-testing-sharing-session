package app

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	db     *sqlx.DB
	router *gin.Engine
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() (err error) {
	if a.router == nil {
		return errors.New("router is not initialized yet")
	}

	if os.Getenv("APP_PORT") == "" {
		if err = a.router.Run(); err != nil {
			return err
		}
	} else {
		if err = a.router.Run(":" + os.Getenv("APP_PORT")); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Stop() (err error) {
	return
}
