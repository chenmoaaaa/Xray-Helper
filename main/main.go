package main

import (
	"net/http"

	"certe.kim/xray4magisk-helper/app/xray/stats"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/statssys", stats.StatssysHandler)
	http.ListenAndServe(":8080", nil)
}
