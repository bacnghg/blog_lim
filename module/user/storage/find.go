package userstorage

import (
	"context"
	usermodel "gin_form/module/user/model"
)

func (store *sqlStore) FindUser(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	var data usermodel.User

	if err := store.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
