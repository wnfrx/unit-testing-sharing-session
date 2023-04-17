package app

import "github.com/gin-gonic/gin"

func (a *App) InitRouter() (err error) {
	if a.router != nil {
		return nil
	}

	a.router = gin.Default()

	return nil
}
