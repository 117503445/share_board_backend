package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"share_board_backend/app/api"
	"share_board_backend/app/middleware"
	"share_board_backend/library/websocket"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS, middleware.HttpLog)
		group.ALL("/", api.Hello)
		group.Group("/api", func(group *ghttp.RouterGroup) {

			group.GET("/ws", func(r *ghttp.Request) {
				g.Log().Line().Debug("LOG /api/ws")
				if err := websocket.M.HandleRequest(r.Response.Writer, r.Request); err != nil {
					g.Log().Line().Debug(err)
				}
			})
			websocket.M.HandleMessage(api.PageOnMessage)

			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.POST("/login", middleware.Auth.LoginHandler)
				group.POST("/", api.User.SignUp)
				group.Middleware(middleware.JWTLogin)
				group.GET("/", api.User.GetInfo)
			})

			// crud_gen will insert here
		})
	})
}
