package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/models"
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service"
)

type postUsecase struct {
	repo service.PostRepository
}

func NewPostUsecase(repo service.PostRepository) service.PostUsecase {
	return &postUsecase{repo}
}

func (u *postUsecase) Add(ctx context.Context, args models.Post) (id int64, err error) {
	if args.Title == "" {
		err = models.ErrPostTitleEmpty
		return
	}

	if args.Author == "" {
		err = errors.New("post author is empty")
		return
	}

	id, err = u.repo.Add(ctx, args)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (u *postUsecase) Get(ctx context.Context, id int64) (result models.Post, err error) {
	if id <= 0 {
		err = models.ErrInvalidId
		return
	}

	result, err = u.repo.Get(ctx, id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (u *postUsecase) Update(ctx context.Context, args models.Post) (err error) {
	if args.ID <= 0 {
		return models.ErrInvalidId
	}

	err = u.repo.Update(ctx, args)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (u *postUsecase) Delete(ctx context.Context, id int64) (err error) {
	if id <= 0 {
		return models.ErrInvalidId
	}

	err = u.repo.Delete(ctx, id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (u *postUsecase) GetByAuthor(ctx context.Context, author string) (result models.Post, err error) {

	return
}
