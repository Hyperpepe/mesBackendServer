package EsopScreen

import (
	Client "first/tcp_Client"
	_ "github.com/denisenkom/go-mssqldb"
	"testing"
)

func TestConnect(t *testing.T) {
	//var rowsData []screen
	//rowsData = scanIpData()
	//for i,row := range rowsData{
	//	fmt.Print("id: ",i," ip: ",row.Ip,"\n")
	//}
	//========================================
	//err := SendMessageToAll()
	//if err != "" {
	//	log.Println("failed!")
	//}else {
	//	log.Println("ok!")
	//}
	//=========================================
	//CheckStatues()
	//Ip := "192.168.80.4:10001"
	//Server.GetConnection(Ip)
	Client.SendMessage("192.168.2.25:10000", "h123")
}
