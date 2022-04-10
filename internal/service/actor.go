package service

import "TicketSales/internal/model"

type GetActorReq struct {
	Query string `form:"query" json:"query"`
}
type UpdateActorReq struct {
	ActorID      int    `json:"actor_id" binding:"required"`
	Name1        string `json:"name1" `
	Name2        string `json:"name2"`
	Birthday     int64  `json:"birthday"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
}
type CreateActorReq struct {
	Name1        string `json:"name1" binding:"required"`
	Name2        string `json:"name2" binding:"required"`
	Birthday     int64  `json:"birthday" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Avatar       string `json:"avatar" binding:"required"`
}

func (svc *Service) GetActorList(param *GetActorReq, offset, size int) ([]*model.Actor, error) {
	return svc.Dao.ActorList(param.Query, offset, size)
}

func (svc *Service) UpdateActor(param *UpdateActorReq) error {
	return svc.Dao.UpdateActor(param.ActorID, param.Name1, param.Name2, param.Introduction, param.Avatar, param.Birthday)
}

func (svc *Service) GetActor(id int) (model.Actor, error) {
	return svc.Dao.GetActor(id)
}

func (svc *Service) DelActor(id int) error {
	return svc.Dao.DelActor(id)
}

func (svc *Service) CreateActor(param *CreateActorReq) error {
	return svc.Dao.CreateActor(param.Name1, param.Name2, param.Introduction, param.Avatar, param.Birthday)
}
