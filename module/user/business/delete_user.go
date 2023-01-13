package userbusiness

import (
	"context"
	usermodel "gin_form/module/user/model"
)

type DeleteUserStore interface {
	FindUser(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
	DeleteUser(ctx context.Context, id int) error
}

func NewDeleteUser(store DeleteUserStore) *deleteUserBusiness {
	return &deleteUserBusiness{store: store}
}

type deleteUserBusiness struct {
	store DeleteUserStore
}

func (business *deleteUserBusiness) DeleteUser(ctx context.Context, id int) error {

	_, err := business.store.FindUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := business.store.DeleteUser(ctx, id); err != nil {
		return err
	}
	return nil
}
