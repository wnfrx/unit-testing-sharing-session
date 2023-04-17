package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service"
)

type restHandler struct {
	router *gin.Engine

	postUsecase service.PostUsecase
}

func NewRestHandler(
	r *gin.Engine,
	postUsecase service.PostUsecase,
) *restHandler {
	return &restHandler{
		router:      r,
		postUsecase: postUsecase,
	}
}

func (h *restHandler) RegisterRoutes() {
	h.router.GET("/", h.testWelcome)

	h.router.POST("/post", h.addPost)
	h.router.GET("/post/:id", h.getPost)
	h.router.PATCH("/post/:id", h.updatePost)
	h.router.DELETE("/post/:id", h.deletePost)
}
