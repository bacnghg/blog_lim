package articlebusiness

import (
	"context"
	articlemodel "gin_form/module/article/model"
)

type UpdateArticleStore interface {
	FindArticle(
		ctx context.Context,
		cond map[string]interface{},
		morekeys ...string,
	) (*articlemodel.Article, error)

	UpdateArticle(
		ctx context.Context,
		cond map[string]interface{},
		data *articlemodel.ArticleUpdate,
	) error
}

func UpdateArticleBusiness(store UpdateArticleStore) *updateArticleBusiness {
	return &updateArticleBusiness{store: store}
}

type updateArticleBusiness struct {
	store UpdateArticleStore
}

func (business *updateArticleBusiness) UpdateArticleById(ctx context.Context, id int, data *articlemodel.ArticleUpdate) error {
	_, err := business.store.FindArticle(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if err := business.store.UpdateArticle(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
