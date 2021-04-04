package api

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"gopkg.in/olahol/melody.v1"
	"share_board_backend/library/websocket"
)

func PageOnMessage(s *melody.Session, msg []byte) {
	var requestBody map[string]interface{}

	if err := json.Unmarshal(msg, &requestBody); err != nil {
		g.Log().Line().Debug(err)
		return
	}

	route := requestBody["route"]
	switch route {
	case "stroke-create":
		broadcastToOther(s, msg)
	case "strokes-delete":
		broadcastToOther(s, msg)
	case "strokes-clear":
		strokesClear(s, requestBody)
	case "change-page-index":
		changePageIndex(s, requestBody)
	}

}

//broadcastToOther 转发消息给其他的 Page
func broadcastToOther(s *melody.Session, msg []byte) {
	err := websocket.M.BroadcastFilter(msg, func(q *melody.Session) bool {
		return q != s
	})

	if err != nil {
		g.Log().Line().Debug(err)
	}
}

func strokeCreate(s *melody.Session, msg []byte) {

}

func strokesDelete(s *melody.Session, msg []byte) {

}

func strokesClear(s *melody.Session, requestBody map[string]interface{}) {

}

func changePageIndex(s *melody.Session, requestBody map[string]interface{}) {

}
