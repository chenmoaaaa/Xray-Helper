package inbound

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

//"Content-Type: application/json"
func AddInboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Println("AddInboundHandler: ParseForm")
		log.Println(err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params map[string]interface{}
	if err := decoder.Decode(&params); err != nil {
		log.Println("AddInboundHandler: jsonDecode")
		log.Println(err)
		return
	}
	res, err := json.Marshal(params)
	if err != nil {
		log.Println("AddInboundHandler: jsonMarshal")
		log.Println(err)
		return
	}
	xray.Client.AddInbound(res)
	//log.Fprintf(w, "%v\n", string(res))
}
