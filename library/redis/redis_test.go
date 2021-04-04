package redis

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestRedis(t *testing.T) {
	g.Redis().Do("SET", "k", "v")
	v, _ := g.Redis().DoVar("GET", "k")
	g.Log().Line().Debug(v)
	g.Redis().Do("DEL", "k")
	v, _ = g.Redis().DoVar("GET", "k")
	g.Log().Line().Debug(v.IsNil())
}
