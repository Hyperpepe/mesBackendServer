package EsopScreen

import (
	"errors"
	"first/SQL"
	"first/readConfig"
	Client "first/tcpClient"
	"log"
	"sync"
)

//定义从数据库取数据的类型
type screen struct {
	Ip    string
	Image string
}

//
func SendMessageToAll() error {
	//声明esop的IP和data
	var rowsData []screen
	conf := readConfig.ReadConfig()
	//声明数据库连接字符串
	conn := SQL.ConnSQL()
	defer conn.Close()
	//编写查询语句
	stmt, err := conn.Prepare(`select 设备网络IP,显示图片 from dbo.esop表单`)
	if err != nil {
		log.Println("Sql Prepare failed:", err.Error())
		return errors.New("从数据库查询失败")
	}
	defer stmt.Close()

	//执行查询语句
	rows, err := stmt.Query()
	if err != nil {
		log.Println("Query failed:", err.Error())
		return errors.New("从数据库查询失败")
	}
	//将数据读取到实体中
	for rows.Next() {
		var row screen
		rows.Scan(&row.Ip, &row.Image)
		rowsData = append(rowsData, row)
	}
	//读取不到信息则返回空
	if rowsData == nil {
		log.Println("Can not get Data,Please Check SqlServer!!!")
		return errors.New("数据库返回消息长度为0")
	}
	//读取到信息则通过tcp传递信息
	esopPort := (*conf)["esop_port"]
	ftpAddr := (*conf)["ftpAddr"]
	var wg sync.WaitGroup
	for _, row := range rowsData {
		ip := row.Ip
		image := row.Image
		go func() {
			wg.Add(1)
			defer wg.Done()
			//传递的信息为IP + TCP地址
			ret, err := Client.SendMessage(ip+esopPort, "Pic:"+ftpAddr+image)
			if err != nil {
				log.Print(err)
				log.Print(ip + " : 发送图片出现错误！")
			} else {
				//log.Println(ip + esopPort + " " + "Pic:" + ftpAddr + image)
				log.Print(ip + " : 图片发送成功！终端返回参数为: " + ret)
			}
		}()
	}
	wg.Wait()
	return nil
}
