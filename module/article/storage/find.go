package articlestorage

import (
	"context"
	articlemodel "gin_form/module/article/model"
)

func (store *sqlStore) FindArticle(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*articlemodel.Article, error) {
	var data articlemodel.Article

	if err := store.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
