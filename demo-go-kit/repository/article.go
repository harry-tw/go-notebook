package repository

import (
	"context"
	"go-notebook/demo-go-kit/model"
	"go-notebook/demo-go-kit/model/query"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(ctx context.Context, article *model.Article) error
	Get(ctx context.Context, id string) (*model.Article, error)
	Update(ctx context.Context, article *model.Article) (gen.ResultInfo, error)
	Delete(ctx context.Context, id string) (gen.ResultInfo, error)
}

type articleRepo struct {
	dao *query.Query
}

func (a *articleRepo) Create(ctx context.Context, article *model.Article) error {
	return a.dao.WithContext(ctx).Article.Create(article)
}

func (a *articleRepo) Get(ctx context.Context, id string) (*model.Article, error) {
	queryArticle := query.Article
	return a.dao.WithContext(ctx).Article.Where(queryArticle.ID.Eq(id)).First()
}

func (a *articleRepo) Update(ctx context.Context, article *model.Article) (gen.ResultInfo, error) {
	queryArticle := query.Article
	return a.dao.WithContext(ctx).Article.Where(queryArticle.ID.Eq(article.ID)).Updates(article)
}

func (a *articleRepo) Delete(ctx context.Context, id string) (gen.ResultInfo, error) {
	queryArticle := query.Article
	return a.dao.WithContext(ctx).Article.Where(queryArticle.ID.Eq(id)).Delete()
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	dao := query.Use(db)
	return &articleRepo{
		dao: dao,
	}
}
