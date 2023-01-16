package commentmodel

import (
	"errors"
	"strings"
)

type Comment struct {
	CommentId   string `json:"comment_id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	UsereId     string `json:"user_id" gorm:"column:user_id"`
	Category    string `json:"category" gorm:"column:category"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   string `json:"deleted_at" gorm:"column:deleted_at"`
}

func (Comment) TableName() string {
	return "comments"
}

type CommentCreate struct {
	CommentId   string `json:"comment_id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	UsereId     string `json:"user_id" gorm:"column:user_id;"`
	Category    string `json:"category" gorm:"column:category"`
}

func (CommentCreate) TableName() string {
	return Comment{}.TableName()
}

func (res *CommentCreate) Validate() error {
	res.Title = strings.TrimSpace(res.Title)

	if len(res.Title) == 0 {
		return errors.New("Title can't be blank")
	}
	if len(res.Description) == 0 {
		return errors.New("Description can't be blank")
	}
	return nil
}

type CommentUpdate struct {
	CommentId   string `json:"comment_id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	UsereId     string `json:"user_id" gorm:"column:user_id;"`
	Category    string `json:"category" gorm:"column:category"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   string `json:"deleted_at" gorm:"column:deleted_at"`
}

func (CommentUpdate) TableName() string {
	return Comment{}.TableName()
}

func (res *CommentUpdate) Validate() error {
	res.Title = strings.TrimSpace(res.Title)

	if len(res.Title) == 0 {
		return errors.New("Title can't be blank")
	}
	if len(res.Description) == 0 {
		return errors.New("Description can't be blank")
	}
	return nil
}
