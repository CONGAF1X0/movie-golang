package v1

import (
	"TicketSales/global"
	"TicketSales/internal/service"
	"TicketSales/pkg/app"
	"TicketSales/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Ticket struct {
}

func NewTicket() Ticket {
	return Ticket{}
}

func (t Ticket) List(c *gin.Context) {
	resp := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	data, err := svc.TicketList(c.GetUint64("uid"))
	if err != nil {
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponseList(data, len(data))
}

func (t Ticket) Create(c *gin.Context) {
	param := service.CreateTicketReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if param.UID != c.GetUint64("uid") {
		resp.ToErrorResponse(errcode.Forbidden)
		return
	}

	err := svc.CreateTicket(&param)
	if err != nil {
		if err.Error() == "sold" {
			resp.ToErrorResponse(errcode.Sold)
			return
		}
		global.Logger.Errorf("svc.CreateTicket err: %v", err)
		resp.ToErrorResp(err)
		return
	}

	resp.ToResponse(gin.H{
		"code": 201,
	})
}
