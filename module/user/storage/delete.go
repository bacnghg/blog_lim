package userstorage

import (
	"context"
	usermodel "gin_form/module/user/model"
)

func (store *sqlStore) DeleteUser(ctx context.Context, id int) error {
	if err := store.db.Model(usermodel.User{}).Where("id = ? ", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
