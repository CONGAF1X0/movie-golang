package v1

import (
	"TicketSales/global"
	"TicketSales/internal/service"
	"TicketSales/pkg/aliyun"
	"TicketSales/pkg/app"
	"TicketSales/pkg/errcode"
	"TicketSales/pkg/uid"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
}

func NewUser() User {
	return User{}
}

func (u User) GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	auth, err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	ts, err := app.GenerateToken(auth.UID)
	//token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	saveErr := app.CreateAuth(auth.UID, ts)
	if saveErr != nil {
		global.Logger.Errorf("app.CreateAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	response.ToResponse(gin.H{
		"code":   200,
		"uid":    auth.UID,
		"tokens": tokens,
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
	base, err := svc.LoginByMobileCaptcha(&param)
	if err != nil {
		global.Logger.Errorf("svc.LoginByMobileCaptcha err: %v", err)
		response.ToErrorResp(err)
		return
	}

	ts, err := app.GenerateToken(base.UID)
	//token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	saveErr := app.CreateAuth(base.UID, ts)
	if saveErr != nil {
		global.Logger.Errorf("app.CreateAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	response.ToResponse(gin.H{
		"code":   200,
		"uid":    base.UID,
		"tokens": tokens,
	})
}
func (u User) IsAccountExist(c *gin.Context) {
	param := service.CheckAccountReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if svc.IsSignup(param.Username) {
		response.ToErrorResponse(errcode.IsSignup)
		return
	}
	response.ToResponse(gin.H{
		"code": 200,
	})
}
func (u User) Logout(c *gin.Context) {
	au, err := app.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := app.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr = app.DeleteAuth(au.RefreshUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
func (u User) Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.JWTSetting.RefreshSecret), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := app.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := app.GenerateToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		//save the tokens metadata to redis
		saveErr := app.CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}

//func (u User) CreateAuth(c *gin.Context) {
//	param := service.AuthRequest{}
//	resp := app.NewResponse(c)
//	valid, errs := app.BindAndValid(c, &param)
//	if !valid {
//		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
//		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
//		return
//	}
//	svc := service.New(c.Request.Context())
//	flag := svc.IsSignup(param.Identity)
//	if flag {
//		resp.ToErrorResponse(errcode.IsSignup)
//		return
//	}
//	uid, err := uid.Sf.NextID(0)
//	if err != nil {
//		global.Logger.Errorf("app.GenSnowflake err: %v", err)
//		resp.ToErrorResponse(errcode.ServerError)
//		return
//	}
//	err = svc.CreateAuth(param.IdentityType, uid, param.Identity, param.Certificate)
//	if err != nil {
//		global.Logger.Errorf("svc.CreateAuth err: %v", err)
//		resp.ToErrorResponse(errcode.ServerError)
//		return
//	}
//
//	resp.ToResponse(gin.H{
//		"flag": "success",
//	})
//}

func (u User) SignupByMobile(c *gin.Context) {
	param := service.SignupByMobileReq{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if svc.IsSignup(param.Username) {
		resp.ToErrorResponse(errcode.IsSignup)
		return
	}

	flag, err := aliyun.CheckMobileCaptcha(param.Mobile, param.Captcha)
	if !flag || err != nil {
		global.Logger.Errorf("svc.CheckMobileCaptcha err: %v", err)
		resp.ToErrorResp(err)
		return
	}

	uid, err := uid.Sf.NextID(0)
	if err != nil {
		global.Logger.Errorf("app.GenSnowflake err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	err = svc.CreateAuth(uid, &param)
	if err != nil {
		global.Logger.Errorf("svc.CreateAuth err: %v", err)
		resp.ToErrorResponse(errcode.ServerError)
		return
	}
	resp.ToResponse(gin.H{
		"code": 200,
	})

}
func (u User) GetInfo(c *gin.Context) {
	resp := app.NewResponse(c)
	tokenAuth, err := app.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	uid, err := app.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	svc := service.New(c.Request.Context())
	info, err := svc.GetUserInfo(uid)
	if err != nil {
		global.Logger.Errorf("svc.GetInfo err: %v", err)
		resp.ToErrorResp(err)
		return
	}
	resp.ToResponse(gin.H{
		"code": 200,
		"data": info,
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

	response.ToResponse(gin.H{"res": "test"})
	return
}
