package commentbusiness

import (
	"context"
	"gin_form/common"
	commentmodel "gin_form/module/comment/model"
)

type ListCommentStore interface {
	ListComment(
		ctx context.Context,
		filter *commentmodel.Filter,
		paging *common.Paging,
		morekeys ...string,
	) ([]commentmodel.Comment, error)
}

type listCommentStruct struct {
	store ListCommentStore
}

func ListCommentBusiness(store ListCommentStore) *listCommentStruct {
	return &listCommentStruct{store: store}
}

func (business *listCommentStruct) ListComment(
	ctx context.Context,
	filter *commentmodel.Filter,
	paging *common.Paging,
) ([]commentmodel.Comment, error) {
	result, err := business.store.ListComment(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
