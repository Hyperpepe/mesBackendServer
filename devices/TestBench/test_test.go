package TestBench

import (
	"testing"
)

func Test(t *testing.T) {
	////SaveResToSql(resInfo{})
	//info := resInfo{
	//	typ:   "C5",
	//	staff: "admin",
	//	sn:    "1001",
	//	stime: time.Now(),
	//	etime: time.Now(),
	//	file:  "/hello",
	//	res:   "合格",
	//}
	//SaveResToSql("水路气密性检测", info)
	message := "--SN:11010209004AE323000002--ItemName:Safety_test--Result:OK--Safety_ST:2020-04-01 18:49:10--Safety_ET:2020-04-01 18:49:10--ACW:ACW_V,1800;ACW_C,0,10,25--IRT:IRT_V,500;IRT_R,10,500,600--GRT:GRT_C,25;GRT_R,0,50,200--LCT:LCT_V,233;LCT_C,0,1.5,3.5"
	TestBenchFuncManage(message)
	//log.Println((*ret)["ItemName"])
}
