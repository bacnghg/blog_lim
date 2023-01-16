package commentbusiness

import (
	"context"
	commentmodel "gin_form/module/comment/model"
)

type CreateCommentStore interface {
	InsertComment(ctx context.Context, data *commentmodel.CommentCreate) error
}

func NewCreateComment(store CreateCommentStore) *createCommentBusiness {
	return &createCommentBusiness{store: store}
}

type createCommentBusiness struct {
	store CreateCommentStore
}

func (business *createCommentBusiness) CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := business.store.InsertComment(ctx, data); err != nil {
		return err
	}
	return nil
}
