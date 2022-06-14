package TestBench

import (
	"errors"
	TestBench "first/devices/TestBench/AirTightness"
	CptbTest "first/devices/TestBench/CPTBTestBench"
	SafeTest "first/devices/TestBench/SafetyTest"
	TestBench2 "first/devices/TestBench/WaterTightness"
	"log"
	"net"
	"strings"
)

func FuncManage(conn net.Conn) error {
	TcpMessages := getMessages(conn)
	defer conn.Close()
	//验证字符串的头和尾
	if (strings.Count(TcpMessages, "#")) != 2 {
		log.Printf("接收到的字符串 # 验证错误！ 请检查发送格式是否正确")
		_, err := conn.Write([]byte("fail"))
		if err != nil {
			log.Printf("写入返回值时的连接错误！")
		}
		return errors.New("接收到的字符串验证符验证错误（#）")
	}
	//将字符换的#信息替换掉
	TcpMessages = strings.Replace(TcpMessages, "#", "", -1)
	//将传递的字符串转换为MAP格式
	messMap, err := Convert(TcpMessages)
	if err != nil {
		log.Print(err)
		_, err = conn.Write([]byte("fail"))
		if err != nil {
			log.Printf("写入返回值时的连接错误！")
		}
		return errors.New("转换字符串格式失败，请查看日志！")
	}
	//function调用过程
	switch (*messMap)["ItemName"] {
	case "Safety_test":
		err := SafeTest.SafetyTestFunc(*messMap)
		if err != nil {
			log.Print(err)
			_, err = conn.Write([]byte("fail"))
			if err != nil {
				log.Printf("写入返回值时的连接错误！")
			}
			return errors.New("调用方法SafetyTest失败！")
		}
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			log.Printf("写入返回值时的连接错误！")
		}
		return nil
	case "CPTB_test_bench":
		err := CptbTest.CptbTestFunc(*messMap)
		if err != nil {
			log.Print(err)
			_, err = conn.Write([]byte("fail"))
			if err != nil {
				log.Printf("写入返回值时的连接错误！")
			}
			return errors.New("调用方法SafetyTest失败！")
		}
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			log.Printf("写入返回值时的连接错误！")
		}
		return nil
	case "Air_tightness_test":
		err := TestBench.AirTightTestFunc(*messMap)
		if err != nil {
			log.Print(err)
			_, err = conn.Write([]byte("fail"))
			if err != nil {
				log.Printf("写入返回值时的连接错误！")
			}
			return errors.New("调用方法SafetyTest失败！")
		}
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			log.Printf("写入返回值时的连接错误！")
		}
		return nil
	case "Water_tightness_test":
		err := TestBench2.WaterTightnessTestFunc(*messMap)
		if err != nil {
			log.Print(err)
			_, err = conn.Write([]byte("fail"))
			if err != nil {
				log.Printf("写入返回值时的连接错误！")
			}
			return errors.New("调用方法SafetyTest失败！")
		}
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			log.Printf("写入返回值时的连接错误！")
		}
		return nil
	default:
		log.Printf("接收到的字符串无法解析！")
		return errors.New("找不到需要调用的方法！")
	}
}

func getMessages(conn net.Conn) (message string) {
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
