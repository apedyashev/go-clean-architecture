package controller

import (
	"fmt"
	"net/http"
	"posts/usecase/usecase"
)

type postController struct {
	postUsecase usecase.Post
}

// interface for the Post controller
type PostController interface {
	GetPosts(ctx Context) error
}

// By using the Post interface we specify that this controller must implement this interface
func NewPostController(pu usecase.Post) PostController {
	return &postController{
		postUsecase: pu,
	}
}

func (pc postController) GetPosts(ctx Context) error {
	pp, err := pc.postUsecase.FindAll()
	if err != nil {
		fmt.Println("get posts controller err", err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, pp)
}
