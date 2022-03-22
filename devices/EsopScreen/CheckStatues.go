package EsopScreen

import (
	"first/ReadConfig"
	"first/SQL"
	Client "first/tcp_Client"
	"fmt"
	"strings"
)

func CheckStatues() {
	conn := SQL.ConnSQL()
	conf := ReadConfig.ReadConfig()
	defer conn.Close()
	//编写查询语句
	stmt, err := conn.Prepare(`select 设备网络IP from dbo.esop表单`)
	if err != nil {
		fmt.Println("Prepare failed:", err.Error())
		return
	}
	defer stmt.Close()

	//执行查询语句
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("Query failed:", err.Error())
		return
	}
	//将数据读取到实体中
	var Ip []string
	esopPort := (*conf)["esop_port"]
	for rows.Next() {
		//tmp为每行暂存数据
		var tmp string
		rows.Scan(&tmp)
		Ip = append(Ip, tmp)
		checkStatues(tmp+esopPort, "status:")
		fmt.Println("ip:", tmp, " status:")
	}
	fmt.Println(Ip)
}
func checkStatues(Ip string, Status string) {
	ret := Client.SendMessage(Ip, "status:")
	fmt.Println(ret)
	if strings.Contains(ret, "online") {

	} else {

	}
}
