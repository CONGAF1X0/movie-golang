package v1

import (
	"TicketSales/global"
	"TicketSales/internal/service"
	"TicketSales/pkg/app"
	"TicketSales/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Actor struct {
}

func NewActor() Actor {
	return Actor{}
}

func (a Actor) Get(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	resp := app.NewResponse(c)
	if err != nil {
		resp.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	data, err := svc.GetActor(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ToErrorResponse(errcode.NotFound)
			return
		}
		global.Logger.Errorf("svc.GetCinema err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{
		"code": errcode.Success.Code(),
		"data": data,
	})
}

func (a Actor) List(c *gin.Context) {
	param := service.GetActorReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	size := app.GetPageSize(c)
	offset := app.GetPageOffset(app.GetPage(c), size)

	svc := service.New(c.Request.Context())
	data, err := svc.GetActorList(&param, offset, size)
	if err != nil {
		global.Logger.Errorf("svc.GetActorList err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponseList(data, 0)
}
func (a Actor) Create(c *gin.Context) {
	param := service.CreateActorReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	err := svc.CreateActor(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateActor err: %v", err)
		resp.ToErrorResp(err)
		return
	}

	resp.ToResponse(gin.H{
		"code": 201,
	})
}

func (a Actor) Update(c *gin.Context) {
	param := service.UpdateActorReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	err := svc.UpdateActor(&param)
	if err != nil {
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{"code": 200})
}

func (a Actor) Del(c *gin.Context) {
	param := service.UpdateActorReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	err := svc.DelActor(param.ActorID)
	if err != nil {
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{"code": 200})
}
