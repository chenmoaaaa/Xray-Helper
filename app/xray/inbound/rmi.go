package inbound

import (
	"fmt"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/utils"
	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func RemoveInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ret := ps.ByName("tag")
	data := utils.Extract(ret)
	xray.Client.RemoveInbound(data)
	fmt.Fprintf(w, "Success")
}
