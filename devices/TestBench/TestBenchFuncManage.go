package TestBench

import "first/devices/TestBench/SafetyTest"

func TestBenchFuncManage(TcpMessages string) string {
	messMap := convert(TcpMessages)
	switch (*messMap)["ItemName"] {
	case "Safety_test":
		SafetyTest.SafetyTestFunc(*messMap)
		return "ok"
	default:
		return "fail"
	}
}
