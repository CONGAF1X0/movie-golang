package service

import (
	"TicketSales/internal/model"
	"TicketSales/pkg/aliyun"
	"TicketSales/pkg/errcode"
)

type CreateMobileCapReq struct {
	ActionType string `form:"action_type" binding:"required"`
	Mobile     string `form:"mobile" binding:"required"`
}
type LoginByCaptchaReq struct {
	Identity string `form:"identity" binding:"required"'`
	Captcha string `form:"captcha" binding:"required,len=6"`
}
type SignupByCaptchaReq struct {
	LoginByCaptchaReq
	Certificate string `form:"certificate" binding:"required,min=6,max=32"`
}

func (svc *Service) CreateMobileCaptcha(param *CreateMobileCapReq) (err error){
	isSignup:=svc.IsSignup(param.Mobile)
	switch param.ActionType {
	case "login":
		if isSignup{
			err = aliyun.SendMobileCaptcha(param.Mobile)
			if err != nil {
				return err
			}
		} else {
			return errcode.NotSignup
		}
	case "signup":
		if isSignup{
			return errcode.IsSignup
		} else {
			err = aliyun.SendMobileCaptcha(param.Mobile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (svc *Service) LoginByMobileCaptcha(param *LoginByCaptchaReq) (model.UserBase,error) {
	flag,err := aliyun.CheckMobileCaptcha(param.Identity,param.Captcha)
	if !flag || err!=nil {
		return model.UserBase{},err
	}
	return svc.Dao.GetBase(param.Identity)
}