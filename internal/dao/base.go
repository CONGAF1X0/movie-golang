package dao

import (
	"TicketSales/internal/model"
	"strconv"
)

func (d Dao) GetBase(str string) (model.UserBase,error){
	uid,_:=strconv.ParseUint(str,10,64)
	var base = model.UserBase{
		UID: uid,
		Mobile: str,
		Email: str,
		UserName: str,
	}
	return	base.Get(d.engine)
}
