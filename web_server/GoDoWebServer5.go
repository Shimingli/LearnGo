package main

import (
	"fmt"
	"net/http"
    //草拟吗 自己创建的目录 哈哈哈哈哈    还好我比较聪明  要不然 就完蛋了  麻痹
	"golang.org/x/net/websocket"
	"log"
)

func main() {
	fmt.Println("Go语言标准包里面没有提供对WebSocket的支持，但是在由官方维护的go.net子包中有对这个的支持 go get golang.org/x/net/websocket")
	//打印这个信息就，os.Exit(1)  退出程序
	//log.Fatal("shiming")  todo  草拟吗 啊   看清楚啊   后面的域名的地址 有个老子的名字啊
    http.Handle("/shiming",websocket.Handler(Echo))
	 if err:=http.ListenAndServe(":8080",nil);err!=nil{
	 	log.Fatal(err)
	 }


}

func Echo(w *websocket.Conn)  {
    var error error
	for   {
		var reply string
		if  error= websocket.Message.Receive(w,&reply);error!=nil{
			fmt.Println("不能够接受消息 error==",error)
			break
		}
		fmt.Println("能够接受到消息了--- ",reply)
		msg:="我已经收到消息 Received:"+reply
		//  连接的话 只能是   string；类型的啊
		fmt.Println("发给客户端的消息： "+msg)

		if error = websocket.Message.Send(w, msg); error != nil {
			fmt.Println("不能够发送消息 悲催哦")
			break
		}




	}


}
