package main

import (
	Client "first/tcp_Client"
)

func main() {
	conn := "192.168.2.43:10000"
	Client.SendMessage(conn, "Pic:ftp://ftp@192.168.2.46/home/ftp/mes/04739040-0e30-4ae5-94c3-e1bec2a028cd_1.png")
	//EsopSqlServer.TestConnect(t *testing.T)
}
