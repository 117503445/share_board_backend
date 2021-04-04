package api

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"gopkg.in/olahol/melody.v1"
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
		strokeCreate(s, requestBody)
	case "strokes-delete":
		strokesDelete(s, requestBody)
	case "strokes-clear":
		strokesClear(s, requestBody)
	case "change-page-index":
		changePageIndex(s, requestBody)
	}

}

func strokeCreate(s *melody.Session, requestBody map[string]interface{}) {

}

func strokesDelete(s *melody.Session, requestBody map[string]interface{}) {

}

func strokesClear(s *melody.Session, requestBody map[string]interface{}) {

}

func changePageIndex(s *melody.Session, requestBody map[string]interface{}) {

}
