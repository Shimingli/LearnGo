package main

import (
	"fmt"
	//log "github.com/sirupsen/logrus"
	"os"
	"github.com/sirupsen/logrus"
)



func main() {
	fmt.Println("我们期望开发的Web应用程序能够把整个程序运行过程中出现的各种事件一一记录下来，Go语言中提供了一个简易的log包，我们使用该包可以方便的实现日志记录的功能，这些日志都是基于fmt包的打印再结合panic之类的函数来进行一般的打印、抛出错误处理。Go目前标准包只是包含了简单的功能，如果我们想把我们的应用日志保存到文件，然后又能够结合日志实现很多复杂的功能（编写过Java或者C++的读者应该都使用过log4j和log4cpp之类的日志工具），可以使用第三方开发的日志系统:logrus和seelog，它们实现了很强大的日志功能，可以结合自己项目选择")

	// https://github.com/cihub/seelog

   // https://github.com/sirupsen/logrus  https://github.com/sirupsen/logrus
   // todo  logrus是用Go语言实现的一个日志系统，与标准库log完全兼容并且核心API很稳定,是Go语言目前最活跃的日志库

   // 运行这个 会报错啊 ----》 todo  麻痹还要有一个依赖啊
	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//}).Info("A walrus appears")
   // fmt.Println("start ")
	//logrus.Info("nihao ")

    //  1.cannot find package "golang.org/x/sys/windows"   需要在这个目录下创建一个文件夹 x\sys\windows
	// 设置输出，这和package的方式有所不同，它是以属性的方式赋值
	//log.Out = os.Stdout
	//log.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")

     //创建一个文本
	file, err := os.OpenFile("DeploymentAndMaintenance/20180726.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		fmt.Println("this ---")
		logone.Out = file
	} else {
		fmt.Println("there  ---")
		logone.Info("Failed to log to file, using default stderr")
	}
	// 加入 一个文件插入日志去写
	logone.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("打开文件失败ddd")
}
// 创建一个logrus示例
var logone = logrus.New()
