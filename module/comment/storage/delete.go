package commentstorage

import (
	"context"

	"gorm.io/gorm"
)

type Comment struct {
	ID int
	DeletedAt gorm.DeletedAt
}
func (store *sqlStore) DeleteComment(ctx context.Context, id int) error {
	if err := store.db.Where("id = ? ", id).Delete(&Comment{}).Error; err != nil {
		return err
	}
	return nil
}
