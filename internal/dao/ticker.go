package dao

import "TicketSales/internal/model"

func (d *Dao) CreateTicket(tid uint64, sid int, uid uint64, seat string) error {
	return model.Ticket{TicketID: tid, SessionID: sid, UID: uid, Seat: seat}.Create(d.engine)
}
func (d *Dao) TicketList(uid uint64) ([]*model.TicketDetail, error) {
	return model.Ticket{UID: uid}.DetailList(d.engine)
}
