package dao

import (
	"TicketSales/internal/model"
)

func (d *Dao) GetBase(p interface{}) (model.UserBase, error) {
	var base = model.UserBase{}
	switch p.(type) {
	case uint64:
		base.UID = p.(uint64)
	case string:
		base.Mobile = p.(string)
	}

	return base.Get(d.engine)
}

func (d *Dao) UserBindCinema(uid, cid uint64) error {
	return model.UserBase{UID: uid, CinemaID: cid}.Update(d.engine)
}
