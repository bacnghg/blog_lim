package commentmodel

import (
	"errors"
	"gin_form/common"
	"strings"
)

type Comment struct {
	common.SQLModel
	UserId      int `json:"user_id" gorm:"column:user_id"`
	ArticlesId  int `json:"articles_id" gorm:"column:articles_id"`
	Description string `json:"description" gorm:"column:description;"`
}

func (Comment) TableName() string {
	return "comments"
}

type CommentCreate struct {
	Id      int    `json:"id" gorm:"column:id;"`
	UserId      int    `json:"user_id" gorm:"column:user_id;"`
	ArticlesId  int `json:"articles_id" gorm:"column:articles_id;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (CommentCreate) TableName() string {
	return Comment{}.TableName()
}

func (res *CommentCreate) Validate() error {
	res.Description = strings.TrimSpace(res.Description)

	if len(res.Description) == 0 {
		return errors.New("Description can't be blank")
	}
	return nil
}

type CommentUpdate struct {
	Description   string `json:"description" gorm:"column:description;"`
}

func (CommentUpdate) TableName() string {
	return Comment{}.TableName()
}

func (res *CommentUpdate) Validate() error {
	res.Description = strings.TrimSpace(res.Description)

	if len(res.Description) == 0 {
		return errors.New("Description can't be blank")
	}
	return nil
}
