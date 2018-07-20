package main

import "fmt"

func main() {
	fmt.Println("Web服务 --让Http协议的基础上通过XML或者是JSON来交换信息")
    //Web 服务   :REST  SOAP
    //  REST即表述性状态传递（英文：Representational State Transfer，简称REST）是Roy Fielding博士在2000年他的博士论文中提出来的一种软件架构风格。它是一种针对网络应用的设计和开发方式，可以降低开发的复杂性，提高系统的可伸缩性。

    fmt.Println("REST 请求是很直观的，因为REST是基于HTTP协议的 ，他的每一次请求都是一个HTTP请求，然后根据不同的method来处理不同的逻辑，")


	//  SOAP  简单对象访问协议是交换数据的一种协议规范，是一种轻量的、简单的、基于XML（标准通用标记语言下的一个子集）的协议，它被设计成在WEB上交换结构化的和固化的信息。
	fmt.Println("SOAP 是W3C 在跨网络信息传递和远程计算机函数调用方面的一个标准。但是SOAP非常复杂，然后完整的规范篇幅很长，而且内容任然在增加。Go语言提供了一种天生性能很不错的，开发起来非常方便的RPC机制，")

	// Go语言是21世纪的C语言，性能和简单，很多的游戏的服务都是采用的是Socket，因为HTTP协议相对而言比较耗费性能，
	// 随着HTML5的发展，webSockets 也逐渐发展成为很多页游公司接下来开发的一些手段


	//  Socket
	// WebSocket
	//REST
	// RPC

	// Socket编程，现在的网络正在向着云的方向快速进化，技术演化的基石是 Socket
	// HTML5  一个重要的 特性是 WebSocket ，通过它 ，可以主动地push下坡熊，以简化ajax轮询的模式
	// REST 编写模式，特别适合来开发网络应用的API ，目前移动应用的快速的发展，这个一个潮流
	// Go RPC 实现RPC的相关的知识

}
