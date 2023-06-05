package dao

import (
	"github.com/go-nunu/nunu-layout-base/internal/model"
)

type UserDao struct {
	*Dao
}

func NewUserDao(dao *Dao) *UserDao {
	return &UserDao{
		Dao: dao,
	}
}

func (r *UserDao) FirstById(id int64) (*model.User, error) {
	var user model.User
	// TODO: query db
	return &user, nil
}
