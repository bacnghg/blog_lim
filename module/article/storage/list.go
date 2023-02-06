package articlestorage

import (
	"context"
	"gin_form/common"
	articlemodel "gin_form/module/article/model"
)

func (store sqlStore) ListArticle(
	ctx context.Context,
	filter *articlemodel.FilterArticle,
	paging *common.Paging,
	morekeys ...string,
) ([]articlemodel.Article, error) {

	offset := (paging.Page - 1) * paging.Limit

	var result []articlemodel.Article

	db := store.db
	if v := filter.ArticleId; v > 0 {
		db = db.Where("id = ?", v)
	}
	if v := filter.Title; v != "" {
		db = db.Where("title LIKE  ?", "%"+v+"%")
	}
	if v := filter.Description; v != "" {
		db = db.Where("description LIKE  ?", "%"+v+"%")
	}
	if v := filter.UsereId; v > 0 {
		db = db.Where("user_id = ?", v)
	}
	if v := filter.Category; v > 0 {
		db = db.Where("category = ?", v)
	}
	if err := db.Table(articlemodel.Article{}.TableName()).
		Count(&paging.Total).
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
