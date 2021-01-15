package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"math"
	"shareboard/api"
	"shareboard/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	m := melody.New()
	m.Config.MaxMessageSize = math.MaxInt64

	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		user := v1.Group("user")
		{
			// 用户注册
			user.POST("", api.UserCreate)

			// 用户登录
			user.POST("login", middleware.JwtMiddleware.LoginHandler)

			auth := user.Group("")
			// 需要登录保护
			auth.Use(middleware.JwtMiddleware.MiddlewareFunc())
			{
				// User Routing
				auth.GET("me", api.UserRead)
				auth.PUT("", api.UserUpdate)
			}
		}
		admin := v1.Group("admin")
		// 需要登录保护
		admin.Use(middleware.JwtMiddleware.MiddlewareFunc())
		admin.Use(middleware.HasRole("admin"))
		{
			user := admin.Group("user")
			{
				user.GET(":id", api.AdminUserRead)
			}
		}

		v1.GET("/ws", func(c *gin.Context) {
			err := m.HandleRequest(c.Writer, c.Request)
			if err != nil {
				fmt.Println(err)
			}
		})
		m.HandleMessage(func(s *melody.Session, msg []byte) {
			err := m.BroadcastOthers(msg, s)
			if err != nil {
				fmt.Println(err)
			}
		})

	}
	return r
}