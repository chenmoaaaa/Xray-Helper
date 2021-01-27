package outbound

import (
	"log"
	"net/http"

	"github.com/certekim/xray-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func RemoveOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := ps.ByName("tag")
	log.Println("remove outbound : " + data)
	xray.Client.RemoveOutbound(data)
	// fmt.Fprintf(w, "Success")
}
