package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"time"
	"errors"
)

func init() {
	//错误处理是每个编程语言都要考虑的一个重要话题。在Go语言的错误处理中，错误是软件包API和应用程序用户界面的一个重要组成部分。
	fmt.Println("错误和异常")

	//在程序中总有一部分函数总是要求必须能够成功的运行。比如strconv.Itoa将整数转换为字符串，从数组或切片中读写元素，从map读取已经存在的元素等。这类操作在运行时几乎不会失败，除非程序中有BUG，或遇到灾难性的、不可预料的情况，比如运行时的内存溢出。如果真的遇到真正异常情况，我们只要简单终止程序就可以了

	//在C语言中，默认采用一个整数类型的errno来表达错误，这样就可以根据需要定义多种错误类型。在Go语言中，syscall.Errno就是对应C语言中errno类型的错误。在syscall包中的接口，如果有返回错误的话，底层也是syscall.Errno错误类型。

	//在Go语言中，错误被认为是一种可以预期的结果；而异常则是一种非预期的结果，发生异常可能表示程序中存在BUG或发生了其它不可控的问题。Go语言推荐使用recover函数将内部异常转为错误处理，这使得用户可以真正的关心业务相关的错误处理。

	i := 0
	//  i=i/0
	fmt.Println(i)

}
func main() {

	copyDemo()
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err1 := errors.New(x)
				fmt.Println("err1=",err1)
			case error:
				err2 := x
				fmt.Println("err2=",err2)
			default:
				//err3= Unknown panic: 102
				err3 := fmt.Errorf("Unknown panic: %v", r)
				fmt.Println("err3=",err3)

				fmt.Println("后面已经发生了什么异常了--然后我还是可以执行的哦  ***********************")
			}
		}
	}()
	// 下面的 panic 的运用，当发生了
	//panic(errors.New("我是new出来的错误"))
	//panic("TODO")
	fmt.Println("后面发生了异常了")
	panic(102)

	//基于这个代码模板，我们甚至可以模拟出不同类型的异常。通过为定义不同类型的保护接口，我们就可以区分异常的类型了：
	//
	//defer func() {
	//	if r := recover(); r != nil {
	//		switch x := r.(type) {
	//		case runtime.Error:
	//			// 这是运行时错误类型异常
	//		case error:
	//			// 普通错误类型异常
	//		default:
	//			// 其他类型异常
	//		}
	//	}
	//}()



}
func copyDemo() {
	//记录开始时间
	start := time.Now().Nanosecond()
	copyFile("New20180730/FirstGoBase/我是源文件---Go语言实现RPC.md", "New20180730/FirstGoBase/我是copy的文件(Go语言实现RPC).md")
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	//计算过程
	sum := 0
	for i := 0; i < 5; i++ {
		for i := 0; i <= 10000000; i++ {
			sum += i
		}
	}

	//记录结束时间 todot Hours将时间段表示为int64类型的纳秒数，等价于int64(d)。
	end := time.Now().Nanosecond()
	//输出执行时间，单位为毫秒。
	fmt.Println("开始的纳秒值：start=", start)
	fmt.Println("结束的纳秒值：end=", end)
	fmt.Println("copy文件的纳秒值的差：end-start=", (end - start))
	fmt.Println("copy文件的秒值的差：end-start=", (end-start)/1000000)

	//但是对于那些提供类似Web服务的框架而言；它们经常需要接入第三方的中间件。因为第三方的中间件是否存在BUG是否会抛出异常，Web框架本身是不能确定的。为了提高系统的稳定性，Web框架一般会通过recover来防御性地捕获所有处理流程中可能产生的异常，然后将异常转为普通的错误返回
	//让我们以JSON解析器为例，说明recover的使用场景。考虑到JSON解析器的复杂性，即使某个语言解析器目前工作正常，也无法肯定它没有漏洞。因此，当某个异常出现时，我们不会选择让解析器崩溃，而是会将panic异常当作普通的解析错误，并附加额外信息提醒用户报告此错误。

	defer func() {
		if p := recover(); p != nil {
			//Go语言库的实现习惯: 即使在包内部使用了panic，但是在导出函数时会被转化为明确的错误值
			//err1 := fmt.Errorf("JSON: internal error: %v", p)
		}
	}()
	// ...parser...
}

// 函数需要打开两个文件，然后将其中一个文件的内容复制到另一个文件：
// todo  代码虽然能够工作，但是隐藏一个bug。如果第一个os.Open调用成功，但是第二个os.Create调用失败，那么会在没有释放src文件资源的情况下返回
func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(dstName)
	if err != nil {
		fmt.Println("打开文件出错", err)
		return
	}
	dst, err := os.Create(srcName)
	if err != nil {
		fmt.Println("创建文件出错了", err)
		return
	}
	//我们可以通过defer语句来确保每个被正常打开的文件都能被正常关闭
	written, err = io.Copy(dst, src)
	fmt.Println("written=", written)
	dst.Close()
	src.Close()
	return written, err
}

/*
可以通过defer语句来确保每个被正常打开的文件都能被正常关闭
 */
func copyFileGood(dstName, drcName string) (written int64, err error) {
	dst, err := os.Open(dstName)
	if err != nil {
		return
	}
	//defer语句可以让我们在打开文件时马上思考如何关闭文件。不管函数如何返回，文件关闭语句始终会被执行。同时defer语句可以保证，即使io.Copy发生了异常，文件依然可以安全地关闭
	defer dst.Close()
	src, err := os.Create(drcName)

	if err != nil {
		return
	}
	defer src.Close()
	return io.Copy(dst, src)
}

