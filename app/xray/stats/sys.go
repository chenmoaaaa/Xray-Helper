package stats

import (
	"fmt"
	"log"
	"net/http"

	"certe.kim/xray4magisk-helper/app/utils"
)

func StatssysHandler(w http.ResponseWriter, r *http.Request) {
	err, out, errout := utils.Shellout("xray api statssys --server=127.0.0.1:65533")
	if err != nil {
		log.Printf("error when running command: %v\n", err)
	}
	if errout != "" {
		log.Printf("stderr: %v\n", err)
	}
	fmt.Fprintf(w, string(out))
}
