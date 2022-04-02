package EsopScreen

import (
	"database/sql"
	"first/ReadConfig"
	"first/SQL"
	Client "first/tcp_Client"
	"fmt"
	"log"
	"strings"
	"sync"
)

func CheckStatus() bool {
	conn := SQL.ConnSQL()
	conf := ReadConfig.ReadConfig()
	defer conn.Close()
	//编写查询语句
	stmt, err := conn.Prepare(`select 设备网络IP from dbo.esop表单`)
	if err != nil {
		fmt.Println("Prepare failed:", err.Error())
		return false
	}
	defer stmt.Close()

	//执行查询语句
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("Query failed:", err.Error())
		return false
	}
	//将数据读取到实体中
	esopPort := (*conf)["esop_port"]
	//创建线程数
	var wg sync.WaitGroup
	for rows.Next() {
		//tmp为每行暂存数据
		var tmp string
		rows.Scan(&tmp)
		wg.Add(1)
		go checkStatus(tmp, esopPort, "status:", conn, &wg)
		//fmt.Println("ip:", tmp, " status:")
	}
	wg.Wait()
	return true
}

//对每一个Ip进行访问并将结果返回到数据库中
func checkStatus(Ip string, port string, Status string, conn *sql.DB, wg *sync.WaitGroup) {
	//超线程完成操作
	defer wg.Done()
	//对Ip进行访问并获得访问结果
	ret := Client.SendMessage(Ip+port, Status)
	log.Println(ret)
	if strings.Contains(ret, "online") {
		_, err := conn.Exec("update dbo.esop表单 set 状态=1 where 设备网络IP='" + Ip + "'")
		if err != nil {
			log.Println(err)
		}
	} else {
		_, err := conn.Exec("update dbo.esop表单 set 状态=-1 where 设备网络IP='" + Ip + "'")
		if err != nil {
			log.Println(err)
		}
	}
}
