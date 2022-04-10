package service

import (
	"TicketSales/internal/model"
	"strconv"
	"strings"
)

type GetMovieReq struct {
	Query string `form:"query" json:"query"`
}

func (svc *Service) GetMovie(id int) (model.MovieApi, error) {
	return svc.Dao.Movie(id)
}

func (svc *Service) GetRuntime(id int) (int, error) {
	str, err := svc.Dao.GetRuntime(id)
	if err != nil {
		return 0, err
	}
	id, err = strconv.Atoi(strings.Trim(str, "分钟"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *Service) GetMovieList(param *GetMovieReq, offset, size int) ([]*model.Movie, error) {
	return svc.Dao.MovieList(param.Query, offset, size)
}

type HotMoviesReq struct {
	City string `form:"city" binding:"required" json:"city"`
}

func (svc *Service) GetHotMovies(param *HotMoviesReq, offset, size int) ([]model.MovieSimple, int, error) {
	return svc.Dao.HotMovies(param.City, offset, size)
}
