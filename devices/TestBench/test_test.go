package TestBench

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	message := "#--SN:11010209004AE323000001--ItemName:Safety_test--Result:OK--Safety_ST:2020-04-01 18:49:10--Safety_ET:2020-04-01 18:49:10--ACW:ACW_V,1800;ACW_C,0,10,25--IRT:IRT_V,500;IRT_R,10,500,600--GRT:GRT_C,25;GRT_R,0,50,200--LCT:LCT_V,233;LCT_C,0,1.5,3.5#"
	_, err := TestBenchFuncManage(message)
	if err != nil {
		log.Print(err)
	}
	//log.Println((*ret)["ItemName"])
}
