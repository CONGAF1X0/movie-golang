package dao

import "TicketSales/internal/model"

func (d Dao) Movie(id int) (model.MovieApi,error) {
	movie := model.Movie{MovieID: id}
	return movie.Get(d.engine)
}

func (d Dao) MovieList(name string,offset,size int) ([]*model.Movie,error) {
	movie :=model.Movie{MovieName: name}
	return movie.List(d.engine,offset,size)
}

func (d Dao) HotMovies(city string) ([]model.MovieSimple, error) {
	return model.Cinema{City: city}.HotMovies(d.engine)
}