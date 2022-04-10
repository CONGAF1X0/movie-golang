package model

import (
	"github.com/jinzhu/gorm"
)

type UserBase struct {
	Model
	UID      uint64 `json:"uid"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Gender   uint16 `json:"gender"`
	Birthday uint64 `json:"birthday"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Face     string `json:"face"`
	CinemaID uint64 `json:"cinema_id"`
}

func (u UserBase) IsExist(db *gorm.DB) (count int, err error) {
	err = db.Model(&UserBase{}).Where("user_name = ? OR mobile = ? OR email = ?", u.UserName, u.Mobile, u.Email).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u UserBase) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u UserBase) GetByMobile(db *gorm.DB) (UserBase, error) {
	var base UserBase
	err := db.Where("mobile=?", u.Mobile).Take(&base).Error
	if err != nil {
		return base, err
	}
	return base, nil
}

func (u UserBase) Get(db *gorm.DB) (UserBase, error) {
	var base UserBase
	err := db.Where(&u).Take(&base).Error
	if err != nil {
		return base, err
	}
	return base, nil
}

func (u UserBase) Update(db *gorm.DB) error {
	return db.Model(&UserBase{}).Where("uid = ?", u.UID).Update(u).Error
}

func (u UserBase) Delete(db *gorm.DB) error {
	return db.Where("uid = ?", u.UID).Delete(&u).Error
}
