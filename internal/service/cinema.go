package service

import (
	"TicketSales/internal/model"
)

type GetCinemaReq struct {
	City     string `form:"city" binding:"required"`
	District string `form:"district"`
	Query    string `form:"query"`
	MID      int    `form:"mid"`
	t        int64  `form:"t"`
}

type CinemaSessResp struct {
	Cinema  model.Cinema
	Session model.SessApi
}

func (svc *Service) GetCinemaList(param *GetCinemaReq, offset, size int) (interface{}, error) {
	if param.MID != 0 {
		return svc.Dao.CinemaSessList(param.City,param.District,param.Query,param.MID,param.t,offset,size)
	}
	return svc.Dao.CinemaList(param.City, param.District, param.Query, offset, size)
}

func (svc Service) GetCinema(id uint64) (CinemaSessResp, error) {
	var resp CinemaSessResp
	var err error
	resp.Cinema, err = svc.Dao.Cinema(id)
	if err != nil {
		return resp, err
	}
	resp.Session, err = svc.Dao.SessionListGroupByMovie(id)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
