package SafetyTest

import (
	"first/SQL"
	"fmt"
	"log"
	"strings"
	"time"
)

func SafetyTestFunc(message map[string]string) error {
	//创建连接
	conn := SQL.ConnSQL()
	defer conn.Close()
	//接收参数并序列化为Map
	Result, SN, SafetySt, SafetyEt, ACW, GRT, IRT, LCT :=
		message["Result"], message["SN"],
		message["Safety_ST"], message["Safety_ET"],
		message["ACW"], message["GRT"],
		message["IRT"], message["LCT"]
	//将时间字符串序列化为标准时间格式
	StartTime, _ := time.ParseInLocation("2006-01-02 15:04:05", SafetySt, time.Local)
	EndTime, _ := time.ParseInLocation("2006-01-02 15:04:05", SafetyEt, time.Local)

	//从数据库查询相关信息
	var (
		orderNumber    string
		productNo      string
		productionLine string
		productCode    string
	)
	rows, err := conn.Query(`select 订单编号,产品型号,生产线别,产品代码 from dbo.设备ID生成情况 where 设备序列号 = ?`, SN)
	if err != nil {
		fmt.Println("Query failed:", err.Error())
		return fmt.Errorf("Query failed错误")
	}
	for rows.Next() {
		err := rows.Scan(&orderNumber, &productNo, &productionLine, &productCode)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	//将接收到的结果插入数据库
	stmt, err := conn.Prepare("INSERT INTO dbo.安规检测 (订单编号, 产品代码,工序代码,设备序列ID,开始时间,结束时间,测试流程文件名,测试结果) VALUES (?,?,?,?,?,?,?,?)")
	res, err := stmt.Exec(orderNumber, productCode, "P0110", SN, StartTime, EndTime, SN, Result)
	log.Println(&res)
	log.Println("=========================操作完成===============================")
	//将检测细节进行语义分割
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

	stmt, err = conn.Prepare("INSERT INTO dbo.安规检测行表 (订单编号,工序代码,设备序列ID,测试项目名称,子项目名称,测试参数一,测试参数一标准,测试参数二,合格下限,检测数据,合格上限,测试结果) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")
	res, err = stmt.Exec(orderNumber, "P0110", SN, "安规检测", "交流耐压", "电流", ACW_V, "电压", ACW_C_min, ACW_C_val, ACW_C_max, Result)

	res, err = stmt.Exec(orderNumber, "P0110", SN, "安规检测", "绝缘电阻", "电阻", IRT_V, "电压", IRT_R_min, IRT_R_val, IRT_R_max, Result)

	res, err = stmt.Exec(orderNumber, "P0110", SN, "安规检测", "接地电阻", "电阻", GRT_C, "电压", GRT_R_min, GRT_R_val, GRT_R_max, Result)

	res, err = stmt.Exec(orderNumber, "P0110", SN, "安规检测", "泄漏电流", "电流", LCT_V, "电压", LCT_C_min, LCT_C_val, LCT_C_max, Result)
	//log.Println(Result)
	//log.Println(SN)
	//log.Println(StartTime)
	//log.Println(EndTime)
	//log.Println(ACW)
	//log.Println(GRT)
	//log.Println(IRT)
	//log.Println(LCT)
	//log.Println(ACW_V)
	//log.Println(IRT_V)
	//log.Println(GRT_C)
	//log.Println(LCT_V)
	//log.Println("--------------------")
	//log.Println(ACW_C_min)
	//log.Println(ACW_C_max)
	//log.Println(ACW_C_val)
	//log.Println("--------------------")
	//log.Println(IRT_R_min)
	//log.Println(IRT_R_max)
	//log.Println(IRT_R_val)
	//log.Println("--------------------")
	//log.Println(GRT_R_min)
	//log.Println(GRT_R_max)
	//log.Println(GRT_R_val)
	//log.Println("--------------------")
	//log.Println(LCT_C_min)
	//log.Println(LCT_C_max)
	//log.Println(LCT_C_val)
	return nil
}
