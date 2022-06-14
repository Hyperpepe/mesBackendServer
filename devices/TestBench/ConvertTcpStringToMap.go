package TestBench

import (
	"errors"
	"strings"
)

// Convert 将TCP传递来的STRING转换为MAP对
func Convert(str string) (*map[string]string, error) {
	//将接受到的字符串转化为MAP
	messages := strings.Split(str, "--")
	var ret map[string]string
	ret = make(map[string]string)
	for _, paraMess := range messages {
		if paraMess == "" {
		} else {
			message := strings.SplitN(paraMess, ":", 2)
			if len(message[0]) == 0 || len(message[1]) == 0 {
				return nil, errors.New("主机传递的字符串包含空信息，请检查主机传递的信息！！！" + str)
			}
			ret[message[0]] = message[1]
		}
	}
	return &ret, nil
}
