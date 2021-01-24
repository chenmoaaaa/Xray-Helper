package inbound

import (
	"encoding/json"
	"log"

	"github.com/certekim/xray4magisk-helper/app/middle"
	"github.com/certekim/xray4magisk-helper/app/xray"
	"github.com/xtls/xray-core/infra/conf"
)

var InboundObj *conf.InboundDetourConfig

// var InboundObj map[string]interface{}
var InboundObjs []*conf.InboundDetourConfig

//var InboundObjs []map[string]interface{}

func ReadAll() []byte {
	records, err := middle.JsonDB.ReadAll("fish")
	if err != nil {
		log.Println("Error", err)
	}
	for _, f := range records {
		if err := json.Unmarshal([]byte(f), &InboundObj); err != nil {
			log.Println("Error", err)
		}
		InboundObjs = append(InboundObjs, InboundObj)
	}
	ret, err := json.Marshal(InboundObjs)
	if err != nil {
		log.Println("Error", err)
	}
	return ret
}

func Write(tag string, in *conf.InboundDetourConfig) {
	if err := middle.JsonDB.Write("Inbound", tag, in); err != nil {
		log.Println("Error", err)
	}
}

func Read(tag string) []byte {
	if err := middle.JsonDB.Read("Inbound", tag, &InboundObj); err != nil {
		log.Println("Error", err)
	}
	ret, err := json.Marshal(InboundObj)
	if err != nil {
		log.Println("Error", err)
	}
	return ret
}

func Del(tag string) {
	if err := middle.JsonDB.Delete("Inbound", tag); err != nil {
		log.Println("Error", err)
	}
}

func Apply(tag string) {
	xray.Client.AddInbound(Read(tag))
}
