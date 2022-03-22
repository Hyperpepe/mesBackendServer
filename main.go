package main

import (
	Client "first/tcp_Client"
	"fmt"
)

func main() {
	conn := "192.168.2.34:10000"
	go fmt.Print(Client.SendMessage(conn, "status"))
	//EsopSqlServer.TestConnect(t *testing.T)
}
