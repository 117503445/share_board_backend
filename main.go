package main

import (
	_ "share_board_backend/boot"
	_ "share_board_backend/router"

	"github.com/gogf/gf/frame/g"
)

// @title       share_board_backend API
// @version     1.0.4
// @description `share_board_backend` 117503445 的 共享白板 api

// @contact.name 117503445
// @contact.url http://www.117503445.top
// @contact.email t117503445@gmail.com

// @license.name GNU GPL 3.0

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

func main() {
	s := g.Server()
	s.SetDumpRouterMap(g.Cfg().GetBool("server.DumpRouterMap"))
	s.Run()
}
