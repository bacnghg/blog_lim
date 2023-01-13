package userbusiness

import (
	"context"
	usermodel "gin_form/module/user/model"
)

type CreateUserStore interface {
	InsertUser(ctx context.Context, data *usermodel.UserCreate) error
}

func NewCreateUser(store CreateUserStore) *createUserBusiness {
	return &createUserBusiness{store: store}
}

type createUserBusiness struct {
	store CreateUserStore
}

func (business *createUserBusiness) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := business.store.InsertUser(ctx, data); err != nil {
		return err
	}
	return nil
}
