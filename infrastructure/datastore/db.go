package datastore

import (
	"fmt"
	"posts/adapter/repository"
	"posts/domain/model"
)

// type InMemoryDB interface {
// 	Insert(d any) error
// 	FindAll() ([]model.Post, error)
// }

type InMemoryDB struct {
	items []model.Post
}

func NewDB() repository.InMemoryDB {
	return &InMemoryDB{
		items: []model.Post{
			{ID: 1, Title: "post 1", Images: []string{"uds-1d-cd"}},
			{ID: 2, Title: "post 2", Images: []string{"fd-2d-hds", "fd-1cds-ls"}},
			{ID: 3, Title: "post 3", Images: []string{"fd-3d-lsd"}},
			{ID: 4, Title: "post 4", Images: []string{"cf-4d-jdcdU"}},
		},
	}
}

func (db InMemoryDB) Insert(d any) error {
	return nil
}

func (db InMemoryDB) FindAll() ([]model.Post, error) {
	fmt.Println("InMemoryDB find all")
	return db.items, nil
}
