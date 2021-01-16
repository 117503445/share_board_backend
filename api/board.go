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
	"strings"
)

//BoardGetData 获取 Board 的 Data
func BoardGetData(c *gin.Context) {
	var board model.Board
	id := 1
	model.DB.First(&board, id)
	//fmt.Println(board.Data)
	c.String(200, board.Data)
}

func BoardOnMessage(s *melody.Session, msg []byte) {
	var board model.Board
	id := 1
	model.DB.First(&board, id)

	if board.ID == 0 {
		board.Data = "{\"version\":\"2.1.0\",\"objects\":[]}"
		model.DB.Create(&board)
		fmt.Println(fmt.Sprintf("board create %v", board.ID))
	}

	var mapResult map[string]interface{}
	if err := json.Unmarshal(msg, &mapResult); err != nil {
		fmt.Println(err)
	}

	requestType := mapResult["type"]
	requestData := mapResult["data"]

	dataBytes, _ := json.Marshal(requestData)

	switch requestType {
	case "add":
		oldObjects := gjson.Get(board.Data, "objects").Array()
		strObjects := "["
		for _, o := range oldObjects {
			strObjects += o.String() + ","
		}
		strObjects += string(dataBytes)
		strObjects += "]"

		strObjects, err := sjson.SetRaw(board.Data, "objects", strObjects)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(strObjects)

		js := strings.ReplaceAll(strObjects, "\\\"", "\"")

		//fmt.Println(js)

		board.Data = js

		model.DB.Save(&board)
		err = global.M.BroadcastOthers(msg, s)
		if err != nil {
			fmt.Println(err)
		}
	case "replace":
		board.Data = string(dataBytes)
		//fmt.Println(board.Data)
		model.DB.Save(&board)
		err := global.M.BroadcastOthers(msg, s)
		if err != nil {
			fmt.Println(err)
		}
	}
}
