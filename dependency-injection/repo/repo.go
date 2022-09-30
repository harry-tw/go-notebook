package repo

import (
	"go-notebook/dependency-injection/dao"
)

type OrderRepo struct {
	dao *dao.OrderDao
}

func NewOrderRepo(dao *dao.OrderDao) (*OrderRepo, error) {
	return &OrderRepo{
		dao: dao,
	}, nil
}

type UserRepo struct {
	dao *dao.UserDao
}

func NewUserRepo(dao *dao.UserDao) (*UserRepo, error) {
	return &UserRepo{
		dao: dao,
	}, nil
}
