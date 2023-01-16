package commentstorage

import (
	"context"
	commentmodel "gin_form/module/comment/model"
)

func (store *sqlStore) UpdateComment(
	ctx context.Context,
	cond map[string]interface{},
	data *commentmodel.CommentUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
