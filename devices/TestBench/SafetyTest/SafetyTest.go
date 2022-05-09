package TestBench

import (
	"errors"
	"first/SQL"
	"log"
	"strings"
	"time"
)

func SafetyTestFunc(message map[string]string) error {
	log.Println("=========================正在写入安规检测报告===============================")
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
	StartTime, err := time.ParseInLocation("2006-01-02 15:04:05", SafetySt, time.Local)
	if err != nil {
		log.Printf("安规检测转换时间格式失败")
	}
	EndTime, err := time.ParseInLocation("2006-01-02 15:04:05", SafetyEt, time.Local)
	if err != nil {
		log.Printf("安规检测转换时间格式失败")
	}

	//从数据库查询订单所属的编号等信息
	var (
		orderNumber    string
		productNo      string
		productionLine string
		productCode    string
	)
	rows, err := conn.Query(`select 订单编号,产品型号,生产线别,产品代码 from dbo.设备ID生成情况 where 设备序列号 = ?`, SN)
	if err != nil {
		log.Println("Query failed:", err.Error())
		//return errors.New("Query failed错误")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&orderNumber, &productNo, &productionLine, &productCode)
		if err != nil {
			//log.Fatal(err)
			return errors.New("遍历行错误！请查看错误信息")
		}
	}

	//将接收到的结果插入数据库
	stmt, err := conn.Prepare("INSERT INTO dbo.安规检测 (订单编号, 产品代码,工序代码,设备序列ID,开始时间,结束时间,测试流程文件名,测试结果) VALUES (?,?,?,?,?,?,?,?)")
	if len(orderNumber) == 0 || len(productCode) == 0 || stmt == nil {
		return errors.New("从和数据库查询的订单信息为空，请检查数据库是否已录入该订单！")
	}
	_, err = stmt.Exec(orderNumber, productCode, "P0110", SN, StartTime, EndTime, SN, Result)
	if err != nil {
		return errors.New("安规测试台执行数据库写入错误！")
	}
	//将检测数据进行语义分割
	AcwV := strings.Split(strings.Split(ACW, ";")[0], ",")[1]
	AcwCMin := strings.Split(strings.Split(ACW, ";")[1], ",")[1]
	AcwCMax := strings.Split(strings.Split(ACW, ";")[1], ",")[3]
	AcwCVal := strings.Split(strings.Split(ACW, ";")[1], ",")[2]
	IrtV := strings.Split(strings.Split(IRT, ";")[0], ",")[1]
	IrtRMin := strings.Split(strings.Split(IRT, ";")[1], ",")[1]
	IrtRMax := strings.Split(strings.Split(IRT, ";")[1], ",")[3]
	IrtRVal := strings.Split(strings.Split(IRT, ";")[1], ",")[2]
	GrtC := strings.Split(strings.Split(GRT, ";")[0], ",")[1]
	GrtRMin := strings.Split(strings.Split(GRT, ";")[1], ",")[1]
	GrtRMax := strings.Split(strings.Split(GRT, ";")[1], ",")[3]
	GrtRVal := strings.Split(strings.Split(GRT, ";")[1], ",")[2]
	LctV := strings.Split(strings.Split(LCT, ";")[0], ",")[1]
	LctCMin := strings.Split(strings.Split(LCT, ";")[1], ",")[1]
	LctCMax := strings.Split(strings.Split(LCT, ";")[1], ",")[3]
	LctCVal := strings.Split(strings.Split(LCT, ";")[1], ",")[2]

	//将分割后的数据插入数据库
	stmt, err = conn.Prepare("INSERT INTO dbo.安规检测行表 (订单编号,工序代码,设备序列ID,测试项目名称,子项目名称,测试参数一,测试参数一标准,测试参数二,合格下限,检测数据,合格上限,测试结果) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")
	if stmt == nil {
		return errors.New("stmt Prepare 为空，请检查数据库")
	}
	_, err = stmt.Exec(orderNumber, "P0110", SN, "安规检测", "交流耐压", "电流", AcwV, "电压", AcwCMin, AcwCVal, AcwCMax, Result)

	_, err = stmt.Exec(orderNumber, "P0110", SN, "安规检测", "绝缘电阻", "电阻", IrtV, "电压", IrtRMin, IrtRVal, IrtRMax, Result)

	_, err = stmt.Exec(orderNumber, "P0110", SN, "安规检测", "接地电阻", "电阻", GrtC, "电压", GrtRMin, GrtRVal, GrtRMax, Result)

	_, err = stmt.Exec(orderNumber, "P0110", SN, "安规检测", "泄漏电流", "电流", LctV, "电压", LctCMin, LctCVal, LctCMax, Result)
	log.Println("========================安规检测报告写入操作完成==============================")
	return nil
}
