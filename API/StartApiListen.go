package API

import (
	"first/devices/EsopScreen"
	"log"
	"net/http"
	"sync"
)

func StartApiListen(ApiIp string) {
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
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		getArgs := r.URL.Query()
		//参数调用请求,有参数则解参数并调用相关的方法，无参数则进入下一个分支调用相应的动作
		if len(getArgs) != 0 {
			log.Print("Receive API Args: ")
			log.Println(getArgs)
			//解析参数并调用相应的方法
			switch getArgs.Get("act") {
			case "下发工艺卡片":
				err := EsopScreen.SendMessageToAll()
				//传递完成之后对MES服务器返回是否成功
				if !err {
					_, _ = w.Write([]byte("failed"))
				} else {
					_, _ = w.Write([]byte("ok"))
				}
			case "检查esop状态":
				err := EsopScreen.CheckStatus()
				if !err {
					_, _ = w.Write([]byte("failed"))
				} else {
					_, _ = w.Write([]byte("ok"))
				}
			default:
				log.Println("API找不到调用的指令!,接收参数: ", getArgs)
			}
		} else {
			log.Println("接收不到参数!")
		}
	}()
	wg.Wait()
}
