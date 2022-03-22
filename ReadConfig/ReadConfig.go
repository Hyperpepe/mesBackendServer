package ReadConfig

import (
	"github.com/astaxie/beego/config"
	"log"
)

func ReadConfig() *map[string]string {
	//var ret map[string]string
	conf, err := config.NewConfig("ini", "C:/Users/tss05/goProject/first/config.conf")
	if err != nil {
		log.Print("config read error!")
		log.Println(err)
	}
	return &map[string]string{"sql_server": conf.String("sql_server"),
		"sql_port":       conf.String("sql_port"),
		"sql_user":       conf.String("sql_user"),
		"sql_pass":       conf.String("sql_pass"),
		"sql_database":   conf.String("sql_database"),
		"esop_port":      conf.String("esop_port"),
		"API_ListenAddr": conf.String("API_ListenAddr")}
}
