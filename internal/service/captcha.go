package service

import (
	"TicketSales/internal/model"
	"TicketSales/pkg/aliyun"
	"TicketSales/pkg/errcode"
	"errors"
)

type CreateMobileCapReq struct {
	ActionType string `json:"action_type" form:"action_type" binding:"required"`
	Mobile     string `json:"mobile" form:"mobile" binding:"required"`
}
type LoginByCaptchaReq struct {
	Identity string `json:"identity" form:"identity" binding:"required"'`
	Captcha  string `json:"captcha" form:"captcha" binding:"required,len=6"`
}
type SignupByMobileReq struct {
	Username string `json:"username" binding:"required,min=5,max=16"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	Captcha  string `json:"captcha" binding:"required,len=6"'`
}
type CheckAccountReq struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
}

func (svc *Service) CreateMobileCaptcha(param *CreateMobileCapReq) (err error) {
	isSignup := svc.IsSignup(param.Mobile)
	switch param.ActionType {
	case "login":
		if !isSignup {
			return errcode.NotSignup
		}
	case "signup":
		if isSignup {
			return errcode.IsSignup
		}
	default:
		return errors.New("invalid params")
	}
	err = aliyun.SendMobileCaptcha(param.Mobile)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) LoginByMobileCaptcha(param *LoginByCaptchaReq) (model.UserBase, error) {
	flag, err := aliyun.CheckMobileCaptcha(param.Identity, param.Captcha)
	if !flag || err != nil {
		return model.UserBase{}, err
	}
	return svc.Dao.GetBase(param.Identity)
}
