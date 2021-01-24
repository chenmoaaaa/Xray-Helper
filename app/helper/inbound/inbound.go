package inbound

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/middle/inbound"
	"github.com/certekim/xray4magisk-helper/app/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/xtls/xray-core/infra/conf"
)

func WriteInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, err, params := utils.Convert(r)
	if err != nil {
		log.Println(err)
		return
	}
	var ret *conf.InboundDetourConfig
	err = json.Unmarshal(res, &ret)
	if err != nil {
		log.Println(err)
		return
	}
	if tag, err := params["tag"].(string); err {
		inbound.Write(tag, ret)
		log.Println(err)
		return
	} else {
		log.Println(err)
	}
}

func ReadInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")
	var data *conf.InboundDetourConfig
	err := json.Unmarshal(inbound.Read(tag), &data)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(w, "%s", res)
}
