package service

import (
	"TicketSales/internal/model"
	"errors"
)

type AuthRequest struct {
	Identity     string `form:"username" binding:"required" json:"username"`
	Certificate  string `form:"password" binding:"required" json:"password"`
	IdentityType uint8  `form:"type" binding:"required" json:"type"`
}

func (svc *Service) CheckAuth(param *AuthRequest) (model.UserAuth, error) {
	auth, err := svc.Dao.GetAuth(
		param.Identity,
		param.Certificate,
		param.IdentityType,
	)
	if err != nil {
		return auth, err
	}

	if auth.ID > 0 {
		return auth, nil
	}

	return auth, errors.New("auth info not exist")
}

func (svc *Service) CreateAuth(uid uint64, req *SignupByMobileReq) error {
	return svc.Dao.CreateAuth(uid, req.Username, req.Mobile, req.Password)
}

func (svc *Service) IsSignup(identity string) bool {
	count, _ := svc.Dao.IsSignup(identity)
	if count > 0 {
		return true
	}
	return false
}
func (svc *Service) GetUserInfo(p interface{}) (model.UserBase, error) {
	return svc.Dao.GetBase(p)
}
func (svc *Service) UserBindCinema(uid, cid uint64) error {
	return svc.Dao.UserBindCinema(uid, cid)
}
