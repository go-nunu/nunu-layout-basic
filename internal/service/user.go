package service

import (
	"github.com/go-nunu/nunu-layout-base/internal/dao"
	"github.com/go-nunu/nunu-layout-base/internal/model"
)

type UserService struct {
	*Service
	userDao *dao.UserDao
}

func NewUserService(service *Service, userDao *dao.UserDao) *UserService {
	return &UserService{
		Service: service,
		userDao: userDao,
	}
}

func (s *UserService) GetUserById(id int64) (*model.User, error) {
	return s.userDao.FirstById(id)
}
