package commentbusiness

import (
	"context"
	commentmodel "gin_form/module/comment/model"
)

type FindCommentStore interface {
	FindComment(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*commentmodel.Comment, error)
}

func NewFindCommentBusiness(store FindCommentStore) *findCommentBusiness {
	return &findCommentBusiness{store: store}
}

type findCommentBusiness struct {
	store FindCommentStore
}

func (business *findCommentBusiness) FindComment(ctx context.Context, id int) (*commentmodel.Comment, error) {
	data, err := business.store.FindComment(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return data, nil
}
