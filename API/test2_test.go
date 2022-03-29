package API

import (
	Server "first/tcp_Server"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"testing"
)

//测试程序
func Test2(t *testing.T) {
	startApiListen()
	Server.StartListen("192.168.2.25:10000")
	log.Println("------------------------------------------------------------")
	//EsopScreen.SendMessageToAll()

	//conn := "192.168.2.34:10000"
	//Client.SendMessage(conn, "Pic:ftp://ftp@192.168.2.46/home/ftp/mes/9f13d2ea-ac5e-40b2-94d9-79fe584dcc68_12.png")
}
