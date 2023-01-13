package userbusiness

import (
	"context"
	usermodel "gin_form/module/user/model"
)

type FindUserStore interface {
	FindUser(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

func NewFindUserBusiness(store FindUserStore) *findUserBusiness {
	return &findUserBusiness{store: store}
}

type findUserBusiness struct {
	store FindUserStore
}

func (business *findUserBusiness) FindUser(ctx context.Context, id int) (*usermodel.User, error) {
	data, err := business.store.FindUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return data, nil
}
