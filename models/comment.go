package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

// Comment ..
type Comment struct {
	gorm.Model
	PostID uint `gorm:"not null" json:"post_id" validate:"required"`
	// Name      string `gorm:"not null" json:"name" validate:"required"`
	Content   string `gorm:"type:text not null" json:"content" validate:"required"`
	LikeCount int    `gorm:"default:0" json:"like_count"`
	ParentID  uint   `json:"Parent_id"`
	IP        string `json:"ip"`
	UserID    uint   `json:"user_id" validate:"required"`
	User      User   `json:"user" validate:"-"`
}

var validate *validator.Validate

// NewComment ..
func NewComment(c Comment) (interface{}, error) {
	validate = validator.New()
	err := validate.Struct(&c)
	if err != nil {
		return c, err
	}
	ret := orm.Create(&c)
	if ret.Error != nil {
		return c, ret.Error
	}
	return c, nil
}

// DeleteComment ..
func DeleteComment(id uint) (Comment, error) {
	var c Comment
	c.ID = id
	ret := orm.First(&c).Delete(&c)
	if ret.Error != nil {
		return c, ret.Error
	}
	return c, nil
}
