package stats

import (
	"fmt"
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func StatsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := ps.ByName("name")
	log.Println("stats : " + data)
	_, value := xray.Client.GetStats(data, false)
	fmt.Fprintf(w, "%s\n", fmt.Sprintf("%v", value))
	log.Println(value)
}
