package Server

import (
	"first/devices/TestBench"
	"fmt"
	"log"
	"net"
)

// StartListen 根据ip地址监听相关信息
func StartListen(conf *map[string]string) {
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
		message := doServerStuff(conn)
		ret, err := TestBench.FuncManage(message)
		if err != nil {
			log.Print("调用程序错误，请检查错误信息->: s", err)
		}
		_, err = conn.Write([]byte(ret))
		if err != nil {
			log.Printf("写入返回值时的连接错误！")
		}
		conn.Close()
	}
}

func doServerStuff(conn net.Conn) (message string) {
	for {
		buf := make([]byte, 512)
		lenConn, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading", err.Error())
			return "" //终止程序
		}
		log.Printf("Received data: %v", string(buf[:lenConn]))
		return string(buf[:lenConn])
	}
}
