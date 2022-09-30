package main

import (
	"go-notebook/dependency-injection/cache"
	"go-notebook/dependency-injection/config"
	"go-notebook/dependency-injection/dao"
)

func provideUserDao(cfg *config.Config) (*dao.UserDao, error) {
	return dao.NewUserDao(cfg.DbUrl)
}

func provideOrderDao(cfg *config.Config) (*dao.OrderDao, error) {
	return dao.NewOrderDao(cfg.DbUrl)
}

func provideCacheLRU(cfg *config.Config) (*cache.LRU, error) {
	return cache.NewLRU(cfg.CacheSize)
}
