### 项目为MES系统部分后台
------
### 主要负责三个功能：
* TCP Server
* HTTP Server

### 主要功能介绍
* #### TCP Server
主要负责与设备之间的通信，接收到信息后对信息做解析，并执行针对不同的信息类型执行对应的方法，并向数据库存储相应的设备信息。

* #### HTTP Server
负责MES网页端的数据请求，充当数据库与前端之间的接口层。


详情查看函数介绍

-----
#### （*需修改配置文件）