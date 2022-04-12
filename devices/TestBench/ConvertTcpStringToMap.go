package TestBench

import (
	"strings"
)

func convert(str string) *map[string]string {
	messages := strings.Split(str, "--")
	var ret map[string]string
	ret = make(map[string]string)
	for _, paraMess := range messages {
		//log.Println(paraId,paraMess)
		if paraMess == "" {
		} else {
			message := strings.SplitN(paraMess, ":", 2)
			//log.Print(message[1])
			ret[message[0]] = message[1]
		}
	}
	return &ret
}
