package articlemodel

import (
	"errors"
	"strings"
)

type Article struct {
	ArticleId   string `json:"article_id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	UsereId     string `json:"user_id" gorm:"column:user_id"`
	Category    string `json:"category" gorm:"column:category"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   string `json:"deleted_at" gorm:"column:deleted_at"`
}

func (Article) TableName() string {
	return "articles"
}

type ArticleCreate struct {
	ArticleId   string `json:"article_id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	UsereId     string `json:"user_id" gorm:"column:user_id"`
	Category    string `json:"category" gorm:"column:category"`
}

func (ArticleCreate) TableName() string {
	return Article{}.TableName()
}

func (res *ArticleCreate) Validate() error {
	res.Title = strings.TrimSpace(res.Title)

	if len(res.Title) == 0 {
		return errors.New("Title can't be blank")
	}
	if len(res.Description) == 0 {
		return errors.New("Description can't be blank")
	}
	return nil
}

type ArticleUpdate struct {
	ArticleId   string `json:"article_id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	UsereId     string `json:"user_id" gorm:"column:user_id"`
	Category    string `json:"category" gorm:"column:category"`
}

func (ArticleUpdate) TableName() string {
	return Article{}.TableName()
}

func (res *ArticleUpdate) Validate() error {
	res.Title = strings.TrimSpace(res.Title)

	if len(res.Title) == 0 {
		return errors.New("Title can't be blank")
	}
	return nil
}
