package xray

import (
	"log"

	"github.com/certekim/xctl"
)

const port = 65534

var Client = xctl.NewServiceClient("127.0.0.1", port)

func init() {
	log.Println("rpc Client created")
}
