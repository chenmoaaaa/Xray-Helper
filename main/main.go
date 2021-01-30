package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/certekim/xray-helper/app/config"
	hin "github.com/certekim/xray-helper/app/helper/inbound"
	hout "github.com/certekim/xray-helper/app/helper/outbound"
	xin "github.com/certekim/xray-helper/app/xray/inbound"
	xout "github.com/certekim/xray-helper/app/xray/outbound"
	xstat "github.com/certekim/xray-helper/app/xray/stats"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	dir := config.Conf["dir"].(string)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("Specific folder not found : " + dir)
		log.Println("Serve static website on : ./xray-webui")
		router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			http.ServeFile(w, r, "xray-webui/index.html")
		})
		router.ServeFiles("/helper/*filepath", http.Dir("xray-webui"))
	} else {
		log.Println("Server static website on : " + dir)
		router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			http.ServeFile(w, r, dir+"/index.html")
		})
		router.ServeFiles("/helper/*filepath", http.Dir(dir))
	}
	router.GET("/api/xray/statssys", xstat.StatssysHandler)
	router.GET("/api/xray/statsquery", xstat.StatsqueryHandler)
	router.GET("/api/xray/statsquery/:pattern", xstat.StatsqueryHandler)
	router.GET("/api/xray/stats/:name", xstat.StatsHandler)
	router.GET("/api/xray/inbound/remove/:tag", xin.RemoveInboundHandler)
	router.POST("/api/xray/inbound/add", xin.AddInboundHandler)
	router.GET("/api/xray/outbound/remove/:tag", xout.RemoveOutboundHandler)
	router.POST("/api/xray/outbound/add", xout.AddOutboundHandler)
	router.POST("/api/helper/inbound/add", hin.WriteInboundHandler)
	router.GET("/api/helper/inbound/read/:tag", hin.ReadInboundHandler)
	router.GET("/api/helper/inbound/delete/:tag", hin.DeleteInboundHandler)
	router.GET("/api/helper/inbound/apply/:tag", hin.ApplyInboundHandler)
	router.POST("/api/helper/outbound/add", hout.WriteOutboundHandler)
	router.GET("/api/helper/outbound/read/:tag", hout.ReadOutboundHandler)
	router.GET("/api/helper/outbound/delete/:tag", hout.DeleteOutboundHandler)
	router.GET("/api/helper/outbound/apply/:tag", hout.ApplyOutboundHandler)
	port := strconv.FormatFloat(config.Conf["port"].(float64), 'f', 0, 64)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func init() {
	log.Println("Started Xray-helper")
}
