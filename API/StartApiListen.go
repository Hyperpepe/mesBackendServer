package API

import (
	"first/devices/EsopScreen"
	"log"
	"net/http"
)

func StartApiListen(ApiIp string) {
	//读取配置文件
	//conf := ReadConfig.ReadConfig()
	//读取配置文件中的监听端口
	//ApiIp = (*conf)["API_ListenAddr"]
	log.Println(ApiIp)
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
		switch getArgs.Get("act") {
		case "综合性能检测台":
			go func() {

			}()
		case "下发工艺卡片":
			go func() {
				err := EsopScreen.SendMessageToAll()
				//传递完成之后对MES服务器返回是否成功
				if !err {
					_, err := w.Write([]byte("failed"))
					if err != nil {
						log.Print("调用图片显示错误")
					}
				} else {
					_, _ = w.Write([]byte("ok"))
					//log.Println("2.API写入返回值成功")
				}
			}()
		case "检查esop状态":
			EsopScreen.CheckStatues()
			//sta, err := json.Marshal(status)
			//log.Println(sta)
			//if err != nil {
			//	log.Println("状态值转json失败", err)
			//	return
			//}
			//w.Header().Set("Content-Type", "application/json")
			//_,err  = w.Write(sta)
			//if err != nil {
			//	log.Print("返回esop状态值到mes主机错误")}
		default:
			log.Println("API找不到调用的指令!,接收参数: ", getArgs)
		}
	} else {
		log.Println("接收不到参数!")
	}
}
