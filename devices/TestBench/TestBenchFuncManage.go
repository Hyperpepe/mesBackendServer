package TestBench

import "first/devices/TestBench/SafetyTest"

func TestBenchFuncManage(TcpMessages string) string {
	//根据传递来的MAP选择需要调用的FUNCTION
	messMap := Convert(TcpMessages)
	switch (*messMap)["ItemName"] {
	case "Safety_test":
		err := SafetyTest.SafetyTestFunc(*messMap)
		if err != nil {
			return "fail"
		}
		return "ok"
	default:
		return "fail"
	}
}
