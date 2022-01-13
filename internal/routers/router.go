package routers

import (
	"TicketSales/internal/middleware"
	v1 "TicketSales/internal/routers/api/v1"
	"TicketSales/pkg/limiter"
	"github.com/gin-gonic/gin"
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
	r.Use(middleware.Translations())

	user := v1.NewUser()
	captcha := v1.NewCaptcha()
	movie := v1.NewMovie()
	cinema := v1.NewCinema()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("login", user.GetAuth)
		apiv1.POST("signup", user.CreateAuth)
		apiv1.POST("create_mobile_captcha", captcha.CreateMobileCaptcha)
		apiv1.POST("login_by_mobile_captcha", user.LoginByMobileCaptcha)
		apiv1.POST("signup_by_mobile", user.SignupByMobile)
		apiv1.GET("test", user.List)

		m := apiv1.Group("/movie")
		m.GET("get/:id", movie.Get)
		m.GET("get", movie.List)
		m.GET("hot", movie.Hot)

		cin := apiv1.Group("/cinema")
		cin.GET("get", cinema.List)
		cin.GET("get/:id",cinema.Get)
	}

	return r
}
