package main

import "fmt"

func init() {
	fmt.Println("很多不安全的问题都是轻信了第三方提供的数据赵成的，用户输入的数据，没有验证之前都应该是不安全的数据，直接把这些数据输出到客户端，就可能造成跨脚本攻击（XSS）的问题，如果把不安全的数据写入到数据库中，就可能赵成SQL注入的问题")

}

func main() {

	// 1  需要对第三方提供的数据，包括用户提供的数据，验证数据的合法性
	// 2  CSRF攻击 会导致受骗者发送攻击者指定的请求从而造成一些破坏
	// 3  能够增强Web应用程序的强大手段就是加密，如何存储密码，如何让密码的传输更加的安全
	// 4  双向加密的方式



}
