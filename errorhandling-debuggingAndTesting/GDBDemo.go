package main

import (
	"fmt"
	"time"
)

func main() {
	//GDB调试简介
	//GDB是FSF(自由软件基金会)发布的一个强大的类UNIX系统下的程序调试工具。使用GDB可以做如下事情：
	//
	//启动程序，可以按照开发者的自定义要求运行程序。
	//可让被调试的程序在开发者设定的调置的断点处停住。（断点可以是条件表达式）
	//当程序被停住时，可以检查此时程序中所发生的事。
	//动态的改变当前程序的执行环境。
	//目前支持调试Go程序的GDB版本必须大于7.1。
	//
	//编译Go程序的时候需要注意以下几点
	//
	//传递参数-ldflags "-s"，忽略debug的打印信息
	//传递-gcflags "-N -l" 参数，这样可以忽略Go内部做的一些优化，聚合变量和函数等优化，这样对于GDB调试来说非常困难，所以在编译的时候加入这两个参数避免这些优化。


	fmt.Println("")
	msg := "黄色小电影------》》开始入侵电脑了------------>"
	fmt.Println(msg)
	// c:=make(chan int,10)//创建可以存储 长度10的 类型为int的 channel
	// todo  类型为int的 channel
	bus := make(chan int)
	go counting(bus)
	//  todo 	for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了
	for count := range bus {
		fmt.Println("入侵病毒数量 count=", count)
	}

	//编译文件，生成可执行文件gdbfile:
	//  注意是编译的是 -l  不是-1
	//go build -gcflags "-N -l" gdbfile.go
	//	通过gdb命令启动调试：
	//
	//gdb gdbfile

}
/*
  go  routine  运行在相同的地址，因此访问共享内存必须做好同步，那么go routine 是如何进行数据的通讯呢，Go提供了一个很好的通信机制channel
     Channel是Go中的一个核心类型,你可以把它看成一个管道, 可以通过 channel 发送和接受值，这些值只能是 channle的类型，定义一个channel时，也需要定义
    发送到channel的值的类型，注意必须使用make创建channel
 */
//	c <- x  // send total to c  其实发送的是地址
func counting(c chan<- int) {
	fmt.Println("c==",c)
	for i := 0; i < 100; i++ {
		//睡两秒钟 然后把这个值才发回给 c
		time.Sleep(2 * time.Second)
		c <- i //c <- x  // send total to c  其实发送的是地址
	}
	//		for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了
	close(c)
}
func init() {
	fmt.Println("GDB 的使用")
}
