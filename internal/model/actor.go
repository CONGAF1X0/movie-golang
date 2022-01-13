package model

import (
	"github.com/jinzhu/gorm"
	"strings"
)

type Actor struct {
	ActorID int `json:"actor_id" gorm:"primaryKey;column:id"`
	Name1 string `json:"name1"`
	Name2 string `json:"name2"`
	Birthday string `json:"birthday"`
	Introduction string `json:"introduction"`
	Image string `json:"image"`
}

func (a Actor) TableName() string {
	return "actor"
}

func (a Actor) List(db *gorm.DB,ids []int) ([]*Actor,error) {
	var arr []*Actor
	var err error
	if err = db.Model(Actor{}).Find(&arr,ids).Error;err !=nil {
		return nil,err
	}
	return arr,nil
}

func (a Actor) StrList(db *gorm.DB,ids []int) (string,error) {
	arr, err := a.List(db, ids)
	if err != nil {
		return "", err
	}
	str := ""
	for i := 0; i < len(arr); i++ {
		str += arr[i].Name1 + "、"
		if i == 2{
			break
		}
	}
	return strings.TrimRight(str,"、"),nil
}

func (a Actor) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Actor) Update(db *gorm.DB) error {
	return db.Model(&Actor{}).Where("id=?",a.ActorID).Update(a).Error
}

func (a Actor) Delete(db *gorm.DB) error {
	return db.Where("id=?",a.ActorID).Delete(&a).Error
}