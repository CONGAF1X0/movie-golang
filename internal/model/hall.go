package model

import (
	"github.com/jinzhu/gorm"
)

type Hall struct {
	HallID   int    `json:"hall_id" gorm:"primaryKey"`
	HallName string `json:"hall_name"`
	CinemaID uint64 `json:"cinema_id"`
	Capacity string `json:"capacity"`
}

func (h Hall) TableName() string {
	return "hall"
}

func (h Hall) List(db *gorm.DB, offset, size int) ([]Hall, error) {
	var hall []Hall
	var err error
	if offset >= 0 && size > 0 {
		db = db.Offset(offset).Limit(size)
	}
	if err = db.Where(&h).Find(&hall).Error; err != nil {
		return nil, err
	}
	return hall, nil
}

func (h Hall) Create(db *gorm.DB) error {
	return db.Create(&h).Error
}

func (h Hall) Update(db *gorm.DB) error {
	return db.Model(&Hall{}).Where("hall_id = ?", h.HallID).Update(h).Error
}

func (h Hall) Delete(db *gorm.DB) error {
	return db.Delete(&h).Error
}
