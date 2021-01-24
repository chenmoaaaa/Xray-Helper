package middle

import (
	"log"

	scribble "github.com/nanobox-io/golang-scribble"
)

var JsonDB *scribble.Driver

const dir = "/data/xray/helper/data"

func Init() {
	var err error
	JsonDB, err = scribble.New(dir, nil)
	if err != nil {
		log.Println("Error", err)
	}
	log.Println("storage pool connected")
}
