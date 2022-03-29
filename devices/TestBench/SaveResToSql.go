package TestBench

import (
	"first/SQL"
	"fmt"
	"time"
)

type resInfo struct {
	typ   string
	staff string
	sn    string
	stime time.Time
	etime time.Time
	file  string
	res   string
}

func SaveResToSql(dbo string, info resInfo) {
	conn := SQL.ConnSQL()
	defer conn.Close()
	//执行相应的字符串
	insert, err := conn.Exec(
		`INSERT INTO dbo.? (设备编号,测试人员,设备序列ID,开始时间,结束时间,流程文件名,测试结果) VALUES (?, ?, ?, ?, ?, ?, ?)`, dbo, info.typ, info.staff, info.sn, info.stime, info.etime, info.file, info.res)
	if err != nil {
		fmt.Println("Insert data err=", err)
	}
	id, err := insert.LastInsertId()
	fmt.Printf("成功插入, id是%v\n", id)
}
