package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go-notebook/demo-go-kit/model"
	"go-notebook/demo-go-kit/repository"
)

type ArticleService interface {
	Create(ctx context.Context, title, content string) error
	Get(ctx context.Context, id string) (*model.Article, error)
	Update(ctx context.Context, id string, title, content string) error
	Delete(ctx context.Context, id string) error
}

type articleSvc struct {
	logger *logrus.Logger
	repo   repository.ArticleRepository
}

func (a *articleSvc) Create(ctx context.Context, title, content string) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	article := &model.Article{
		Base: model.Base{
			ID: id.String(),
		},
		Title:   title,
		Content: content,
	}
	return a.repo.Create(ctx, article)
}

func (a *articleSvc) Get(ctx context.Context, id string) (*model.Article, error) {
	return a.repo.Get(ctx, id)
}

func (a *articleSvc) Update(ctx context.Context, id string, title, content string) error {
	article := &model.Article{
		Base: model.Base{
			ID: id,
		},
		Title:   title,
		Content: content,
	}
	_, err := a.repo.Update(ctx, article)
	return err
}

func (a *articleSvc) Delete(ctx context.Context, id string) error {
	_, err := a.repo.Delete(ctx, id)
	return err
}

func NewArticleService(logger *logrus.Logger, repository repository.ArticleRepository) ArticleService {
	return &articleSvc{
		logger: logger,
		repo:   repository,
	}
}
