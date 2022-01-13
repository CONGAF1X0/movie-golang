package service

import (
	"TicketSales/global"
	"TicketSales/internal/dao"
	"context"

)

type Service struct {
	ctx context.Context
	Dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.Dao = dao.New(global.DBEngine)
	return svc
}
