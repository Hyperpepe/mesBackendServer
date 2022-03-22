package EsopScreen

import (
	"database/sql"
	Client "first/tcp_Client"
	"fmt"
	"github.com/astaxie/beego/config"
	"log"
)

//定义从数据库取数据的类型
type screen struct {
	Ip    string
	Image string
}

func SendMessageToAll() bool {
	//声明esop的IP和data
	var rowsData []screen
	//声明数据库连接字符串
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
		log.Println("Sql Open Connection failed:", err.Error())
		return false
	}
	defer conn.Close()
	//编写查询语句
	stmt, err := conn.Prepare(`select 设备网络IP,显示图片 from dbo.esop表单`)
	if err != nil {
		log.Println("Sql Prepare failed:", err.Error())
		return false
	}
	defer stmt.Close()

	//执行查询语句
	rows, err := stmt.Query()
	if err != nil {
		log.Println("Query failed:", err.Error())
		return false
	}
	//将数据读取到实体中
	for rows.Next() {
		var row screen
		rows.Scan(&row.Ip, &row.Image)
		rowsData = append(rowsData, row)
	}
	//读取不到信息则返回空
	if rowsData == nil {
		log.Println("Can not get Data,Please Check DataSources!!!")
		return false
	}
	//读取到信息则通过tcp传递信息
	esopPort := conf.String("esop_port")
	for _, row := range rowsData {
		//fmt.Println("Now send to device ",i,row.Ip+":10000","Pic:"+row.Image)
		log.Println(row.Ip + esopPort)
		log.Println("Pic:" + row.Image)
		go Client.SendMessage(row.Ip+esopPort, "Pic:"+row.Image)
	}
	return true
}
