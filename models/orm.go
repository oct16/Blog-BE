package models

import (
	"echo-blog/conf"
	"fmt"
	"log"
	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var orm *gorm.DB

// NewOrm ..
func NewOrm() {
	var err error
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "blog_" + defaultTableName
	}
	var url string
	url = "localhost:3306"
	if conf.Env == "production" {
		url = "oct16.cn:3307"
	}
	o, err := gorm.Open("mysql", "root:123456@tcp("+url+")/blog?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Fail to create engine")
	}
	o.LogMode(true)

	// defer o.Close()
	o.AutoMigrate(&Post{}, &SuperUser{}, &Tag{}, &Comment{}, &Comment{}, &User{})

	orm = o
}
