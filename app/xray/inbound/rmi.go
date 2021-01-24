package inbound

import (
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func RemoveInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := ps.ByName("tag")
	xray.Client.RemoveInbound(data)
	// fmt.Fprintf(w, "Success")
}
