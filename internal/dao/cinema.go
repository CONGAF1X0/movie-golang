package dao

import (
	"TicketSales/internal/model"
)

func (d *Dao) CinemaList(city, dis, query string, lat, lng float64, offset, size int) ([]model.CinemaWithDistance, int, error) {
	c := model.Cinema{CinemaName: query, City: city, District: dis, Latitude: lat, Longitude: lng}
	return c.List(d.engine, offset, size)
}

func (d *Dao) Cinema(id uint64) (model.Cinema, error) {
	return model.Cinema{CinemaID: id}.Get(d.engine)
}

func (d *Dao) CinemaSessList(city, dis, query string, lat, lng float64, mid int, t int64, offset, size int) ([]model.CinemaSess, int, error) {

	return model.Cinema{City: city, District: dis, CinemaName: query, Latitude: lat, Longitude: lng}.CinemaSessList(d.engine, model.Session{MovieID: mid, StartTime: t}, offset, size)
}

func (d *Dao) UpdateCinema(cid uint64, cname, mob, prov, city, dist, loc string, lng, lat float64) error {
	return model.Cinema{CinemaID: cid, CinemaName: cname, Mobile: mob, Province: prov, City: city, District: dist, Location: loc, Longitude: lng, Latitude: lat}.Update(d.engine)
}

func (d *Dao) CreateCinema(cid uint64, cname, mob, prov, city, dist, loc string, lng, lat float64) error {
	return model.Cinema{CinemaID: cid, CinemaName: cname, Mobile: mob, Province: prov, City: city, District: dist, Location: loc, Longitude: lng, Latitude: lat}.Create(d.engine)
}
