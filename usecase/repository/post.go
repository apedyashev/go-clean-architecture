/*
Repositroy in the Use case layer only declares interfaces
*/
package repository

import "posts/domain/model"

type PostRepository interface {
	CreateOne(model.Post) error
	FindAll() ([]model.Post, error)
}
