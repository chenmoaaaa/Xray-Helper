package stats

import (
	"fmt"
	"log"
	"net/http"

	"certe.kim/xray4magisk-helper/app/utils"
	"github.com/julienschmidt/httprouter"
)

func StatsqueryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pattern := ps.ByName("pattern")
	value := utils.Extract(pattern)
	// log.Println(value)
	// log.Println("xray api statsquery --server=127.0.0.1:65534 -pattern \"" + value + "\"")
	err, out, errout := utils.Shellout("xray api statsquery --server=127.0.0.1:65534 -pattern \"" + value + "\"")
	if err != nil {
		log.Printf("error when running command: %v\n", err)
	}
	if errout != "" {
		log.Printf("stderr: %v\n", err)
	}
	fmt.Fprintf(w, "%s\n", string(out))
}
