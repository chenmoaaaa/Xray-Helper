package stats

import (
	"fmt"
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/utils"
	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func StatsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ret := ps.ByName("name")
	data := utils.Extract(ret)
	_, value := xray.Client.GetStats(data, false)
	fmt.Fprintf(w, "%s\n", fmt.Sprintf("%v", value))
	log.Println(value)
}
