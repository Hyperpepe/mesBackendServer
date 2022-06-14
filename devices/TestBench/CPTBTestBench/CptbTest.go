package TestBench

import (
	"errors"
	"first/SQL"
	"log"
	"time"
)

func CptbTestFunc(message map[string]string) error {
	log.Println("=========================正在写入综合性能检测报告===============================")
	conn := SQL.ConnSQL()
	defer conn.Close()
	SN, Result, CptbSt, CptbEt, _ :=
		message["SN"], message["Result"],
		message["CPTB_ST"], message["CPTB_ET"],
		message["CPTB_ITEM_NM"]

	//将时间字符串序列化为标准时间格式
	StartTime, err := time.ParseInLocation("2006-01-02 15:04:05", CptbSt, time.Local)
	if err != nil {
		log.Printf("综合性能检测台转换时间格式失败！请检查TCP连接与传递信息是否正确")
	}
	EndTime, err := time.ParseInLocation("2006-01-02 15:04:05", CptbEt, time.Local)
	if err != nil {
		log.Printf("综合性能检测台转换时间格式失败！请检查TCP连接与传递信息是否正确")
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
		return errors.New("从设备ID生成情况拉取相关信息失败")
	}
	for rows.Next() {
		err := rows.Scan(&orderNumber, &productNo, &productionLine, &productCode)
		if err != nil {
			//log.Fatal(err)
			return errors.New("遍历行错误！请查看错误信息")
		}
	}

	//将接收到的结果插入数据库
	stmt, err := conn.Prepare("INSERT INTO dbo.综合性能检测台 ('订单编号','产品代码','工序代码',设备序列ID,开始时间,结束时间,测试流程文件名,测试结果) VALUES (?,?,?,?,?,?,?,?)")
	if len(orderNumber) == 0 || len(productCode) == 0 || stmt == nil {
		return errors.New("从和数据库查询的订单信息为空，请检查数据库是否已录入该订单！")
	}
	_, err = stmt.Exec(orderNumber, productCode, "P0110", SN, StartTime, EndTime, SN, Result)
	if err != nil {
		return errors.New("安规测试台执行数据库写入错误！")
	}

	log.Println("=========================综合性能检测台数据写入操作完成=============================")
	return nil
}
