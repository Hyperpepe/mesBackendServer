package Client

import (
	"errors"
	"log"
	"net"
	"time"
)

// SendMessage 向目标IP地址传递信息并接收返回值
func SendMessage(ipAddr, Message string) (string, error) {
	//设置超时时间为：2s
	connTimeout := 2 * time.Second
	//打开连接:
	conn, err := net.DialTimeout("tcp", ipAddr, connTimeout)
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		log.Println("Error dialing", ipAddr, err.Error())
		return "", errors.New("Error dialing:" + ipAddr) // 终止程序
	}
	defer conn.Close()
	_, err = conn.Write([]byte(Message))
	if err != nil {
		log.Print("Error Send", err.Error())
		return "", errors.New("error Send") //发送失败
	}
	//设置接收寄存器
	buf := [512]byte{}
	//设置读取信息超时时间
	err = conn.SetReadDeadline(time.Now().Add(connTimeout))
	n, err := conn.Read(buf[:])
	if err != nil {
		return "", errors.New("获取返回值超时，请检查相关联的设备:" + ipAddr)
	}
	//发送图片正常时返回值为ok!
	return string(buf[:n]), nil
}
