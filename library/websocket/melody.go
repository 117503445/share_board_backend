package websocket

import (
	"gopkg.in/olahol/melody.v1"
	"math"
	"net/http"
)

var M *melody.Melody

func init() {
	M = melody.New()
	M.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	M.Config.MaxMessageSize = math.MaxInt64
}
