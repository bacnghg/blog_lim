package articlebusiness

import (
	"context"
	articlemodel "gin_form/module/article/model"
)

type FindArticleStore interface {
	FindArticle(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*articlemodel.Article, error)
}

func NewFindArticleBusiness(store FindArticleStore) *findArticleBusiness {
	return &findArticleBusiness{store: store}
}

type findArticleBusiness struct {
	store FindArticleStore
}

func (business *findArticleBusiness) FindArticle(ctx context.Context, id int) (*articlemodel.Article, error) {

	// data, err := business.store.FindArticle(ctx, map[string]interface{}{"id": id})
	data, err := business.store.FindArticle(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return data, nil
}
