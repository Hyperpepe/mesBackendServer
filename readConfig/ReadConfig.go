package readConfig

import (
	"github.com/astaxie/beego/config"
	"log"
)

// ReadConfig 用于将config中配置信息解析成MAP格式
func ReadConfig() *map[string]string {
	//conf, err := config.NewConfig("ini", "C:\\Users\\tss05\\goProject\\first\\config.conf")
	conf, err := config.NewConfig("ini", "config.conf")
	if err != nil {
		log.Print("config read error!")
		log.Println(err)
	}
	return &map[string]string{
		"sql_server":     conf.String("sql_server"),
		"sql_port":       conf.String("sql_port"),
		"sql_user":       conf.String("sql_user"),
		"sql_pass":       conf.String("sql_pass"),
		"sql_database":   conf.String("sql_database"),
		"esop_port":      conf.String("esop_port"),
		"API_ListenAddr": conf.String("API_ListenAddr"),
		"TCP_ListenAddr": conf.String("TCP_ListenAddr"),
		"ID_Computer":    conf.String("ID_Computer"),
		"ftpAddr":        conf.String("ftpAddr")}
}
