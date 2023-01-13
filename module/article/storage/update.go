package articlestorage

import (
	"context"
	articlemodel "gin_form/module/article/model"
)

func (store *sqlStore) UpdateArticle(
	ctx context.Context,
	cond map[string]interface{},
	data *articlemodel.ArticleUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
