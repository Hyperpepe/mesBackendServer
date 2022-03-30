package main

import (
	"first/API"
	Server "first/tcp_Server"
)

func main() {
	API.StartApiListen("192.168.2.46:11000")
	Server.StartListen("192.168.2.46:10000")
	//log.Println("------------------------------------------------------------")
}
