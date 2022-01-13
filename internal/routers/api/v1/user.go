package v1

import (
	"TicketSales/global"
	"TicketSales/internal/service"
	"TicketSales/pkg/aliyun"
	"TicketSales/pkg/app"
	"TicketSales/pkg/errcode"
	"TicketSales/pkg/uid"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUser() User{
	return User{}
}

func (u User) GetAuth(c *gin.Context){
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	//token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	//if err != nil {
	//	global.Logger.Errorf("app.GenerateToken err: %v", err)
	//	response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	//	return
	//}

	response.ToResponse(gin.H{
		"flag": "success",
	})
}
func (u User) LoginByMobileCaptcha(c *gin.Context) {
	param := service.LoginByCaptchaReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	base,err := svc.LoginByMobileCaptcha(&param)
	if err != nil {
		global.Logger.Errorf("svc.LoginByMobileCaptcha err: %v", err)
		response.ToErrorResp(err)
		return
	}
	response.ToResponse(gin.H{
		"code": 200,
		"data": base,
	})
}

func (u User) CreateAuth(c *gin.Context) {
	param := service.AuthRequest{}
	resp := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	flag := svc.IsSignup(param.Identity)
	if flag {
		resp.ToErrorResponse(errcode.IsSignup)
		return
	}
	uid,err := uid.Sf.NextID(0)
	if err != nil {
		global.Logger.Errorf("app.GenSnowflake err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	err = svc.CreateAuth(param.IdentityType,uid,param.Identity,param.Certificate)
	if err != nil {
		global.Logger.Errorf("svc.CreateAuth err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponse(gin.H{
		"flag": "success",
	})
}

func (u User) SignupByMobile(c *gin.Context) {
	param := service.SignupByCaptchaReq{}
	resp := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	flag,err:=aliyun.CheckMobileCaptcha(param.Identity,param.Captcha)
	if !flag || err != nil{
		global.Logger.Errorf("svc.CheckMobileCaptcha err: %v", err)
		resp.ToErrorResp(err)
		return
	}
	svc := service.New(c.Request.Context())
	uid,err := uid.Sf.NextID(0)
	if err != nil {
		global.Logger.Errorf("app.GenSnowflake err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	err = svc.CreateAuth(1,uid,param.Identity,param.Certificate)
	if err != nil {
		global.Logger.Errorf("svc.CreateAuth err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{
		"flag": "success",
	})


}

func (u User) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
		/*
		required	必填
		gt	大于
		gte	大于等于
		lt	小于
		lte	小于等于
		min	最小值
		max	最大值
		oneof	参数集内的其中之一
		len	长度要求与 len 给定的一致
		*/
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	response.ToResponse(gin.H{})
	return
}

