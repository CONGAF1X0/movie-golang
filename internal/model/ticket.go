package model

import (
	"github.com/jinzhu/gorm"
)

type Ticket struct {
	TicketID    uint64 `json:"ticket_id"`
	UID         uint64 `json:"uid"`
	SessionID   int    `json:"session_id"`
	Seat        string `json:"seat"`
	CreatedTime uint32 `json:"created_time"`
}
type TicketDetail struct {
	Ticket
	Session
	Hall
	Cinema
	Movie
}

func (t Ticket) TableName() string {
	return "ticket"
}
func (t Ticket) List(db *gorm.DB, pageOffset, pageSize int) ([]*Ticket, error) {
	var ticket []*Ticket
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Where(&t).Find(&ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}
func (t Ticket) DetailList(db *gorm.DB) (res []*TicketDetail, err error) {
	db = db.Raw("SELECT * FROM ticket as t,`session` as s ,movie as m,cinema as c,hall as h WHERE t.session_id=s.session_id AND s.movie_id=m.movie_id AND s.cinema_id=c.cinema_id AND h.hall_id=s.hall_id "+
		"AND uid=?", t.UID).Order("created_time desc")
	if err = db.Find(&res).Error; err != nil {
		return
	}
	return
}

func (t Ticket) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Ticket) Update(db *gorm.DB) error {
	return db.Model(&Ticket{}).Where("ticket_id = ?", t.TicketID).Update(t).Error
}

func (t Ticket) Delete(db *gorm.DB) error {
	return db.Where("ticket_id = ?", t.TicketID).Delete(&t).Error
}
