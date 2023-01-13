package userbusiness

import (
	"context"
	"gin_form/common"
	usermodel "gin_form/module/user/model"
)

type ListUserStore interface {
	ListUser(
		ctx context.Context,
		filter *usermodel.Filter,
		paging *common.Paging,
		morekeys ...string,
	) ([]usermodel.User, error)
}

type listUserStruct struct {
	store ListUserStore
}

func ListUserBusiness(store ListUserStore) *listUserStruct {
	return &listUserStruct{store: store}
}

func (business *listUserStruct) ListUser(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
) ([]usermodel.User, error) {
	result, err := business.store.ListUser(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
