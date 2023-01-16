package usermodel

import (
	"errors"
	"gin_form/common"
	"strings"
)

type User struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
	Phone   string `json:"phone" gorm:"column:phone;"`
	About   string `json:"about" gorm:"column:about;"`
	Email   string `json:"email" gorm:"column:email;"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
	Phone   string `json:"phone" gorm:"column:phone;"`
	About   string `json:"about" gorm:"column:about;"`
	Email   string `json:"email" gorm:"column:email;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (res *UserCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("user name can't be blank")
	}
	if len(res.Address) == 0 {
		return errors.New("user addr can't be blank")
	}
	return nil
}

type UserUpdate struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
	Phone   string `json:"phone" gorm:"column:phone;"`
	About   string `json:"about" gorm:"column:about;"`
	Email   string `json:"email" gorm:"column:email;"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (res *UserUpdate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("user name can't be blank")
	}
	if len(res.Address) == 0 {
		return errors.New("user addr can't be blank")
	}
	return nil
}
