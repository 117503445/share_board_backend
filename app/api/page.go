package api

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"gopkg.in/olahol/melody.v1"
	"share_board/app/dao"
	"share_board/app/model"
	"share_board/app/service"
	"share_board/library/websocket"
	"strings"
)

func PageOnMessage(s *melody.Session, msg []byte) {
	var mapResult map[string]interface{}
	if err := json.Unmarshal(msg, &mapResult); err != nil {
		fmt.Println(err)
	}

	requestType := mapResult["type"]
	requestData := mapResult["data"]

	dataBytes, _ := json.Marshal(requestData)

	switch requestType {
	case "add":
		var page model.Page
		boardID := s.Keys["boardID"].(string)
		pageNumber := s.Keys["pageNumber"].(int)
		page = *service.GetPageByBoardIDPageNumber(boardID, pageNumber)

		oldObjects := gjson.Get(page.Data, "objects").Array()
		strObjects := "["
		for _, o := range oldObjects {
			strObjects += o.String() + ","
		}
		strObjects += string(dataBytes)
		strObjects += "]"

		strObjects, err := sjson.SetRaw(page.Data, "objects", strObjects)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(strObjects)

		js := strings.ReplaceAll(strObjects, "\\\"", "\"")

		//fmt.Println(js)

		page.Data = js

		if _, err := dao.Page.Save(page); err != nil {
			g.Log().Line().Error(err)
		}
		//model.DB.Save(&page)

		err = websocket.M.BroadcastFilter(msg, func(q *melody.Session) bool {
			isOk := q != s && boardID == q.Keys["boardID"].(string) && pageNumber == q.Keys["pageNumber"].(int) // 发送给其他 同白板 同页 的客户端
			return isOk
		})
		if err != nil {
			fmt.Println(err)
		}
	case "replace":
		var page model.Page
		boardID := s.Keys["boardID"].(string)
		pageNumber := s.Keys["pageNumber"].(int)
		page = *service.GetPageByBoardIDPageNumber(boardID, pageNumber)

		page.Data = string(dataBytes)
		//fmt.Println(page.Data)
		if _, err := dao.Page.Save(page); err != nil {
			g.Log().Line().Error(err)
		}

		err := websocket.M.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q != s && boardID == q.Keys["boardID"].(string) && pageNumber == q.Keys["pageNumber"].(int) // 发送给其他 同白板 同页 的客户端
		})

		if err != nil {
			fmt.Println(err)
		}
	case "status":
		data := requestData.(map[string]interface{})
		boardID := data["boardid"].(string)
		pageNumber := int(data["pagenumber"].(float64))

		page :=service.GetPageByBoardIDPageNumber(boardID, pageNumber)
		s.Keys = make(map[string]interface{})
		s.Keys["boardID"] = boardID
		s.Keys["pageNumber"] = pageNumber
		//response, _ := json.Marshal(gin.H{
		//	"code": 1,
		//	"msg":  "ok",
		//	"type": "status",
		//	"data": page.Data})
		response := fmt.Sprintf("{\"code\": 1,\"msg\":  \"ok\",\"type\": \"status\",\"data\": %v}", page.Data)

		_ = websocket.M.BroadcastFilter([]byte(response), func(q *melody.Session) bool {
			return q == s
		})
	}
}
