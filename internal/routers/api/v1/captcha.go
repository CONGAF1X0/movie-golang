package v1

import (
	"TicketSales/global"
	"TicketSales/internal/service"
	"TicketSales/pkg/app"
	"TicketSales/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Captcha struct {
}

func NewCaptcha() Captcha {
	return Captcha{}
}

func (cap Captcha) CreateMobileCaptcha(c *gin.Context) {
	param := service.CreateMobileCapReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateMobileCaptcha(&param)
	if err == errcode.NotSignup {
		response.ToErrorResponse(errcode.NotSignup)
		return
	}
	if err == errcode.IsSignup {
		response.ToErrorResponse(errcode.IsSignup)
		return
	}
	if err != nil {
		global.Logger.Errorf("svc.CreateMobileCaptcha err: %v", err)
		response.ToErrorResp(err)
		return
	}
	response.ToResponse(gin.H{
		"code": 200,
	})
}
