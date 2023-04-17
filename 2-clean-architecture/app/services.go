package app

import (
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service/handler/rest"
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service/repository/postgres"
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service/usecase"
)

func (a *App) InitServices() (err error) {
	postRepository := postgres.NewPostRepository(a.db)
	postUsecase := usecase.NewPostUsecase(postRepository)

	rh := rest.NewRestHandler(a.router, postUsecase)
	rh.RegisterRoutes()

	return
}
