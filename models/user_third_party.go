package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

// User UserThirdParty 第三方用户
type User struct {
	Login    string `gorm:"size:30" json:"login" validate:"required"`
	Name     string `gorm:"size:30" json:"name" validate:"required"`
	Email    string `gorm:"size:30;unique;index;not null" json:"email"`
	Avatar   string `json:"avatar"`
	PID      uint   `gorm:"not null" json:"p_id" validate:"required"`
	Platform string `gorm:"size:20;not null" json:"platform" validate:"required"`
	gorm.Model
}

// RegisterUser ..
func RegisterUser(u User) (User, error) {
	fmt.Println(u)
	validate = validator.New()
	err := validate.Struct(&u)
	if err != nil {
		return u, err
	}
	if u.Name == "" {
		u.Name = u.Login
	}
	if orm.Where("p_id = ? AND platform = ?", u.PID, u.Platform).First(&u).RecordNotFound() {
		ret := orm.Create(&u)
		if ret.Error != nil {
			return u, ret.Error
		}
	}
	return u, nil
}
