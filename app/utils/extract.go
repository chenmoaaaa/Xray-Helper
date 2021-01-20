package utils

import (
	"regexp"
)

var rex = regexp.MustCompile(`(.*[^=])=(.*)`)
var res string

func Extract(ps string) string {
	data := rex.FindAllStringSubmatch(ps, -1)
	// log.Println(data)
	for _, m := range data {
		res = m[2]
	}
	return res
}

func ExtractStats() {}
