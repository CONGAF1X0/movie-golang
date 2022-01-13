package service

import (
	"TicketSales/internal/model"
)

type GetMovieReq struct {
	Name string `form:"name"`
}

func (svc Service) GetMovie(id int) (model.MovieApi, error) {
	return svc.Dao.Movie(id)
}

func (svc *Service) GetMovieList(param *GetMovieReq, offset, size int) ([]*model.Movie, error) {
	return svc.Dao.MovieList(param.Name, offset, size)
}

type HotMoviesReq struct {
	City string `form:"city" binding:"required"`
}
func (svc Service) GetHotMovies(param HotMoviesReq) ([]model.MovieSimple, error) {
	return svc.Dao.HotMovies(param.City)
}