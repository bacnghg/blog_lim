package userstorage

import (
	"context"
	"gin_form/common"
	usermodel "gin_form/module/user/model"
)

func (store sqlStore) ListUser(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]usermodel.User, error) {

	offset := (paging.Page - 1) * paging.Limit

	var result []usermodel.User

	db := store.db
	if v := filter.CategoryId; v > 0 {
		db = db.Where("category_id = ?", v)
	}
	if err := db.Table(usermodel.User{}.TableName()).
		Count(&paging.Total).
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
