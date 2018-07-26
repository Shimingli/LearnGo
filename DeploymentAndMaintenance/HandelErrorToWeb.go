package main

import (
	"net/http"
	"fmt"
	"html/template"
	"strconv"
)
//  todo  而panic和recover是针对自己开发package里面实现的逻辑，针对一些特殊情况来设计。
func init() {
    //   todo   网站错误处理


	//我们的Web应用一旦上线之后，那么各种错误出现的概率都有，Web应用日常运行中可能出现多种错误，具体如下所示：

	// 1、数据库错误：指与访问数据库服务器或数据相关的错误。例如，以下可能出现的一些数据库错误。
	//连接错误：这一类错误可能是数据库服务器网络断开、用户名密码不正确、或者数据库不存在。
	//查询错误：使用的SQL非法导致错误，这样子SQL错误如果程序经过严格的测试应该可以避免。
	//数据错误：数据库中的约束冲突，例如一个唯一字段中插入一条重复主键的值就会报错，但是如果你的应用程序在上线之前经过了严格的测试也是可以避免这类问题。

	// 2、应用运行时错误：这类错误范围很广，涵盖了代码中出现的几乎所有错误。可能的应用错误的情况如下：
	//文件系统和权限：应用读取不存在的文件，或者读取没有权限的文件、或者写入一个不允许写入的文件，这些都会导致一个错误。应用读取的文件如果格式不正确也会报错，例如配置文件应该是ini的配置格式，而设置成了json格式就会报错。
	//第三方应用：如果我们的应用程序耦合了其他第三方接口程序，例如应用程序发表文章之后自动调用接发微博的接口，所以这个接口必须正常运行才能完成我们发表一篇文章的功能。

	// 3、HTTP错误：这些错误是根据用户的请求出现的错误，最常见的就是404错误。虽然可能会出现很多不同的错误，但其中比较常见的错误还有401未授权错误(需要认证才能访问的资源)、403禁止错误(不允许用户访问的资源)和503错误(程序内部出错)。

	// 4、操作系统出错：这类错误都是由于应用程序上的操作系统出现错误引起的，主要有操作系统的资源被分配完了，导致死机，还有操作系统的磁盘满了，导致无法写入，这样就会引起很多错误。

	// 5、网络出错：指两方面的错误，一方面是用户请求应用程序的时候出现网络断开，这样就导致连接中断，这种错误不会造成应用程序的崩溃，但是会影响用户访问的效果；另一方面是应用程序读取其他网络上的数据，其他网络断开会导致读取失败，这种需要对应用程序做有效的测试，能够避免这类问题出现的情况下程序崩溃。

}
func main() {
	//   todo  错误处理的目标
	//在实现错误处理之前，我们必须明确错误处理想要达到的目标是什么，错误处理系统应该完成以下工作：

	//1、通知访问用户出现错误了：不论出现的是一个系统错误还是用户错误，用户都应当知道Web应用出了问题，用户的这次请求无法正确的完成了。例如，对于用户的错误请求，我们显示一个统一的错误页面(404.html)。出现系统错误时，我们通过自定义的错误页面显示系统暂时不可用之类的错误页面(error.html)。

	//2、记录错误：系统出现错误，一般就是我们调用函数的时候返回err不为nil的情况，可以使用前面小节介绍的日志系统记录到日志文件。如果是一些致命错误，则通过邮件通知系统管理员。一般404之类的错误不需要发送邮件，只需要记录到日志系统。

	//3、回滚当前的请求操作：如果一个用户请求过程中出现了一个服务器错误，那么已完成的操作需要回滚。下面来看一个例子：一个系统将用户递交的表单保存到数据库，并将这个数据递交到一个第三方服务器，但是第三方服务器挂了，这就导致一个错误，那么先前存储到数据库的表单数据应该删除(应告知无效)，而且应该通知用户系统出现错误了。

	//4、保证现有程序可运行可服务：我们知道没有人能保证程序一定能够一直正常的运行着，万一哪一天程序崩溃了，那么我们就需要记录错误，然后立刻让程序重新运行起来，让程序继续提供服务，然后再通知系统管理员，通过日志等找出问题。


	// 如何处理错误 ?
	// 通知用户出现错误
	//demo404()




	//如何处理异常
	//我们知道在很多其他语言中有try...catch关键词，用来捕获异常情况，但是其实很多错误都是可以预期发生的，而不需要异常处理，应该当做错误来处理，这也是为什么Go语言采用了函数返回错误的设计，这些函数不会panic，例如如果一个文件找不到，os.Open返回一个错误，它不会panic；如果你向一个中断的网络连接写数据，net.Conn系列类型的Write函数返回一个错误，它们不会panic。这些状态在这样的程序里都是可以预期的。你知道这些操作可能会失败，因为设计者已经用返回错误清楚地表明了这一点。这就是上面所讲的可以预期发生的错误。
	//
	//但是还有一种情况，有一些操作几乎不可能失败，而且在一些特定的情况下也没有办法返回错误，也无法继续执行，这样情况就应该panic。举个例子：如果一个程序计算x[j]，但是j越界了，这部分代码就会导致panic，像这样的一个不可预期严重错误就会引起panic，在默认情况下它会杀掉进程，它允许一个正在运行这部分代码的goroutine从发生错误的panic中恢复运行，发生panic之后，这部分代码后面的函数和代码都不会继续执行，这是Go特意这样设计的，因为要区别于错误和异常，panic其实就是异常处理。如下代码，我们期望通过uid来获取User中的username信息，但是如果uid越界了就会抛出异常，这个时候如果我们没有recover机制，进程就会被杀死，从而导致程序不可服务。因此为了程序的健壮性，在一些地方需要建立recover机制。
	//GetUser(12121)


	//是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。recover仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
	b:=panicDemo(initDemo(1))

	fmt.Println("b==",b)
}
func panicDemo(f func())(b  bool) {
	fmt.Println("f===","开始执行了哦")
  defer func() {
	  if x:=recover();x!=nil {
	  	  fmt.Println("x===",x)
		  b=true
	  }
  }()
  f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
  return  b
}
func initDemo(i int) func() {
	fmt.Println("这个方法执行了啊 initDemo")
	str1 := strconv.Itoa(i)
	srt:="initDemo 产生了  panic  i="+str1
	fmt.Println(srt)
	panic(srt)

}

func GetUser(uid int) (username string) {
	defer func() {
		if x := recover(); x != nil {
			username = ""
		}
	}()
	//username = User[uid]
	return
}
func demo404() {
	http.HandleFunc("/demoError",demoError)
	http.HandleFunc("/demoError404",demoError404)
	http.HandleFunc("/SystemError",SystemError)
	http.ListenAndServe(":8080",nil)
}
func demoError404(w http.ResponseWriter, r *http.Request) {
	notFound404(w,r)

}
func demoError(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求的方式",r.Method)
	fmt.Println("请求的结尾是：",r.URL.Path)
	if r.URL.Path == "/demoError" {
		fmt.Fprintf(w,"ni hao  shiming")//输入到客户端的内容
		return
	}
}
func notFound404(writer http.ResponseWriter, request *http.Request) {
	//log.Error("页面找不到")   //记录错误日志
	t, _:= template.ParseFiles("DeploymentAndMaintenance/404.html")  //解析模板文件
	ErrorInfo := "文件找不到" //获取当前用户信息
	t.Execute(writer, ErrorInfo)  //执行模板的merger操作
}

func SystemError(w http.ResponseWriter, r *http.Request) {
	//log.Critical("系统错误")   //系统错误触发了Critical，那么不仅会记录日志还会发送邮件
	t, _ := template.ParseFiles("DeploymentAndMaintenance/error.html")  //解析模板文件
	ErrorInfo := "系统暂时不可用" //获取当前用户信息
	t.Execute(w, ErrorInfo)  //执行模板的merger操作
}
