package router

import (
	"share_board_backend/app/api"
	"share_board_backend/app/middleware"
	"share_board_backend/library/websocket"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS, middleware.HttpLog)
		group.ALL("/", api.Hello)
		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.POST("/login", middleware.Auth.LoginHandler)
				group.POST("/", api.User.SignUp)
				group.Middleware(middleware.JWTLogin)
				group.GET("/", api.User.GetInfo)
			})

			group.GET("/ws", func(r *ghttp.Request) {
				if err := websocket.M.HandleRequest(r.Response.Writer, r.Request); err != nil {
					g.Log().Line().Panic(err)
				}
			})
			websocket.M.HandleMessage(api.PageOnMessage)

			group.Group("/task", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Task.ReadAll)
				group.GET("/{id}", api.Task.ReadOne)
				group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
				group.POST("/", api.Task.Create)
				group.DELETE("/{id}", api.Task.Delete)
				group.PUT("/{id}", api.Task.Update)
				group.PATCH("/{id}", api.Task.Update)
			})

			// crud_gen will insert here
		})
	})
}
