package service

import (
	"TicketSales/internal/model"
	"TicketSales/pkg/uid"
	"errors"
	"strings"
)

type CreateTicketReq struct {
	SessionID int    `json:"session_id" binding:"required"`
	UID       uint64 `json:"uid" binding:"required"`
	Seat      string `json:"seat" binding:"required"`
}

func (svc *Service) CreateTicket(param *CreateTicketReq) error {
	soldArr, _ := svc.SoldSeat(param.SessionID)
	seatArr := strings.Split(param.Seat, ",")
	for _, seat := range seatArr {
		for _, sold := range soldArr {
			if sold == seat {
				return errors.New("sold")
			}
		}
	}
	tid, err := uid.Sf.NextID(1)
	if err != nil {
		return err
	}
	return svc.Dao.CreateTicket(tid, param.SessionID, param.UID, param.Seat)
}

func (svc *Service) TicketList(uid uint64) ([]*model.TicketDetail, error) {
	return svc.Dao.TicketList(uid)
}
