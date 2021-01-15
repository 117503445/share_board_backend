package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"math"
)

func main() {
	r := gin.Default()
	m := melody.New()
	m.Config.MaxMessageSize = math.MaxInt64
	r.GET("/ws", func(c *gin.Context) {
		err := m.HandleRequest(c.Writer, c.Request)
		if err != nil {
			fmt.Println(err)
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		err := m.BroadcastOthers(msg, s)
		if err != nil {
			fmt.Println(err)
		}
	})

	r.Run(":10000")
}
