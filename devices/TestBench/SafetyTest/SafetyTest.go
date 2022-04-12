package SafetyTest

import (
	"log"
	"strings"
	"time"
)

func SafetyTestFunc(message map[string]string) string {
	//conn := SQL.ConnSQL()
	Result, SN, SafetySt, SafetyEt, ACW, GRT, IRT, LCT :=
		message["Result"], message["SN"],
		message["Safety_ST"], message["Safety_ET"],
		message["ACW"], message["GRT"],
		message["IRT"], message["LCT"]

	StartTime, _ := time.ParseInLocation("2006-01-02 15:04:05", SafetySt, time.Local)
	EndTime, _ := time.ParseInLocation("2006-01-02 15:04:05", SafetyEt, time.Local)

	ACW_V := strings.Split(strings.Split(ACW, ";")[0], ",")[1]
	ACW_C_min := strings.Split(strings.Split(ACW, ";")[1], ",")[1]
	ACW_C_max := strings.Split(strings.Split(ACW, ";")[1], ",")[3]
	ACW_C_val := strings.Split(strings.Split(ACW, ";")[1], ",")[2]
	IRT_V := strings.Split(strings.Split(IRT, ";")[0], ",")[1]
	IRT_R_min := strings.Split(strings.Split(IRT, ";")[1], ",")[1]
	IRT_R_max := strings.Split(strings.Split(IRT, ";")[1], ",")[3]
	IRT_R_val := strings.Split(strings.Split(IRT, ";")[1], ",")[2]
	GRT_C := strings.Split(strings.Split(GRT, ";")[0], ",")[1]
	GRT_R_min := strings.Split(strings.Split(GRT, ";")[1], ",")[1]
	GRT_R_max := strings.Split(strings.Split(GRT, ";")[1], ",")[3]
	GRT_R_val := strings.Split(strings.Split(GRT, ";")[1], ",")[2]
	LCT_V := strings.Split(strings.Split(LCT, ";")[0], ",")[1]
	LCT_C_min := strings.Split(strings.Split(LCT, ";")[1], ",")[1]
	LCT_C_max := strings.Split(strings.Split(LCT, ";")[1], ",")[3]
	LCT_C_val := strings.Split(strings.Split(LCT, ";")[1], ",")[2]

	log.Println(Result)
	log.Println(SN)
	log.Println(StartTime)
	log.Println(EndTime)
	//log.Println(ACW)
	//log.Println(GRT)
	//log.Println(IRT)
	//log.Println(LCT)
	log.Println(ACW_V)
	log.Println(IRT_V)
	log.Println(GRT_C)
	log.Println(LCT_V)
	log.Println("--------------------")
	log.Println(ACW_C_min)
	log.Println(ACW_C_max)
	log.Println(ACW_C_val)
	log.Println("--------------------")
	log.Println(IRT_R_min)
	log.Println(IRT_R_max)
	log.Println(IRT_R_val)
	log.Println("--------------------")
	log.Println(GRT_R_min)
	log.Println(GRT_R_max)
	log.Println(GRT_R_val)
	log.Println("--------------------")
	log.Println(LCT_C_min)
	log.Println(LCT_C_max)
	log.Println(LCT_C_val)
	return "ok"
}
