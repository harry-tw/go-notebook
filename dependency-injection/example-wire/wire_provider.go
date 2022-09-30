package main

import (
	"github.com/google/wire"
	"go-notebook/dependency-injection/cache"
	"go-notebook/dependency-injection/config"
	"go-notebook/dependency-injection/dao"
	"go-notebook/dependency-injection/repo"
	"go-notebook/dependency-injection/service"
)

var cacheSet = wire.NewSet(provideCacheLRU)
var daoSet = wire.NewSet(provideUserDao, provideOrderDao)
var repoSet = wire.NewSet(repo.NewUserRepo, repo.NewOrderRepo)
var serviceSet = wire.NewSet(service.NewUserService, service.NewOrderService, service.NewCacheService)
var applicationSet = wire.NewSet(newApplication, config.NewConfig, cacheSet, daoSet, repoSet, serviceSet)

func provideCacheLRU(cfg *config.Config) (*cache.LRU, error) {
	return cache.NewLRU(cfg.CacheSize)
}

func provideUserDao(cfg *config.Config) (*dao.UserDao, error) {
	return dao.NewUserDao(cfg.DbUrl)
}

func provideOrderDao(cfg *config.Config) (*dao.OrderDao, error) {
	return dao.NewOrderDao(cfg.DbUrl)
}
