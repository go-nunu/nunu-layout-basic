package repository

import (
	"github.com/go-nunu/nunu-layout-basic/internal/model"
)

type UserRepository interface {
	FirstById(id int64) (*model.User, error)
}
type userRepository struct {
	*Repository
}

func NewUserRepository(repository *Repository) UserRepository {
	return &userRepository{
		Repository: repository,
	}
}

func (r *userRepository) FirstById(id int64) (*model.User, error) {
	var user model.User
	// TODO: query db
	return &user, nil
}
