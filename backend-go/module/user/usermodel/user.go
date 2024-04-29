package usermodel

import (
	"cinema/common"
	"errors"
	"time"
)

const (
	TableName = "users"
)

type User struct {
	common.SQLModel `json:",inline"`
	DateOfBirth     *time.Time `json:"date_of_birth" gorm:"column:date_of_birth;"`
	Email           string     `json:"email" gorm:"column:email;"`
	Password        string     `json:"-" gorm:"column:password;"`
	Gender          string     `json:"gender" gorm:"gender"`
	LastName        string     `json:"last_name" gorm:"column:last_name;"`
	FirstName       string     `json:"first_name" gorm:"column:first_name;"`
	Tier            string     `json:"tier" gorm:"column:tier;"`
	Salt            string     `json:"-" gorm:"column:salt;"`
	Phone           string     `json:"phone" gorm:"column:phone_number;"`
}

func (u *User) GetUserId() int {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Tier
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

func (User) TableName() string {
	return TableName
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)