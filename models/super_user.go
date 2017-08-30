package models

import (
	"github.com/jinzhu/gorm"
)

// SuperUser ..
type SuperUser struct {
	Name     string `gorm:"size:30;unique" json:"name"`
	Avatar   string `json:"avatar"`
	Email    string `gorm:"size:30;unique;index;not null" json:"email"`
	Password string `gorm:"size:20;not null" json:"password"`
	gorm.Model
}

// GetSuperUser ..
func GetSuperUser(u SuperUser) SuperUser {
	if len(u.Email) > 0 && len(u.Password) > 0 {
		orm.Where(&u).First(&u)
	}
	return u
}
