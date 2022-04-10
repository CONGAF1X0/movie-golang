package dao

import "TicketSales/internal/model"

func (d *Dao) HallList(cid uint64, offset, size int) ([]model.Hall, error) {
	return model.Hall{CinemaID: cid}.List(d.engine, offset, size)
}
func (d *Dao) CreateHall(cid uint64, name, cap string) error {
	return model.Hall{CinemaID: cid, HallName: name, Capacity: cap}.Create(d.engine)
}
func (d *Dao) DelHall(hid int) error {
	return model.Hall{HallID: hid}.Delete(d.engine)
}
func (d *Dao) UpdateHall(hid int, name, cap string) error {
	return model.Hall{HallName: name, HallID: hid, Capacity: cap}.Update(d.engine)
}
