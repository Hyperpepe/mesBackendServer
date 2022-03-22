package EsopScreen

import (
	"database/sql"
	Client "first/tcp_Client"
	"fmt"
	"github.com/astaxie/beego/config"
	"log"
	"strings"
)

func CheckStatues() {
	conf, err := config.NewConfig("ini", "C:/Users/tss05/goProject/first/config.conf")
	if err != nil {
		log.Print("config read error!")
		log.Println(err)
	}
	server, port, database, user, password := conf.String("sql_server"),
		conf.String("sql_port"),
		conf.String("sql_database"),
		conf.String("sql_user"),
		conf.String("sql_pass")
	//编写连接字符串
	connString := fmt.Sprintf("server=%s;port%s;database=%s;user id=%s;password=%s", server, port, database, user, password)
	//建立数据库连接：conn
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Println("Open Connection failed:", err.Error())
		return
	}
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
	for rows.Next() {
		var tmp string
		rows.Scan(&tmp)
		Ip = append(Ip, tmp)
		checkStatues(tmp+":10000", "status:")
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
