package userstorage

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	ID      int
	DeletedAt gorm.DeletedAt
}
  
func (store *sqlStore) DeleteUser(ctx context.Context, id int) error {
	if err := store.db.Where("id = ? ", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
