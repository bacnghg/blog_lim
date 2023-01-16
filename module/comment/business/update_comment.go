package commentbusiness

import (
	"context"
	commentmodel "gin_form/module/comment/model"
)

type UpdateCommentStore interface {
	FindComment(
		ctx context.Context,
		cond map[string]interface{},
		morekeys ...string,
	) (*commentmodel.Comment, error)

	UpdateComment(
		ctx context.Context,
		cond map[string]interface{},
		data *commentmodel.CommentUpdate,
	) error
}

func UpdateCommentBusiness(store UpdateCommentStore) *updateCommentBusiness {
	return &updateCommentBusiness{store: store}
}

type updateCommentBusiness struct {
	store UpdateCommentStore
}

func (business *updateCommentBusiness) UpdateCommentById(ctx context.Context, id int, data *commentmodel.CommentUpdate) error {
	_, err := business.store.FindComment(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if err := business.store.UpdateComment(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
