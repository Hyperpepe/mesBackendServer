package TestBench

import (
	"testing"
	"time"
)

func TestSaveResToSql(t *testing.T) {
	//SaveResToSql(resInfo{})
	info := resInfo{
		typ:   "C5",
		staff: "admin",
		sn:    "1001",
		stime: time.Now(),
		etime: time.Now(),
		file:  "/hello",
		res:   "合格",
	}
	SaveResToSql("水路气密性检测", info)
}
