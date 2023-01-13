package articlebusiness

import (
	"context"
	"gin_form/common"
	articlemodel "gin_form/module/article/model"
)

type ListArticleStore interface {
	ListArticle(
		ctx context.Context,
		filter *articlemodel.FilterArticle,
		paging *common.Paging,
		morekeys ...string,
	) ([]articlemodel.Article, error)
}

type listArticleStruct struct {
	store ListArticleStore
}

func ListArticleBusiness(store ListArticleStore) *listArticleStruct {
	return &listArticleStruct{store: store}
}

func (business *listArticleStruct) ListArticle(
	ctx context.Context,
	filter *articlemodel.FilterArticle,
	paging *common.Paging,
) ([]articlemodel.Article, error) {
	result, err := business.store.ListArticle(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
