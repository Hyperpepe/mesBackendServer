package main

import (
	"first/httpServer"
	"first/readConfig"
	Server "first/tcpServer"
)

func main() {
	//服务器开启
	conf := readConfig.ReadConfig()
	httpServer.StartApiListen(conf)
	Server.StartListen(conf)
}
