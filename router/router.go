package router

import (
	"share_board/app/api"
	"share_board/app/middleware"
	"share_board/library/websocket"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func middlewareAuth(r *ghttp.Request) {
	middleware.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Hello)
		group.Group("/api", func(group *ghttp.RouterGroup) {

			group.GET("/ws", func(r *ghttp.Request) {
				if err := websocket.M.HandleRequest(r.Response.Writer, r.Request); err != nil {
					g.Log().Line().Panic(err)
				}
			})
			websocket.M.HandleMessage(api.PageOnMessage)

			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.POST("/login", middleware.Auth.LoginHandler)
				group.POST("/", api.User.SignUp)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middlewareAuth)
					group.GET("/", api.User.GetInfo)
				})
			})
		})
	})
}
