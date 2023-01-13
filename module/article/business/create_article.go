package articlebusiness

import (
	"context"
	articlemodel "gin_form/module/article/model"
)

type CreateArticleStore interface {
	InsertArticle(ctx context.Context, data *articlemodel.ArticleCreate) error
}

func NewCreateArticle(store CreateArticleStore) *createArticleBusiness {
	return &createArticleBusiness{store: store}
}

type createArticleBusiness struct {
	store CreateArticleStore
}

func (business *createArticleBusiness) CreateArticle(ctx context.Context, data *articlemodel.ArticleCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := business.store.InsertArticle(ctx, data); err != nil {
		return err
	}
	return nil
}
