package dao

import "TicketSales/internal/model"

func (d *Dao) SessionListGroupByMovie(id uint64, t int64) ([]model.MovieSess, error) {
	return model.Session{CinemaID: id, StartTime: t}.ListGroupByMovie(d.engine)
}
func (d *Dao) GetSession(id int) (model.SessInfo, error) {
	return model.Session{SessionID: id}.Get(d.engine)
}

func (d *Dao) UpdateSession(sid, mid, hid int, cid uint64, st, et int64, price float32) error {
	return model.Session{SessionID: sid, MovieID: mid, CinemaID: cid, HallID: hid, StartTime: st, EndTime: et, Price: price}.Update(d.engine)
}
func (d *Dao) CreateSession(cid uint64, mid, hid int, st, et int64, price float32) error {
	return model.Session{CinemaID: cid, MovieID: mid, HallID: hid, StartTime: st, EndTime: et, Price: price}.Create(d.engine)
}
func (d *Dao) DelSession(sid int) error {
	return model.Session{SessionID: sid}.Delete(d.engine)
}

func (d *Dao) SoldSeat(sid int) ([]*model.Ticket, error) {
	return model.Ticket{SessionID: sid}.List(d.engine, 0, 0)
}
