package userstorage

import (
	"context"
	"gin_form/common"
	"math"
	"fmt"
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
		Where("deleted_at IS NULL").
		Count(&paging.Total).
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {			
		return nil, err
	}
	lastPage := float64(paging.Total) / float64(paging.Limit)
	fmt.Println("lastPage", lastPage)
	paging.TotalPage = math.Ceil(lastPage)
	
	return result, nil
}
