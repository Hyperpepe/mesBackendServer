package Server

import (
	"fmt"
	"log"
	"net"
)

//根据ip地址监听相关信息
func StartListen(ipAddr string) {
	fmt.Println("Starting the server ...")
	// 创建 listener
	//go func() {
	//	listener, err := net.Listen("tcp", ipAddr)
	//	if err != nil {
	//		fmt.Println("Error listening", err.Error())
	//	}
	//	// 监听并接受来自客户端的连接
	//	for {
	//		conn, err := listener.Accept()
	//		if err != nil {
	//			fmt.Println("Error accepting", err.Error())
	//		}
	//		message := doServerStuff(conn)
	//		fmt.Println(message)
	//	}
	//}()

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
		log.Println(message)
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
