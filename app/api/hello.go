package api

import (
	"share_board_backend/library/response"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

// @summary 返回 Hello
// @tags    hello
// @produce json
// @router  / [GET]
// @success 200 {object} response.JsonResponse
func Hello(r *ghttp.Request) {
	response.Json(r, 0, "hello", g.Map{"time": gtime.New(time.Now())})
}
