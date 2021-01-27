package inbound

import (
	"log"
	"net/http"

	"github.com/certekim/xray-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func RemoveInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := ps.ByName("tag")
	log.Println("remove inbound : " + data)
	xray.Client.RemoveInbound(data)
	// fmt.Fprintf(w, "Success")
}
