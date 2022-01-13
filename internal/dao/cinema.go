package dao

import "TicketSales/internal/model"

func (d Dao) CinemaList(city,dis,query string, offset, size int) ([]*model.Cinema,error) {
	c := model.Cinema{CinemaName: query,City: city,District: dis}
	return c.List(d.engine,offset,size)
}

func (d Dao) Cinema(id uint64) (model.Cinema,error) {
	return model.Cinema{CinemaID: id}.Get(d.engine)
}

func (d Dao) CinemaSessList(city,dis,query string,mid int,t int64, offset, size int) ([]model.CinemaSess, error) {
	return model.Cinema{City: city,District: dis,CinemaName: query}.CinemaSessList(d.engine,model.Session{MovieID: mid,StartTime: t},offset,size)
}