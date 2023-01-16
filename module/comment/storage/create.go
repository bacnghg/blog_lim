package commentstorage

import (
	"context"
	commentmodel "gin_form/module/comment/model"
)

func (store *sqlStore) InsertComment(ctx context.Context, data *commentmodel.CommentCreate) error {
	if err := store.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
