package outbound

import (
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/middle/outbound"
	"github.com/certekim/xray4magisk-helper/app/utils"
	"github.com/julienschmidt/httprouter"
)

func WriteOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, err, params := utils.Convert(r)
	if err != nil {
		log.Println(err)
		return
	}
	if tag, err := params["tag"].(string); err {
		outbound.Write(tag, res)
	} else {
		log.Println(err)
	}
}
