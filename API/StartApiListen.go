package API

import (
	"first/devices/EsopScreen"
	"first/devices/PdaReturnSN"
	"io"
	"log"
	"net/http"
)

func StartApiListen(conf *map[string]string) {
	ApiIp := (*conf)["API_ListenAddr"]
	log.Println("Opening Api to " + ApiIp + "  ...")
	//开始监听端口
	go func() {
		http.HandleFunc("/", handler)
		err := http.ListenAndServe(ApiIp, nil)
		if err != nil {
			log.Println("API响应错误，错误代码：", err)
		}
	}()
}

//API监听调用方法
func handler(w http.ResponseWriter, r *http.Request) {
	getArgs := r.URL.Query()
	//参数调用请求,有参数则解参数并调用相关的方法，无参数则进入下一个分支调用相应的动作
	if len(getArgs) != 0 {
		log.Print("Receive API Args: ")
		log.Println(getArgs)
		//解析参数并调用相应的方法
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		sn := string(getArgs.Get("sn"))
		switch getArgs.Get("act") {
		case "下发工艺卡片":
			err := EsopScreen.SendMessageToAll()
			log.Print("下发工艺卡片操作------>正在下发！")
			//传递完成之后对MES服务器返回是否成功
			if err != nil {
				_, _ = io.WriteString(w, "failed!")
				log.Print("下发工艺卡片操作------>失败！")
				return
			} else {
				_, _ = io.WriteString(w, "ok!")
				log.Print("下发工艺卡片操作------>成功！")
				return
			}
		case "检查esop状态":
			err := EsopScreen.CheckStatus()
			if err != nil {
				_, _ = io.WriteString(w, "failed!")
				log.Print("检查esop状态操作------>失败！")
				return
			} else {
				_, _ = io.WriteString(w, "ok!")
				log.Print("检查esop状态操作------>成功！")
				return
			}
		case "返回SN":
			err := PdaReturnSN.ReturnSN(sn)
			if err != nil {
				_, _ = w.Write([]byte("failed!"))
				log.Print("返回SN码操作------->失败！")
				return
			} else {
				_, _ = w.Write([]byte("ok!"))
				log.Print("返回SN码操作------->成功！")
				return
			}
		default:
			log.Println("API找不到调用的指令!,接收参数: ", getArgs)
			return
		}
	} else {
		log.Println("接收不到参数")
	}
}
