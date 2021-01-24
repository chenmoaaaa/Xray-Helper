package outbound

import (
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/utils"
	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func AddOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, err := utils.Convert(r)
	if err != nil {
		log.Println(err)
		return
	}
	xray.Client.AddOutbound(res)
	//log.Fprintf(w, "%v\n", string(res))
}
