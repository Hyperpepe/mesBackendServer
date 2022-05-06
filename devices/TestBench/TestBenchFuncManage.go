package TestBench

import (
	"errors"
	"first/devices/TestBench/CPTBTestBench"
	"first/devices/TestBench/SafetyTest"
	"log"
	"strings"
)

func TestBenchFuncManage(TcpMessages string) (string, error) {
	if (strings.Count(TcpMessages, "#")) != 2 {
		log.Printf("接收到的字符串 # 验证错误！ 请检查发送格式是否正确")
		return "fail", errors.New("接收到的字符串验证符验证错误（#）")
	}
	TcpMessages = strings.Replace(TcpMessages, "#", "", -1)
	//根据传递来的MAP选择需要调用的FUNCTION
	messMap, err := Convert(TcpMessages)
	if err != nil {
		log.Print(err)
		return "fail", errors.New("转换字符串格式失败，请查看日志！")
	}
	switch (*messMap)["ItemName"] {
	case "Safety_test":
		err := SafetyTest.SafetyTestFunc(*messMap)
		if err != nil {
			log.Print(err)
			return "fail", errors.New("调用方法SafetyTest失败！")
		}
		return "ok", nil
	case "CPTB_test_bench":
		err := CPTBTestBench.CptbTestFunc(*messMap)
		if err != nil {
			log.Print(err)
			return "fail", errors.New("调用方法CPTBTestBench失败！")
		}
		return "ok", nil
	default:
		log.Printf("接收到的字符串无法解析！")
		return "fail", errors.New("找不到需要调用的方法！")
	}
}
