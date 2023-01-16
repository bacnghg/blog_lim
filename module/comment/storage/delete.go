package commentstorage

import (
	"context"
	commentmodel "gin_form/module/comment/model"
)

func (store *sqlStore) DeleteComment(ctx context.Context, id int) error {
	if err := store.db.Model(commentmodel.Comment{}).Where("id = ? ", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
