package xray

import (
	"log"

	"github.com/certekim/xctl"
)

var Client = xctl.NewServiceClient("127.0.0.1", 65534)

func init() {
	log.Println("connecting to xray")
}
