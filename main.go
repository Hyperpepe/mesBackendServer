package main

import (
	"first/API"
	Server "first/tcp_Server"
)

func main() {
	API.StartApiListen("172.20.10.2:11000")
	Server.StartListen("172.20.10.2:10000")
	//log.Println("------------------------------------------------------------")
}
