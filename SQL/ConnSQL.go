package SQL

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/config"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func ConnSQL() *sql.DB {
	configAddr := "C:/Users/tss05/goProject/first/config.conf"
	conf, err := config.NewConfig("ini", configAddr)
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
	}
	return conn
}
