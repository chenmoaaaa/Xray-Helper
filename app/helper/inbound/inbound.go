package inbound

import (
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/middle/inbound"
	"github.com/certekim/xray4magisk-helper/app/utils"
	"github.com/julienschmidt/httprouter"
)

func WriteInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, err, params := utils.Convert2(r)
	if err != nil {
		log.Println(err)
		return
	}
	if tag, err := params["tag"].(string); err {
		inbound.Write(tag, res)
		log.Println(err)
		return
	} else {
		log.Println(err)
	}
}

func ReadInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
