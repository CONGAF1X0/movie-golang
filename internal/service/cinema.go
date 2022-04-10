package service

import (
	"TicketSales/internal/model"
	"strings"
)

type GetCinemaReq struct {
	City     string  `form:"city" binding:"required" json:"city"`
	Lat      float64 `json:"lat" form:"lat"`
	Lng      float64 `json:"lng" form:"lng"'`
	District string  `form:"district" json:"district"`
	Query    string  `form:"query" json:"query"`
	MID      int     `form:"mid" json:"mid"`
	T        int64   `form:"t" json:"t"`
}

type UpdateCinemaReq struct {
	CinemaID   uint64  `json:"cinema_id" binding:"required"`
	CinemaName string  `json:"cinema_name" `
	Mobile     string  `json:"mobile" `
	Province   string  `json:"province"`
	City       string  `json:"city" `
	District   string  `json:"district" `
	Location   string  `json:"location" `
	Longitude  float64 `json:"longitude" `
	Latitude   float64 `json:"latitude"`
}
type CreateCinemaReq struct {
	CinemaName string  `json:"cinema_name" binding:"required"`
	Mobile     string  `json:"mobile" binding:"required"`
	Province   string  `json:"province" binding:"required"`
	City       string  `json:"city" binding:"required"`
	District   string  `json:"district" binding:"required"`
	Location   string  `json:"location" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
}

type CreateHallReq struct {
	CinemaID uint64 `json:"cinema_id" binding:"required"`
	HallName string `json:"hall_name" binding:"required"`
	Capacity string `json:"capacity" binding:"required"`
}
type DelHallReq struct {
	HallID   int    `json:"hall_id" binding:"required"`
	CinemaID uint64 `json:"cinema_id" binding:"required"`
}
type UpdateHallReq struct {
	HallID   int    `json:"hall_id" binding:"required"`
	CinemaID uint64 `json:"cinema_id" binding:"required"`
	HallName string `json:"hall_name"`
	Capacity string `json:"capacity"`
}

type CinemaSessResp struct {
	Cinema  model.Cinema
	Session []model.MovieSess
}
type UpdateSessionReq struct {
	SessionID int     `json:"session_id"binding:"required"`
	CinemaID  uint64  `json:"cinema_id" binding:"required"`
	MovieID   int     `json:"movie_id"`
	HallID    int     `json:"hall_id"`
	StartTime int64   `json:"start_time" `
	EndTime   int64   `json:"end_time"`
	Price     float32 `json:"price" `
}
type CreateSessionReq struct {
	CinemaID  uint64  `json:"cinema_id" binding:"required"`
	MovieID   int     `json:"movie_id" binding:"required"`
	HallID    int     `json:"hall_id" binding:"required"`
	StartTime int64   `json:"start_time" binding:"required"`
	EndTime   int64   `json:"end_time" binding:"required"`
	Price     float32 `json:"price" binding:"required"`
}
type DelSessionReq struct {
	SessionID int    `json:"session_id"binding:"required"`
	CinemaID  uint64 `json:"cinema_id" binding:"required"`
}

func (svc *Service) GetCinemaList(param *GetCinemaReq, offset, size int) (interface{}, int, error) {
	if param.MID != 0 {
		return svc.Dao.CinemaSessList(param.City, param.District, param.Query, param.Lat, param.Lng, param.MID, param.T, offset, size)
	}
	return svc.Dao.CinemaList(param.City, param.District, param.Query, param.Lat, param.Lng, offset, size)
}
func (svc *Service) UpdateSession(param *UpdateSessionReq) error {
	return svc.Dao.UpdateSession(param.SessionID, param.MovieID, param.HallID, param.CinemaID, param.StartTime, param.EndTime, param.Price)
}
func (svc *Service) CreateSession(param *CreateSessionReq) error {
	return svc.Dao.CreateSession(param.CinemaID, param.MovieID, param.HallID, param.StartTime, param.EndTime, param.Price)
}
func (svc *Service) DelSession(param *DelSessionReq) error {
	return svc.Dao.DelSession(param.SessionID)
}
func (svc *Service) UpdateCinema(param *UpdateCinemaReq) error {
	return svc.Dao.UpdateCinema(param.CinemaID, param.CinemaName, param.Mobile, param.Province, param.City, param.District, param.Location, param.Longitude, param.Latitude)
}
func (svc *Service) CreateCinema(id uint64, param *CreateCinemaReq) error {
	return svc.Dao.CreateCinema(id, param.CinemaName, param.Mobile, param.Province, param.City, param.District, param.Location, param.Longitude, param.Latitude)
}

func (svc *Service) GetCinema(id uint64) (model.Cinema, error) {
	return svc.Dao.Cinema(id)
}

func (svc *Service) GetHallList(cid uint64, offset, size int) ([]model.Hall, error) {
	return svc.Dao.HallList(cid, offset, size)
}

func (svc *Service) CreateHall(param *CreateHallReq) error {
	return svc.Dao.CreateHall(param.CinemaID, param.HallName, param.Capacity)
}
func (svc *Service) DelHall(hid int) error {
	return svc.Dao.DelHall(hid)
}
func (svc *Service) UpdateHall(param *UpdateHallReq) error {
	return svc.Dao.UpdateHall(param.HallID, param.HallName, param.Capacity)
}
func (svc *Service) GetSessionList(cid uint64, t int64, offset, size int) ([]model.MovieSess, error) {
	return svc.Dao.SessionListGroupByMovie(cid, t)
}
func (svc *Service) GetSession(id int) (model.SessInfo, error) {
	return svc.Dao.GetSession(id)
}

func (svc Service) SoldSeat(id int) ([]string, error) {
	data, err := svc.Dao.SoldSeat(id)
	var res []string
	for _, ticket := range data {
		res = append(res, strings.Split(ticket.Seat, ",")...)
	}
	if err != nil {
		return []string{}, err
	}
	return res, nil
}
