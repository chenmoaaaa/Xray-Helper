package xray

import (
	"log"

	"github.com/certekim/xctl"
	"github.com/certekim/xray-helper/app/config"
)

var Client *xctl.ServiceClient

func init() {
	log.Println("creating rpc client")
	Client = xctl.NewServiceClient("127.0.0.1", uint32(config.Conf["xray"].(float64)))
	log.Println("rpc client created")
}
