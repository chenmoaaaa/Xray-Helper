package main

import (
	"log"
	"net/http"

	"github.com/certekim/xray4magisk-helper/app/xray/stats"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.FileServer(http.Dir(("./static"))).ServeHTTP(w, r)
	})
	router.GET("/statssys", stats.StatssysHandler)
	router.GET("/statsquery/:pattern", stats.StatsqueryHandler)
	router.GET("/stats/:name", stats.StatsHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
