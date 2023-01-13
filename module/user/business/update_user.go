package userbusiness

import (
	"context"
	usermodel "gin_form/module/user/model"
)

type UpdateUserStore interface {
	FindUser(
		ctx context.Context,
		cond map[string]interface{},
		morekeys ...string,
	) (*usermodel.User, error)

	UpdateUser(
		ctx context.Context,
		cond map[string]interface{},
		data *usermodel.UserUpdate,
	) error
}

func UpdateUserBusiness(store UpdateUserStore) *updateUserBusiness {
	return &updateUserBusiness{store: store}
}

type updateUserBusiness struct {
	store UpdateUserStore
}

func (business *updateUserBusiness) UpdateUserById(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	_, err := business.store.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if err := business.store.UpdateUser(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
