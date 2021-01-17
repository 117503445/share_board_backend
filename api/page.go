package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"gopkg.in/olahol/melody.v1"
	"shareboard/global"
	"shareboard/model"
	"shareboard/service"
	"strings"
)

//BoardGetData 获取 Page 的 Data
func BoardGetData(c *gin.Context) {
	var page model.Page
	id := 1
	model.DB.First(&page, id)
	//fmt.Println(page.Data)
	c.String(200, page.Data)
}

func BoardOnMessage(s *melody.Session, msg []byte) {

	//id := 1
	//model.DB.First(&page, id)
	//
	//if page.ID == 0 {
	//	page.Data = "{\"version\":\"2.1.0\",\"objects\":[]}"
	//	model.DB.Create(&page)
	//	fmt.Println(fmt.Sprintf("page create %v", page.ID))
	//}
	//
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
		page = *service.PageGet(boardID, pageNumber)

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

		model.DB.Save(&page)
		err = global.M.BroadcastOthers(msg, s)
		if err != nil {
			fmt.Println(err)
		}
	case "replace":
		var page model.Page
		boardID := s.Keys["boardID"].(string)
		pageNumber := s.Keys["pageNumber"].(int)
		page = *service.PageGet(boardID, pageNumber)

		page.Data = string(dataBytes)
		//fmt.Println(page.Data)
		model.DB.Save(&page)
		err := global.M.BroadcastOthers(msg, s)
		if err != nil {
			fmt.Println(err)
		}
	case "status":
		data := requestData.(map[string]interface{})
		boardID := data["boardid"].(string)
		pageNumber := int(data["pagenumber"].(float64))

		page := service.PageGet(boardID, pageNumber)
		s.Keys = make(map[string]interface{})
		s.Keys["boardID"] = boardID
		s.Keys["pageNumber"] = pageNumber
		//response, _ := json.Marshal(gin.H{
		//	"code": 1,
		//	"msg":  "ok",
		//	"type": "status",
		//	"data": page.Data})
		response := fmt.Sprintf("{\"code\": 1,\"msg\":  \"ok\",\"type\": \"status\",\"data\": %v}", page.Data)

		_ = global.M.BroadcastFilter([]byte(response), func(q *melody.Session) bool {
			return q == s
		})
	}
}
