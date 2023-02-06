package commentstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewDbStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
