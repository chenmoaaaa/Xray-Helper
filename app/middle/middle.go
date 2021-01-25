package middle

import (
	"log"

	"github.com/certekim/xray4magisk-helper/app/config"
	scribble "github.com/nanobox-io/golang-scribble"
)

var JsonDB *scribble.Driver

func init() {
	var err error
	path := config.Conf["data"].(string)
	JsonDB, err = scribble.New(path, nil)
	if err != nil {
		log.Println("storsge pool :")
		log.Println("	Error : ", err)
	} else {
		log.Println("storage pool connected : " + path)
	}
}
