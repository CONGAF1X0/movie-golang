package routers

import (
	"TicketSales/global"
	"TicketSales/internal/middleware"
	v1 "TicketSales/internal/routers/api/v1"
	"TicketSales/pkg/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/api/v1/test",
	FillInterval: time.Second,
	Capacity:     1,
	Quantum:      1,
})

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use()
	//r.Use(middleware.Translations())

	user := v1.NewUser()
	captcha := v1.NewCaptcha()
	movie := v1.NewMovie()
	cinema := v1.NewCinema()
	actor := v1.NewActor()
	ticket := v1.NewTicket()
	upload := v1.NewUpload()
	r.POST("/upload", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("login", user.GetAuth)
		apiv1.POST("logout", middleware.TokenAuth(), user.Logout)
		apiv1.POST("refresh", user.Refresh)
		apiv1.POST("signup", user.SignupByMobile)
		apiv1.POST("create_mobile_captcha", captcha.CreateMobileCaptcha)
		apiv1.POST("login_by_mobile_captcha", user.LoginByMobileCaptcha)
		apiv1.POST("is_account_exist", user.IsAccountExist)
		apiv1.GET("test", user.List)

		ug := apiv1.Group("/user")
		ug.GET("info/get", middleware.TokenAuth(), user.GetInfo)

		tg := apiv1.Group("/ticket")
		tg.POST("/create", middleware.TokenAuth(), ticket.Create)
		tg.GET("list", middleware.TokenAuth(), ticket.List)

		m := apiv1.Group("/movie")
		m.GET("get/:id", movie.Get)
		m.GET("list", movie.List)
		m.GET("hot", movie.Hot)
		m.GET("runtime", movie.Runtime)

		ag := apiv1.Group("/actor")
		ag.GET("get/:id", actor.Get)
		ag.GET("list", actor.List)
		ag.PUT("/update", middleware.TokenAuth(), actor.Update)
		ag.DELETE("/del", middleware.TokenAuth(), actor.Del)
		ag.POST("create", middleware.TokenAuth(), actor.Create)

		cin := apiv1.Group("/cinema")
		cin.GET("list", cinema.List)
		cin.GET("get/:id", cinema.Get)
		cin.PUT("/update", middleware.TokenAuth(), cinema.Update)
		cin.POST("/create", middleware.TokenAuth(), cinema.Create)

		cin.GET("/hall/list", cinema.HallList)
		cin.POST("/hall/create", middleware.TokenAuth(), cinema.CreateHall)
		cin.DELETE("/hall/del", middleware.TokenAuth(), cinema.DelHall)
		cin.PUT("/hall/update", middleware.TokenAuth(), cinema.UpdateHall)

		cin.GET("/session/list", cinema.SessionList)
		cin.GET("/session/get/:id", cinema.GetSession)
		cin.PUT("/session/update", middleware.TokenAuth(), cinema.UpdateSession)
		cin.POST("/session/create", middleware.TokenAuth(), cinema.CreateSession)
		cin.DELETE("session/del", middleware.TokenAuth(), cinema.DelSession)

		cin.GET("/seat/sold", cinema.SoldSeat)
	}

	return r
}
