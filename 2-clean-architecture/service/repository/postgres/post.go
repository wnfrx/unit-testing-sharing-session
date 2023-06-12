package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/models"
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service"
)

type postRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) service.PostRepository {
	return &postRepository{db}
}

func (u *postRepository) Add(ctx context.Context, args models.Post) (id int64, err error) {
	fmt.Println("test add")

	return
}

func (u *postRepository) Get(ctx context.Context, id int64) (result models.Post, err error) {

	return
}

func (u *postRepository) Update(ctx context.Context, args models.Post) (err error) {

	return
}

func (u *postRepository) Delete(ctx context.Context, id int64) (err error) {

	return
}

func (u *postRepository) GetByAuthor(ctx context.Context, author string) (result models.Post, err error) {

	return
}
