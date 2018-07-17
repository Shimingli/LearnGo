package main

import "fmt"

func main() {
	//在WebSocket出现之前，为了实现即时通信，采用的技术都是“轮询”，即在特定的时间间隔内，由浏览器对服务器发出HTTP Request，服务器在收到请求后，返回最新的数据给浏览器刷新，“轮询”使得浏览器需要对服务器不断发出请求，这样会占用大量带宽。
   fmt.Println("WebSocket 是HTML5的重要特性。它实现了 基于浏览器的远程的Socket，它使的浏览器可以很服务器进行双向的通信 ")

   fmt.Println("WebSocket采用了一些特殊的报头，使得浏览器和服务器只需要做一个握手的动作，就可以在浏览器和服务器之间建立一条连接通道。且此连接会保持在活动状态，你可以使用JavaScript来向连接写入或从中接收数据，就像在使用一个常规的TCP Socket一样 ")

  // WebSocket  解决了Web 实时话，相比传统的HTTP 有些好处

  // 1、 一个Web客户端只建立一个TCP连接
  // 2、 WebSocket 服务端可以推送 push 数据到 web客户端
  // 3、有更加轻量级的头，减少数据传送量

  //WebSocket URL的起始输入是ws://或是wss://（在SSL上）。下图展示了WebSocket的通信过程，一个带有特定报头的HTTP握手被发送到了服务器端，接着在服务器端或是客户端就可以通过JavaScript来使用某种套接口（socket），这一套接口可被用来通过事件句柄异步地接收数据。
  // todo  图片  WebSocket8.2.websocket.png WebSocket原理图

// WebSocket原理 WebSocket的协议颇为简单，在第一次handshake通过以后，连接便建立成功，其后的通讯数据都是以”\x00″开头，以”\xFF”结尾。在客户端，这个是透明的，WebSocket组件会自动将原始数据“掐头去尾”。

  //todo  注意看 websocket2.png 图 WebSocket的request和response信息

  // 在请求中的"Sec-WebSocket-Key"是随机的，对于整天跟编码打交道的程序员，一眼就可以看出来：这个是一个经过base64编码后的数据。
  // 服务器端接收到这个请求之后需要把这个字符串连接上一个固定的字符串： 258EAFA5-E914-47DA-95CA-C5AB0DC85B11

    //即：f7cb4ezEAl6C3wRaU6JORA==连接上那一串固定字符串，生成一个这样的字符串：
	//f7cb4ezEAl6C3wRaU6JORA==258EAFA5-E914-47DA-95CA-C5AB0DC85B11

	//对该字符串先用 sha1安全散列算法计算出二进制的值，然后用base64对其进行编码，即可以得到握手后的字符串：
	//rE91AJhfC+6JdVcVXOGJEADEJdQ=
	//将之作为响应头Sec-WebSocket-Accept的值反馈给客户端。



}
