package v1

import (
	"TicketSales/global"
	"TicketSales/internal/service"
	"TicketSales/pkg/app"
	"TicketSales/pkg/errcode"
	"TicketSales/pkg/uid"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	size := app.GetPageSize(c)
	offset := app.GetPageOffset(app.GetPage(c), size)

	svc := service.New(c.Request.Context())
	data, total, err := svc.GetCinemaList(&param, offset, size)
	if err != nil {
		global.Logger.Errorf("svc.GetCinemaList err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponseList(data, total)
}

func (ci Cinema) Get(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	resp := app.NewResponse(c)
	if err != nil {
		resp.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	data, err := svc.GetCinema(id)
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

func (ci Cinema) HallList(c *gin.Context) {
	strID := c.Query("cinema_id")
	id, err := strconv.ParseUint(strID, 10, 64)
	resp := app.NewResponse(c)
	if err != nil {
		resp.ToErrorResponse(errcode.InvalidParams)
		return
	}
	size := app.GetPageSize(c)
	offset := app.GetPageOffset(app.GetPage(c), size)

	svc := service.New(c.Request.Context())
	data, err := svc.GetHallList(id, offset, size)
	if err != nil {
		global.Logger.Errorf("svc.GetCinemaList err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponseList(data, len(data))
}

func (ci *Cinema) CreateHall(c *gin.Context) {
	param := service.CreateHallReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := app.CheckPermission(c, &svc, param)
	if err != nil {
		resp.ToErrorResponse(errcode.Forbidden)
		return
	}

	err = svc.CreateHall(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateHall err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponse(gin.H{"code": 201})
}
func (ci *Cinema) DelHall(c *gin.Context) {
	param := service.DelHallReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := app.CheckPermission(c, &svc, param)
	if err != nil {
		resp.ToErrorResponse(errcode.Forbidden)
		return
	}

	err = svc.DelHall(param.HallID)
	if err != nil {
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{"code": 200})
}
func (ci *Cinema) UpdateHall(c *gin.Context) {
	param := service.UpdateHallReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	err := app.CheckPermission(c, &svc, param)
	if err != nil {
		resp.ToErrorResponse(errcode.Forbidden)
		return
	}

	err = svc.UpdateHall(&param)
	if err != nil {
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{"code": 200})
}

func (ci Cinema) SessionList(c *gin.Context) {
	strID := c.Query("cinema_id")
	id, err := strconv.ParseUint(strID, 10, 64)
	resp := app.NewResponse(c)
	if id <= 0 || err != nil {
		resp.ToErrorResponse(errcode.InvalidParams)
		return
	}
	strT := c.Query("t")
	var t int64 = 0
	if strT != "" {
		t, err = strconv.ParseInt(strT, 10, 64)
		if err != nil {
			resp.ToErrorResponse(errcode.InvalidParams)
			return
		}
	}
	size := app.GetPageSize(c)
	offset := app.GetPageOffset(app.GetPage(c), size)

	svc := service.New(c.Request.Context())
	data, err := svc.GetSessionList(id, t, offset, size)
	if err != nil {
		global.Logger.Errorf("svc.GetCinemaList err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponseList(data, len(data))
}
func (ci Cinema) GetSession(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	resp := app.NewResponse(c)
	if id < 0 || err != nil {
		resp.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	data, err := svc.GetSession(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ToErrorResponse(errcode.NotFound)
			return
		}
		global.Logger.Errorf("svc.GetSession err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{
		"code": errcode.Success.Code(),
		"data": data,
	})
}
func (ci Cinema) SoldSeat(c *gin.Context) {
	strID := c.Query("sid")
	id, err := strconv.Atoi(strID)
	resp := app.NewResponse(c)
	if id <= 0 || err != nil {
		resp.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	data, err := svc.SoldSeat(id)
	if err != nil {
		global.Logger.Errorf("svc.SoldSeat err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{
		"code": errcode.Success.Code(),
		"data": data,
	})
}
func (ci Cinema) Update(c *gin.Context) {
	param := service.UpdateCinemaReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := app.CheckPermission(c, &svc, param)
	if err != nil {
		resp.ToErrorResponse(errcode.Forbidden)
		return
	}

	err = svc.UpdateCinema(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateCinema err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{
		"code": 200,
	})
}

func (ci Cinema) Create(c *gin.Context) {
	param := service.CreateCinemaReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	cid, err := uid.Sf.NextID(2)
	if err != nil {
		global.Logger.Errorf("app.GenSnowflake err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	err = svc.CreateCinema(cid, &param)
	if err != nil {
		global.Logger.Errorf("svc.CreateCinema err: %v", err)
		resp.ToErrorResp(err)
		return
	}
	err = svc.UserBindCinema(c.GetUint64("uid"), cid)
	if err != nil {
		global.Logger.Errorf("svc.CreateCinema err: %v", err)
		resp.ToErrorResp(err)
		return
	}
	resp.ToResponse(gin.H{
		"code": 201,
	})
}
func (ci Cinema) UpdateSession(c *gin.Context) {
	param := service.UpdateSessionReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := app.CheckPermission(c, &svc, param)
	if err != nil {
		resp.ToErrorResponse(errcode.Forbidden)
		return
	}
	err = svc.UpdateSession(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateSession err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponse(gin.H{
		"code": 200,
	})
}
func (ci Cinema) CreateSession(c *gin.Context) {
	param := service.CreateSessionReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	err := app.CheckPermission(c, &svc, param)
	if err != nil {
		resp.ToErrorResponse(errcode.Forbidden)
		return
	}
	err = svc.CreateSession(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateSession err: %v", err)
		resp.ToErrorResp(err)
		return
	}

	resp.ToResponse(gin.H{
		"code": 201,
	})
}
func (ci Cinema) DelSession(c *gin.Context) {
	param := service.DelSessionReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	err := app.CheckPermission(c, &svc, param)
	if err != nil {
		resp.ToErrorResponse(errcode.Forbidden)
		return
	}
	err = svc.DelSession(&param)
	if err != nil {
		global.Logger.Errorf("svc.DelSession err: %v", err)
		resp.ToErrorResp(err)
		return
	}

	resp.ToResponse(gin.H{
		"code": 200,
	})
}
