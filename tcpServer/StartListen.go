package Server

import (
	"first/devices/TestBench"
	"fmt"
	"log"
	"net"
)

// StartListen 根据ip地址监听相关信息
func StartListen(conf *map[string]string) {
	//从配置文件读取配置信息
	ipAddr := (*conf)["TCP_ListenAddr"]
	log.Println("Start listening Tcp/Ip from " + ipAddr + "  ...")
	listener, err := net.Listen("tcp", ipAddr)
	if err != nil {
		fmt.Println("Error listening", err.Error())
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting", err.Error())
		}
		go func() {
			err := TestBench.FuncManage(conn)
			if err != nil {
				log.Print("调用程序错误，请检查错误信息->: s", err)
			}
			if err != nil {
				log.Printf("写入返回值时出现连接错误，请检查程序日志。")
			}
		}()
	}
}
