package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-notebook/demo-go-kit/repository"
	"go-notebook/demo-go-kit/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	logger := logrus.New()
	dsn := "host=localhost user=postgres password=password dbname=demo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panicf("failed to open db: %v", err)
	}
	repo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(logger, repo)
	ctx := context.Background()
	err = articleService.Create(ctx, "Hello", "World")
	if err != nil {
		log.Panicf("failed to create article: %v", err)
	}
}
