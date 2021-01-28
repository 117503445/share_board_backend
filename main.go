package main

import (
	_ "share_board/boot"
	_ "share_board/router"

	"github.com/gogf/gf/frame/g"
)

// @title       share_board API
// @version     1.0.1
// @description `share_board` HaoTian Qi 的 共享白板，基于 Goframe + Websocket

// @contact.name 117503445
// @contact.url http://www.117503445.top
// @contact.email t117503445@gmail.com

// @license.name GNU GPL 3.0

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

func main() {
	g.Server().Run()
}
