package main

import (
	"first/API"
	"first/ReadConfig"
	Server "first/tcp_Server"
)

func main() {
	conf := ReadConfig.ReadConfig()
	ApiIp := (*conf)["API_ListenAddr"]
	TcpIp := (*conf)["TCP_ListenAddr"]
	API.StartApiListen(ApiIp)
	Server.StartListen(TcpIp)
	//log.Println("------------------------------------------------------------")
}
