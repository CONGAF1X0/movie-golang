package dao

import (
	"TicketSales/internal/model"
	"github.com/jinzhu/gorm"
)

func (d *Dao) GetAuth(identity, certificate string, idType uint8) (model.UserAuth, error) {
	auth := model.UserAuth{Identity: identity, Certificate: certificate, IdentityType: idType}
	return auth.Get(d.engine)
}

func (d *Dao) CreateAuth(uid uint64, username, mobile, certificate string) error {
	auth1 := model.UserAuth{
		UID:          uid,
		IdentityType: 1,
		Identity:     mobile,
		Certificate:  certificate,
	}
	auth2 := model.UserAuth{
		UID:          uid,
		Identity:     username,
		IdentityType: 2,
		Certificate:  certificate,
	}
	base := model.UserBase{
		UID:      uid,
		UserName: username,
		Mobile:   mobile,
	}
	d.engine.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := auth1.Create(tx); err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		if err := auth2.Create(tx); err != nil {
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

func (d *Dao) IsSignup(identity string) (int, error) {
	return model.UserAuth{Identity: identity}.Exist(d.engine)
}
