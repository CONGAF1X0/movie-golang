package v1

import (
	"TicketSales/global"
	"TicketSales/internal/service"
	"TicketSales/pkg/app"
	"TicketSales/pkg/errcode"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Movie struct {
}

func NewMovie() Movie {
	return Movie{}
}

func (m Movie) Get(c *gin.Context) {
	strID := c.Param("id")
	id , _:= strconv.Atoi(strID)
	resp := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	data,err:=svc.GetMovie(id)
	if err!=nil {
		global.Logger.Errorf("svc.GetMovie err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{
		"code":errcode.Success.Code(),
		"data":data,
	})
}

func (m Movie) List(c *gin.Context) {
	param := service.GetMovieReq{}
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
	data,err := svc.GetMovieList(&param,offset,size)
	if err != nil {
		global.Logger.Errorf("svc.GetMovieList err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponseList(data,len(data))
}

func (m Movie) Hot(c *gin.Context) {
	param := service.HotMoviesReq{}
	resp := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	data,err := svc.GetHotMovies(param)
	if err != nil {
		global.Logger.Errorf("svc.GetHotMovies err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponseList(data,len(data))
}