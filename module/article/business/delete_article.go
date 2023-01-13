package articlebusiness

import (
	"context"
	articlemodel "gin_form/module/article/model"
)

type DeleteArticleStore interface {
	FindArticle(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*articlemodel.Article, error)
	DeleteArticle(ctx context.Context, id int) error
}

func NewDeleteArticle(store DeleteArticleStore) *deleteArticleBusiness {
	return &deleteArticleBusiness{store: store}
}

type deleteArticleBusiness struct {
	store DeleteArticleStore
}

func (business *deleteArticleBusiness) DeleteArticle(ctx context.Context, id int) error {

	_, err := business.store.FindArticle(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := business.store.DeleteArticle(ctx, id); err != nil {
		return err
	}
	return nil
}
