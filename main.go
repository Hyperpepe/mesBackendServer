package main

import (
	"first/API"
	"first/ReadConfig"
	Server "first/tcp_Server"
)

func main() {
	//服务器开启
	conf := ReadConfig.ReadConfig()
	API.StartApiListen(conf)
	Server.StartListen(conf)
}
