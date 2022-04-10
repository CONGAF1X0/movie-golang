package model

import (
	"github.com/jinzhu/gorm"
)

type UserAuth struct {
	Model
	ID           uint64 `json:"id"`
	UID          uint64 `json:"uid"`
	IdentityType uint8  `json:"identity_type"`
	Identity     string `json:"identity"`
	Certificate  string `json:"certificate"`
}

func (u UserAuth) Exist(db *gorm.DB) (count int, err error) {
	err = db.Model(&UserAuth{}).Where("identity =?", u.Identity).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u UserAuth) Get(db *gorm.DB) (UserAuth, error) {
	var auth UserAuth
	db = db.Where("identity = ? AND certificate = ? AND identity_type = ?", u.Identity, u.Certificate, u.IdentityType)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}

func (u UserAuth) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u UserAuth) Update(db *gorm.DB) error {
	return db.Model(UserAuth{}).Where("uid=?", u.UID).Update(u).Error
}

func (u UserAuth) Delete(db *gorm.DB) error {
	return db.Where("uid=?", u.UID).Delete(&u).Error
}
