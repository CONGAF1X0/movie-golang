package dao

import (
	"TicketSales/internal/model"
	"github.com/jinzhu/gorm"
)

func (d *Dao) GetAuth(identity, certificate string , idType uint8) (model.UserAuth, error) {
	auth := model.UserAuth{Identity: identity, Certificate: certificate, IdentityType: idType}
	return auth.Get(d.engine)
}

func (d *Dao) CreateAuth(identityType uint8,uid uint64, identity, certificate string) error {
	auth := model.UserAuth{
		UID: uid,
		IdentityType: identityType,
		Identity: identity,
		Certificate: certificate,
	}
	base := model.UserBase{
		UID:      uid,
		UserName: identity,
	}
	switch identityType {
	case 1:
		base.Mobile = identity
	case 2:
		base.Email = identity
	}
	d.engine.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := auth.Create(tx); err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := base.Create(tx); err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
	return nil
}

func (d *Dao) FindByIdentity(identity string) (int,error) {
	base:= model.UserBase{
		UserName: identity,
		Mobile: identity,
		Email: identity,
	}
	return base.IsExist(d.engine)
}