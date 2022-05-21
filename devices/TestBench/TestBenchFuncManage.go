package TestBench

import (
	"errors"
	CptbTest "first/devices/TestBench/CPTBTestBench"
	SafeTest "first/devices/TestBench/SafetyTest"
	"log"
	"strings"
)

func FuncManage(TcpMessages string) (string, error) {
	//验证字符串的头和尾
	if (strings.Count(TcpMessages, "#")) != 2 {
		log.Printf("接收到的字符串 # 验证错误！ 请检查发送格式是否正确")
		return "fail", errors.New("接收到的字符串验证符验证错误（#）")
	}
	//将字符换的#信息替换掉
	TcpMessages = strings.Replace(TcpMessages, "#", "", -1)
	//将传递的字符串转换为MAP格式
	messMap, err := Convert(TcpMessages)
	if err != nil {
		log.Print(err)
		return "fail", errors.New("转换字符串格式失败，请查看日志！")
	}
	//function调用过程
	switch (*messMap)["ItemName"] {
	case "Safety_test":
		err := SafeTest.SafetyTestFunc(*messMap)
		if err != nil {
			log.Print(err)
			return "fail", errors.New("调用方法SafetyTest失败！")
		}
		return "ok", nil
	case "CPTB_test_bench":
		err := CptbTest.CptbTestFunc(*messMap)
		if err != nil {
			log.Print(err)
			return "fail", errors.New("调用方法CPTBTestBench失败！")
		}
		return "ok", nil
	case "ID_system":
		if err != nil {
			log.Print(err)
			return "fail", errors.New("调用ID_system方法失败！")
		}
		return "ok", nil
	default:
		log.Printf("接收到的字符串无法解析！")
		return "fail", errors.New("找不到需要调用的方法！")
	}
}
