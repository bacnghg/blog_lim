package commentbusiness

import (
	"context"
	commentmodel "gin_form/module/comment/model"
)

type DeleteCommentStore interface {
	FindComment(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*commentmodel.Comment, error)
	DeleteComment(ctx context.Context, id int) error
}

func NewDeleteComment(store DeleteCommentStore) *deleteCommentBusiness {
	return &deleteCommentBusiness{store: store}
}

type deleteCommentBusiness struct {
	store DeleteCommentStore
}

func (business *deleteCommentBusiness) DeleteComment(ctx context.Context, id int) error {

	_, err := business.store.FindComment(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := business.store.DeleteComment(ctx, id); err != nil {
		return err
	}
	return nil
}
