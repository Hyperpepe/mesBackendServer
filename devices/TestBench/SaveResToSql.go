package TestBench

import (
	"first/SQL"
)

type resInfo struct {
	typ   string
	staff string
	sn    string
	stime string
	etime string
	file  string
	res   string
}

func SaveResToSql(dbo string, info resInfo) {
	conn := SQL.ConnSQL()
	defer conn.Close()

	//insert, err := conn.Exec("INSERT INTO dbo.? (Name,Age,Score) VALUES (?, ?, ?)", "tom", 18, 88.5)
	//if err != nil {
	//	fmt.Println("Insert data err=", err)
	//}
	//id, err := insert.LastInsertId()
	//fmt.Printf("成功插入, id是%v\n", id)
}
