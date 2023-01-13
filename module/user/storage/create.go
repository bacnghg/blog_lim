package userstorage

import (
	"context"
	usermodel "gin_form/module/user/model"
)

func (store *sqlStore) InsertUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := store.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
