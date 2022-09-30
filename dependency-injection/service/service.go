package service

import (
	"go-notebook/dependency-injection/cache"
	"go-notebook/dependency-injection/repo"
)

type OrderService struct {
	orderRepo *repo.OrderRepo
}

func NewOrderService(orderRepo *repo.OrderRepo) (*OrderService, error) {
	return &OrderService{orderRepo: orderRepo}, nil
}

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) (*UserService, error) {
	return &UserService{userRepo: userRepo}, nil
}

type CacheService struct {
	cache *cache.LRU
}

func NewCacheService(cache *cache.LRU) (*CacheService, error) {
	return &CacheService{cache: cache}, nil
}
