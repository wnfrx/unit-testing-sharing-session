package service

import (
	"context"

	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/models"
)

type PostUsecase interface {
	Add(ctx context.Context, args models.Post) (id int64, err error)
	Get(ctx context.Context, id int64) (result models.Post, err error)
	Update(ctx context.Context, args models.Post) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
