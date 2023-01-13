package articlestorage

import (
	"context"
	articlemodel "gin_form/module/article/model"
)

func (store *sqlStore) InsertArticle(ctx context.Context, data *articlemodel.ArticleCreate) error {
	if err := store.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
