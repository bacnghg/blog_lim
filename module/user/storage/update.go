package userstorage

import (
	"context"
	usermodel "gin_form/module/user/model"
)

func (store *sqlStore) UpdateUser(
	ctx context.Context,
	cond map[string]interface{},
	data *usermodel.UserUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
