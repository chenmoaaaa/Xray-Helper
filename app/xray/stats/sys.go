package stats

import (
	"fmt"
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/julienschmidt/httprouter"
)

func StatssysHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("statsys")
	fmt.Fprintf(w, "%s\n", xray.Client.GetSysStats())
}
