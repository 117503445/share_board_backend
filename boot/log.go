package boot

import (
	"context"
	"encoding/json"
	"fmt"
	"share_board/library/elasticsearch"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

type EsLogWriter struct {
	logger *glog.Logger
}

func (w *EsLogWriter) Write(p []byte) (n int, err error) {
	body, _ := json.Marshal(g.Map{"body": string(p)}) // todo parse log

	req := esapi.IndexRequest{
		Index:   "log",
		Body:    strings.NewReader(string(body)),
		Refresh: "true",
	}

	// Perform the request with the client.
	if _, err := req.Do(context.Background(), elasticsearch.Es); err != nil {
		fmt.Println(err)
	}

	return w.logger.Write(p)
}

func LogBindEs() {
	if g.Cfg().Get("elasticsearch.enabled").(bool) {
		// g.Log().Line().Debug("LogBindEs")
		g.Log().SetWriter(&EsLogWriter{logger: glog.DefaultLogger()})
		// g.Log().Line().Debug("test log")
	}

}
