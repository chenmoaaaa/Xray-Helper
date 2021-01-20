package stats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/utils"
	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func StatsqueryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ret := ps.ByName("pattern")
	data := utils.Extract(ret)
	pattern := xray.Client.QueryStats(data, false)

	sub := make(map[string]interface{})
	for key, val := range pattern {
		sub[key] = val
	}
	var stat []map[string]interface{}
	stat = append(stat, sub)

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	err := enc.Encode(&stat)
	if err != nil {
		log.Println("failed to convert json")
	}
	fmt.Fprintf(w, "%s", buf.String())
}
