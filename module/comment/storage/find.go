package commentstorage

import (
	"context"
	commentmodel "gin_form/module/comment/model"
)

func (store *sqlStore) FindComment(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*commentmodel.Comment, error) {
	var data commentmodel.Comment

	if err := store.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
