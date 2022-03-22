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
}
