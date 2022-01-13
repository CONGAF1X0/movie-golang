package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Ticket struct {
	TicketID int `json:"ticket_id"`
	//Base UserBase `gorm:"joinForeignKey:UID"'`
	//Session  Session  `gorm:"joinForeignKey:SessionID"`
	UID       uint64 `json:"uid"`
	SessionID uint64 `json:"session_id"`
	Seat      string `json:"seat"`
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
	db=db.Where("uid=?", t.UID)
	if err = db.Model(Ticket{}).Find(&ticket).Error; err != nil {
		return nil, err
	}
	fmt.Println(ticket[1].SessionID)
	return ticket, nil
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
