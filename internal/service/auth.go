package service

import "errors"

type AuthRequest struct {
	Identity     string `form:"identity" binding:"required"`
	Certificate  string `form:"certificate" binding:"required"`
	IdentityType uint8 `form:"identity_type" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.Dao.GetAuth(
		param.Identity,
		param.Certificate,
		param.IdentityType,
	)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}

func (svc *Service) CreateAuth(identityType uint8,uid uint64,identity,certificate string) error {
	return svc.Dao.CreateAuth(identityType, uid, identity, certificate)
}

func (svc *Service) IsSignup(identity string) bool {
	count, _ := svc.Dao.FindByIdentity(identity)
	if count>0 {
		return true
	}
	return false
}