package api

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"gopkg.in/olahol/melody.v1"
	"share_board_backend/library/websocket"
	"strings"
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

	pageKey := fmt.Sprintf("%v-%v", s.Keys["boardID"], s.Keys["pageNumber"])
	//g.Log().Line().Debug(pageKey)

	var requestBody map[string]interface{}
	_ = json.Unmarshal(msg, &requestBody)

	strokeData, _ := json.Marshal(requestBody["data"])

	strokeKey := requestBody["data"].(map[string]interface{})["startTime"].(float64)
	//g.Log().Line().Debug(strokeKey)
	if _, err := g.Redis().Do("HMSET", pageKey, strokeKey, string(strokeData)); err != nil {
		g.Log().Line().Debug(err)
	}

}

func strokesDelete(s *melody.Session, msg []byte) {
	broadcastToOther(s, msg)

	pageKey := fmt.Sprintf("%v-%v", s.Keys["boardID"], s.Keys["pageNumber"])

	j, _ := gjson.DecodeToJson(msg)
	deleteIdArray := j.GetFloats("id")
	//g.Log().Line().Debug(pageKey)
	//g.Log().Line().Debug(deleteIdArray)
	for _, deleteId := range deleteIdArray {
		g.Log().Line().Debug(deleteId)
		_, _ = g.Redis().Do("HDEL", pageKey, deleteId)
	}

}

func strokesClear(s *melody.Session, requestBody map[string]interface{}) {

}

func changePageIndex(s *melody.Session, requestBody map[string]interface{}) {

	boardId := requestBody["boardId"].(string)
	pageId := requestBody["pageId"].(float64)

	// todo 验证 boardId 是否合法

	if pageId <= 0 {
		return
	}
	s.Keys = g.MapStrAny{"boardID": boardId, "pageNumber": pageId}

	pageKey := fmt.Sprintf("%v-%v", boardId, pageId)

	strokes, err := g.Redis().DoVar("HVALS", pageKey)
	if err != nil {
		g.Log().Line().Debug(err)
	}

	response := fmt.Sprintf("{\"code\": 0,\"msg\":  \"ok\",\"route\": \"strokes-set\",\"data\": [%v]}", strings.Join(strokes.Strings(), ","))

	_ = websocket.M.BroadcastFilter([]byte(response), func(q *melody.Session) bool {
		return q == s
	})

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
