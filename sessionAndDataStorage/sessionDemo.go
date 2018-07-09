package main

import "fmt"

func main() {
    fmt.Println("Web开发中一个很重要的议题就是如何做好用户的整个浏览过程的控制，因为HTTP协议是无状态的，所以用户的每一次请求都是无状态的，我们不知道在整个Web操作过程中哪些连接与该用户有关")

    fmt.Println("Web里面经典的解决方案是cookie和session，cookie机制是一种客户端机制，把用户数据保存在客户端，而session机制是一种服务器端的机制，服务器使用一种类似于散列表的结构来保存信息，每一个网站访客都会被分配给一个唯一的标志符,即sessionID,它的存放形式无非两种:要么经过url传递,要么保存在客户端的cookies里.当然,你也可以将Session保存到数据库里,这样会更安全,但效率方面会有所下降。")


    // 1、session 机制和cookie机制的关系和区别
    // 2、Go语言如何实现session，实现一个见得session管理器
    // 3、防止session被劫持的情况，如何有效的保护session。
    // 4、里面实现的session是存储在内存中的，但是如果我们的应用进一步扩展了，要实现应用的session共享，那么我们可以把session存储在数据库中(memcache或者redis)
    // 5、如何实现它


}


