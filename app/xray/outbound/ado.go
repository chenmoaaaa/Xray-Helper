package outbound

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func AddOutboundHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("AddOutboundHandler: ParseForm")
		log.Println(err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params map[string]interface{}
	if err := decoder.Decode(&params); err != nil {
		fmt.Println("AddOutboundHandler: jsonDecode")
		log.Println(err)
		return
	}
	res, err := json.Marshal(params)
	if err != nil {
		fmt.Println("AddOutboundHandler: jsonMarshal")
		log.Println(err)
		return
	}
	xray.Client.AddOutbound(res)
	//fmt.Fprintf(w, "%v\n", string(res))
}
