package app

import (
	"TicketSales/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

func CheckPermission(c *gin.Context, svc *service.Service, param interface{}) error {
	rv := reflect.ValueOf(param)
	cid := rv.FieldByName("CinemaID")
	user, _ := svc.GetUserInfo(c.GetUint64("uid"))
	if cid.Uint() != user.CinemaID {
		return errors.New("forbidden")
	}
	return nil
}
