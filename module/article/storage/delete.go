package articlestorage

import (
	"context"
	articlemodel "gin_form/module/article/model"
)

func (store *sqlStore) DeleteArticle(ctx context.Context, id int) error {
	if err := store.db.Model(articlemodel.Article{}).Where("id = ? ", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
