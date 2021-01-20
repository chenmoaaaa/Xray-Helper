package stats

import (
	"fmt"
	"log"
	"net/http"

	"certe.kim/xray4magisk-helper/app/utils"
	"github.com/julienschmidt/httprouter"
)

func StatsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	value := utils.Extract(name)
	// log.Println(value)
	// log.Println("xray api stats --server=127.0.0.1:65534 -name \"" + value + "\"")
	err, out, errout := utils.Shellout("xray api stats --server=127.0.0.1:65534 -name \"" + value + "\"")
	if err != nil {
		log.Printf("error when running command: %v\n", err)
	}
	if errout != "" {
		log.Printf("stderr: %v\n", err)
	}
	fmt.Fprintf(w, "%s\n", string(out))
}
