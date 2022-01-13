package dao

import "TicketSales/internal/model"

func (d Dao) SessionListGroupByMovie(id uint64) (model.SessApi, error) {
	return model.Session{CinemaID: id}.ListGroupByMovie(d.engine)
}