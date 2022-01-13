package v1

import (
	"TicketSales/global"
	"TicketSales/internal/service"
	"TicketSales/pkg/app"
	"TicketSales/pkg/errcode"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Cinema struct {
}

func NewCinema() Cinema {
	return Cinema{}
}

func (ci Cinema) List(c *gin.Context) {
	param := service.GetCinemaReq{}
	resp := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	size :=app.GetPageSize(c)
	offset:=app.GetPageOffset(app.GetPage(c),size)

	svc := service.New(c.Request.Context())
	data,err := svc.GetCinemaList(&param,offset,size)
	if err != nil {
		global.Logger.Errorf("svc.GetCinemaList err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponseList(data,0)
}

func (ci Cinema) Get(c *gin.Context) {
	strID := c.Param("id")
	id , _:= strconv.ParseUint(strID,10,64)
	resp := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	data,err:=svc.GetCinema(id)
	if err!=nil {
		global.Logger.Errorf("svc.GetCinema err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{
		"code":errcode.Success.Code(),
		"data":data,
	})
}
