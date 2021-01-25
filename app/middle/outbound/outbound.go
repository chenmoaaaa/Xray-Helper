package outbound

import (
	"encoding/json"
	"log"

	"github.com/certekim/xray4magisk-helper/app/middle"
	"github.com/certekim/xray4magisk-helper/app/xray"
)

var OutboundObj map[string]interface{}
var OutboundObjs []map[string]interface{}

func ReadAll() []byte {
	records, err := middle.JsonDB.ReadAll("fish")
	if err != nil {
		log.Println("Error", err)
	}
	for _, f := range records {
		if err := json.Unmarshal([]byte(f), &OutboundObj); err != nil {
			log.Println("Error", err)
		}
		OutboundObjs = append(OutboundObjs, OutboundObj)
	}
	ret, err := json.Marshal(OutboundObjs)
	if err != nil {
		log.Println("Error", err)
	}
	return ret
}

func Write(tag string, in map[string]interface{}) {
	if err := middle.JsonDB.Write("Outbound", tag, in); err != nil {
		log.Println("Error", err)
	}
}

func Read(tag string) []byte {
	if err := middle.JsonDB.Read("Outbound", tag, &OutboundObj); err != nil {
		log.Println("Error", err)
	}
	ret, err := json.Marshal(OutboundObj)
	if err != nil {
		log.Println("Error", err)
	}
	return ret
}

func Del(tag string) {
	if err := middle.JsonDB.Delete("Outbound", tag); err != nil {
		log.Println("Error", err)
	}
}

func Apply(tag string) {
	xray.Client.AddOutbound(Read(tag))
}
