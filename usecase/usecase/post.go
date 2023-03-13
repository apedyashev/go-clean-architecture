/*
Business logic for posts
*/
package usecase

import (
	"fmt"
	"posts/domain/model"
	"posts/usecase/repository"
)

// interface of the use case
type Post interface {
	Create(model.Post) error
	FindAll() ([]model.Post, error)
}

type postUsecase struct {
	repo repository.PostRepository
}

func NewPostUseCase(postRepo repository.PostRepository) Post {
	return &postUsecase{
		repo: postRepo,
	}
}

func (pu postUsecase) Create(post model.Post) error {
	return pu.repo.CreateOne(post)
}

func (pu postUsecase) FindAll() ([]model.Post, error) {
	posts, err := pu.repo.FindAll()
	if err != nil {
		fmt.Println("Use case FindAll err", err)
		return nil, err
	}

	return posts, nil
}
