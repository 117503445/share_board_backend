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
		strokeCreate(s, msg)
	case "strokes-delete":
		strokesDelete(s, msg)
	case "strokes-clear":
		strokesClear(s, requestBody)
	case "change-page-index":
		changePageIndex(s, requestBody)
	}

}

func strokeCreate(s *melody.Session, msg []byte) {
	broadcastToOther(s, msg)

}

func strokesDelete(s *melody.Session, msg []byte) {
	broadcastToOther(s, msg)

}

func strokesClear(s *melody.Session, requestBody map[string]interface{}) {

}

func changePageIndex(s *melody.Session, requestBody map[string]interface{}) {

	boardId := requestBody["boardId"].(string)
	pageId := requestBody["pageId"].(int)

	s.Keys = g.MapStrAny{"boardID": boardId, "pageNumber": pageId}
	//s.Keys["boardID"] = boardId
	//s.Keys["pageNumber"] = pageId

	// todo 返回当前页的所有笔迹

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
