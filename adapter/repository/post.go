package repository

import (
	"posts/domain/model"
	"posts/usecase/repository"
)

type InMemoryDB interface {
	Insert(d any) error
	FindAll() ([]model.Post, error)
}

type postRepository struct {
	// custom InMemoryDB for simplicity
	db InMemoryDB
}

// By using the repository.PostRepository interface we specify that this repo
// must implement this interface
func NewPostRepository(db InMemoryDB) repository.PostRepository {
	return &postRepository{db}
}

func (pr postRepository) CreateOne(p model.Post) error {
	return pr.db.Insert(p)
}

func (pr postRepository) FindAll() ([]model.Post, error) {
	return pr.db.FindAll()
}
