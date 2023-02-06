package commentstorage

import (
	"context"
	"gin_form/common"
	commentmodel "gin_form/module/comment/model"
)

func (store sqlStore) ListComment(
	ctx context.Context,
	filter *commentmodel.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]commentmodel.Comment, error) {

	offset := (paging.Page - 1) * paging.Limit

	var result []commentmodel.Comment

	db := store.db
	if v := filter.ArticleId; v > 0 {
		db = db.Where("articles_id = ?", v)
	}
	if v := filter.UserId; v > 0 {
		db = db.Where("user_id = ?", v)
	}
	if v := filter.Description; v != "" {
		db = db.Where("description like ?", "%"+v+"%")
	}
	if err := db.Table(commentmodel.Comment{}.TableName()).
		Count(&paging.Total).
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
