package outbound

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func AddOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Println("AddOutboundHandler: ParseForm")
		log.Println(err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params map[string]interface{}
	if err := decoder.Decode(&params); err != nil {
		log.Println("AddOutboundHandler: jsonDecode")
		log.Println(err)
		return
	}
	res, err := json.Marshal(params)
	if err != nil {
		log.Println("AddOutboundHandler: jsonMarshal")
		log.Println(err)
		return
	}
	xray.Client.AddOutbound(res)
	//log.Fprintf(w, "%v\n", string(res))
}
