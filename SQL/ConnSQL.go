package SQL

import (
	"database/sql"
	"first/readConfig"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func ConnSQL() *sql.DB {
	conf := readConfig.ReadConfig()
	server, port, database, user, password := (*conf)["sql_server"],
		(*conf)["sql_port"],
		(*conf)["sql_database"],
		(*conf)["sql_user"],
		(*conf)["sql_pass"]
	//编写连接字符串
	connString := fmt.Sprintf("server=%s;port%s;database=%s;user id=%s;password=%s", server, port, database, user, password)
	//建立数据库连接：conn
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Println("Sql Open Connection failed:", err.Error())
	}
	return conn
}
