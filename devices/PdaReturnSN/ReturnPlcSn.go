package PdaReturnSN

import (
	"errors"
	"first/readConfig"
	Client "first/tcpClient"
	"log"
)

func ReturnSN(sn string) error {
	conf := readConfig.ReadConfig()
	ip := (*conf)["ID_Computer"]
	formatSn := "#--" + sn + "#"
	log.Print("IP:" + ip)
	mess, err := Client.SendMessage(ip, formatSn)
	log.Print(err)
	if err != nil {
		return errors.New("发送信息失败，请检查日志！")
	}
	if mess != "ok!" {
		return errors.New("服务器返回结果失败，请检查日志！")
	}
	return nil
}
