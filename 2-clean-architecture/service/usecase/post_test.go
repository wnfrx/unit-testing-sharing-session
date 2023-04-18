package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/models"
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service/repository/mocks"
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service/usecase"
)

func TestAddPost(t *testing.T) {
	type testCase struct {
		name          string
		wantError     bool
		input         models.Post
		expectedError error
		mockFunc      func(postMock *mocks.MockPostRepository)
	}

	var testTable = []testCase{
		{
			name:      "empty post title",
			wantError: true,
			input: models.Post{
				Title:    "",
				Content:  "",
				Author:   "",
				PostedAt: time.Now(),
			},
			expectedError: models.ErrPostTitleEmpty,
		},
		{
			name:      "author empty",
			wantError: true,
			input: models.Post{
				Title:    "this is title",
				Content:  "",
				Author:   "",
				PostedAt: time.Now(),
			},
			expectedError: errors.New("post author is empty"),
		},
		{
			name:      "error add post",
			wantError: true,
			input: models.Post{
				Title:    "this is title",
				Content:  "",
				Author:   "admin",
				PostedAt: time.Now(),
			},
			expectedError: errors.New("this is an unexpected error"),
			mockFunc: func(postMock *mocks.MockPostRepository) {
				postMock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("this is an unexpected error"))
			},
		},
		{
			name:      "success",
			wantError: false,
			input: models.Post{
				Title:    "this is title",
				Content:  "",
				Author:   "admin",
				PostedAt: time.Now(),
			},
			mockFunc: func(postMock *mocks.MockPostRepository) {
				postMock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(int64(123), nil)
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			// Init mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			postRepo := mocks.NewMockPostRepository(mockCtrl)
			postUsecase := usecase.NewPostUsecase(postRepo)

			if tc.mockFunc != nil {
				tc.mockFunc(postRepo)
			}

			id, errx := postUsecase.Add(context.Background(), tc.input)

			if tc.wantError {
				assert.Equal(t, int64(0), id)
				assert.NotNil(t, errx)
				assert.EqualError(t, errx, tc.expectedError.Error())
			} else {
				assert.Greater(t, id, int64(0))
				assert.Nil(t, errx)
			}
		})
	}
}

func TestGetPost(t *testing.T) {
	// TODO: implement me
}

func TestUpdatePost(t *testing.T) {
	type testCase struct {
		name          string
		wantError     bool
		input         models.Post
		expectedError error
		mockFunc      func(postMock *mocks.MockPostRepository)
	}

	var testTable = []testCase{
		{
			name:      "invalid id",
			wantError: true,
			input: models.Post{
				ID:       0,
				Title:    "",
				Content:  "",
				Author:   "",
				PostedAt: time.Now(),
			},
			expectedError: models.ErrInvalidId,
		},
		{
			name:      "error add post",
			wantError: true,
			input: models.Post{
				ID:       123,
				Title:    "this is title",
				Content:  "",
				Author:   "",
				PostedAt: time.Now(),
			},
			expectedError: errors.New("this is an unexpected error"),
			mockFunc: func(postMock *mocks.MockPostRepository) {
				postMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("this is an unexpected error"))
			},
		},
		{
			name:      "success",
			wantError: false,
			input: models.Post{
				ID:       123,
				Title:    "this is title",
				Content:  "",
				Author:   "",
				PostedAt: time.Now(),
			},
			mockFunc: func(postMock *mocks.MockPostRepository) {
				postMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			// Init mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			postRepo := mocks.NewMockPostRepository(mockCtrl)
			postUsecase := usecase.NewPostUsecase(postRepo)

			// Run flow mock
			if tc.mockFunc != nil {
				tc.mockFunc(postRepo)
			}

			errx := postUsecase.Update(context.Background(), tc.input)

			if tc.wantError {
				assert.NotNil(t, errx)
				assert.EqualError(t, errx, tc.expectedError.Error())
			} else {
				assert.Nil(t, errx)
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	// TODO: implement me
}

func TestGetByAuthor(t *testing.T) {
	// test table

	// loop per test case
}
