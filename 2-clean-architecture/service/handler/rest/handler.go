package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/models"
)

type responseBody struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (h *restHandler) testWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, responseBody{
		Status:  http.StatusOK,
		Message: "Welcome!",
	})
}

func (h *restHandler) addPost(c *gin.Context) {
	var args models.Post

	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, responseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

	id, err := h.postUsecase.Add(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseBody{
			Status:  http.StatusInternalServerError,
			Message: "Oops! Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, responseBody{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]interface{}{
			"id": id,
		},
	})
}

func (h *restHandler) getPost(c *gin.Context) {
	var (
		id  int64
		err error
	)

	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid post ID",
		})
		return
	}

	var result models.Post

	result, err = h.postUsecase.Get(c, id)
	if err != nil {
		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, responseBody{
				Status:  http.StatusNotFound,
				Message: "Post not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, responseBody{
			Status:  http.StatusInternalServerError,
			Message: "Oops! Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, responseBody{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func (h *restHandler) updatePost(c *gin.Context) {
	var (
		id  int64
		err error
	)

	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid post ID",
		})
		return
	}

	var args models.Post

	if err = c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, responseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

	args.ID = id

	err = h.postUsecase.Update(c, args)
	if err != nil {
		if err == models.ErrInvalidId {
			c.JSON(http.StatusBadRequest, responseBody{
				Status:  http.StatusBadRequest,
				Message: "Invalid post ID",
			})
			return
		}

		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, responseBody{
				Status:  http.StatusNotFound,
				Message: "Post not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, responseBody{
			Status:  http.StatusInternalServerError,
			Message: "Oops! Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, responseBody{
		Status:  http.StatusOK,
		Message: "Success",
	})
}

func (h *restHandler) deletePost(c *gin.Context) {
	var (
		id  int64
		err error
	)

	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid post ID",
		})
		return
	}

	err = h.postUsecase.Delete(c, id)
	if err != nil {
		if err == models.ErrInvalidId {
			c.JSON(http.StatusBadRequest, responseBody{
				Status:  http.StatusBadRequest,
				Message: "Invalid post ID",
			})
			return
		}

		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, responseBody{
				Status:  http.StatusNotFound,
				Message: "Post not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, responseBody{
			Status:  http.StatusInternalServerError,
			Message: "Oops! Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, responseBody{
		Status:  http.StatusOK,
		Message: "Success",
	})
}
