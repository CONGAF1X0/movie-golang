package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB{
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
