package main

import (
	"log"
	"net/http"

	hin "github.com/certekim/xray4magisk-helper/app/helper/inbound"
	xin "github.com/certekim/xray4magisk-helper/app/xray/inbound"
	xout "github.com/certekim/xray4magisk-helper/app/xray/outbound"
	xstat "github.com/certekim/xray4magisk-helper/app/xray/stats"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.FileServer(http.Dir(("./static"))).ServeHTTP(w, r)
	})
	router.GET("/xray/statssys", xstat.StatssysHandler)
	router.GET("/xray/statsquery/:pattern", xstat.StatsqueryHandler)
	router.GET("/xray/stats/:name", xstat.StatsHandler)
	router.GET("/xray/inbound/remove/:tag", xin.RemoveInboundHandler)
	router.POST("/xray/inbound/add", xin.AddInboundHandler)
	router.GET("/xray/outbound/remove/:tag", xout.RemoveOutboundHandler)
	router.POST("/xray/outbound/add", xout.AddOutboundHandler)
	router.POST("/helper/inbound/add", hin.WriteInboundHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
