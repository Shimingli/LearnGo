package main

import (
	log "github.com/cihub/seelog"
	"fmt"

)
func init() {
	//seelog是用Go语言实现的一个日志系统，它提供了一些简单的函数来实现复杂的日志分配、过滤和格式化。主要有如下特性：
	//
	//XML的动态配置，可以不用重新编译程序而动态的加载配置信息
	//
	//支持热更新，能够动态改变配置而不需要重启应用
	//
	//支持多输出流，能够同时把日志输出到多种流中、例如文件流、网络流等
	//
	//支持不同的日志输出
	//
	//命令行输出
	//文件输出
	//缓存输出
	//支持log rotate
	//SMTP邮件
	//上面只列举了部分特性，seelog是一个特别强大的日志处理系统
}

func main() {
	defer log.Flush()
	//1532588313466000000 [Info] Hello from Seelog!
	fmt.Println("seelog start ")
	log.Info("Hello from Seelog!")


	//addr, _ := configs.MainConfig.String("server", "addr")
	//logs.Logger.Info("Start server at:%v", addr)
	//err := http.ListenAndServe(addr, routes.NewMux())
	//logs.Logger.Critical("Server err:%v", err)

}
