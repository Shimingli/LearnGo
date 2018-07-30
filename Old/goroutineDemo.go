package Old

import (
	"fmt"
	"runtime"
	"time"
)
//并发
func init() {
	fmt.Println("有人把Go比作21世纪的C语言，第一是因为Go语言设计简单，第二，21世纪最重要的就是并行程序设计，而Go从语言层面就支持了并行。")
}

func main() {
    fmt.Println("goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。")
    //goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。
	//go hello(1, 2, 3)  如果这一行输入的话，后面的 world 就不会输入了

	//go关键字很方便的就实现了并发编程
	/**
	todo   多个goroutine 运行在同一进程里面，共享内存数据，不过设计上要遵守：  不要通过共享来通信，而是通过通信来共享
	 */
    go say("world-aaaaaaaaaaa")//  开启一个新的Goroutines 执行
	go say("world-------------")//  开启一个新的Goroutines 执行
    say("hello--bbbb")//  当前的 Goroutines 执行

    /*
角标 i=== 0 hello=----********************
角标 i=== 1 hello=----********************
角标 i=== 4 world-------------
角标 i=== 2 hello=----********************
角标 i=== 3 hello=----********************
角标 i=== 4 hello=----********************

    todo   注意多种的输出的结果
     */
    say("hello=----********************")//  当前的 Goroutines 执行




    /**
    go  routine  运行在相同的地址，因此访问共享内存必须做好同步，那么go routine 是如何进行数据的通讯呢，Go提供了一个很好的通信机制channel
     Channel是Go中的一个核心类型,你可以把它看成一个管道, 可以通过 channel 发送和接受值，这些值只能是 channle的类型，定义一个channel时，也需要定义
    发送到channel的值的类型，注意必须使用make创建channel
     */
     fmt.Println("执行到这里来了")
    channelsDemo()



     //非缓存类型的channel，不过Go也允许指定channel的缓冲大小，很简单，就是channel可以存储多少元素
     bufferedChannelsDemo()


     //Range和Close   ---->Demo
     rangeAndCloseDemo()

     //多个channel 的情况，Go提供了一个关键字 select ，通过select可以监听 channel 上的数据流动

     selectDemo()

    //超时
    //有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？我们可以利用select来设置超时，通过如下的方式实现：
    caoShiDemo()



	 //runtime goroutine
	//runtime包中有几个处理goroutine的函数：
	//runtime.Goexit()
	//Goexit
	//退出当前执行的goroutine，但是defer函数还会继续调用
	 //runtime.Gosched()
	//Gosched
	//让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
	 var numcpu=runtime.NumCPU()
	 println("numcpu==返回 CPU 核数量",numcpu)
	//NumCPU
	//返回 CPU 核数量
	 var numgoroutine=runtime.NumGoroutine()
	 println("返回正在执行和排队的任务总数",numgoroutine)
	//NumGoroutine
	//返回正在执行和排队的任务总数
	//runtime.GOMAXPROCS()
	//GOMAXPROCS
	//用来设置可以并行计算的CPU核数的最大值，并返回之前的值。

}
func caoShiDemo() {
	c := make(chan int)
	o := make(chan bool)
    //o<-false
    //fmt.Println("打印出来---》",o)
	go func() {
		for {
			select {
			case v := <- c:  // 从c中接收数据，并赋值给v  如果成功的话 就打印出来
				println(v)
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true//如果把这段代码，给注释掉的话，就会报错了---》  麻痹
				fmt.Println("0====",o)
				break
			}
		}
	}()
	//最后的结果，打印的结果----
	var  flag= <- o
	fmt.Println(flag)




}
/**
select 默认是阻塞的，只有当监听的 channel 中有发送或者接受可以进行时才会运动，当多个channel 都准备好了，select是随机选择一个执行的

 */
func selectDemo() {

	//for true  {
	//	fmt.Printf("这是无限循环。\n");
	//}
	//其实就是个无线的循环和C一样
	//for   {

	//}
	fmt.Println("<---------------------------------->")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {

			fmt.Println("for循环的结果",<-c)
			//fmt.Println("shiming",<-c)
		}
		quit <- 0
	}()
	//<---------------------------------->
	//c====== 0xc0420b0000
	//quit====== 0xc0420b0060   其实下面就是两个地址  有点意思啊啊啊啊啊啊啊啊
	fmt.Println("c======",c)
	fmt.Println("quit======",quit)
	fibonacciD(c, quit)
}

func fibonacciD(c, quit chan int) {
	x,y :=1,1
	for   {
		fmt.Println("你要运行好多次 x= ",x,"y=",y)
		select {

		case c<- x:  //channel通过操作符<-来接收和发送数据

		    fmt.Println("shiming fibonacciD",c)
		  	x,y=y,x+y
		case <-quit:
			fmt.Println("quit====")
			return

		}
	}

}




/**
  x :=<-c  //c的地址是一样的
	fmt.Println("shiming 我感觉C会打印什么",c)
    y:=<-c
 */
func rangeAndCloseDemo() {
   //注释的例子，我们需要读取两次 c ，这样不是很方便，Go考虑到了这一点，所以也可以通过 range，像操作slice或者是map一样操作缓存类型的channel，
   c:=make(chan int,10)//创建可以存储 长度10的 类型为int的 channel
   go fibonacci(cap(c),c)
	v1, ok1 := <-c  /// 从c中接收数据，并赋值给v1,ok1为true 表示才有值  ，如果没有的话  ，就表示没有值，这个通道已经关闭了
	fmt.Println("ok=",ok1,"v==",v1)//ok= true v== 1 为  true 表示有数据要返回哦
	for i:=range c  {
		/**
		for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了
		 */
		fmt.Println("得到的值是",i)
	}
	v, ok := <-c
	//ok= false v== 0
	fmt.Println("ok=",ok,"v==",v)//如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。



	//记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
	/*
	panic：  恐慌
1、内建函数
2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
4、直到goroutine整个退出，并报告错误

	 */

	//另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的
}
//斐波那契 -----》数列
func fibonacci(n int, c chan int) {
	x,y :=1,1
	for i:=0;i<n ;i++  {
		c<- x //	c <- x  // send total to c  其实发送的是地址
		fmt.Println("x=",x,"y=",y)
		x,y = y,x+y//草拟吗  你看不清除这个预算符号么 ，麻痹
		fmt.Println(x,y)
	}
	close(c)
}




func bufferedChannelsDemo() {
	//在这个channel 中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间
	ch:= make(chan bool, 4)//，创建了可以存储4个元素的bool 型channel
    ch<-true
    ch<-true
    ch<-true
    ch<-true  //	c <- total  // send total to c  其实发送的是地址
    //ch<-true  // 长度只能是4 ，要不然会报错 fatal error: all goroutines are asleep - deadlock!
    fmt.Println("ch=",ch)


	//ch := make(chan type, value)
	//当 value = 0 时，channel 是无缓冲阻塞读写的，当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。


	c := make(chan int, 2)//修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	//修改为1报如下的错误:
	//fatal error: all goroutines are asleep - deadlock!
}
/**
channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具
 */
func channelsDemo() {
   //a := make(chan int)
   //
   //println("shiming aa=====",a)

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	//截取角标3 前面的数组，不包含3
	fmt.Println(a[:len(a)/2])//[7 2 8]
	go sum(a[:len(a)/2], c)
	// a[len(a)/2:]  包含角标3，[-9,4,0]
	go sum(a[len(a)/2:], c)

	// v := <-ch  // 从ch中接收数据，并赋值给v
	fmt.Println("shiming 我感觉C会打印什么",c)
	//x, y := <-c, <-c  // receive from c
    x :=<-c  //c的地址是一样的
	fmt.Println("shiming 我感觉C会打印什么",c)
    y:=<-c
	fmt.Println("shiming 我感觉C会打印什么",c)

    // 感觉 x=-5 或者是 x=17，不确定哪个是什么
	fmt.Println(x, y, x + y)
}

func sum(a []int,c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	//total= 17     total= -5
	fmt.Println("total=",total)
	//channel通过操作符<-来接收和发送数据
	c <- total  // send total to c  其实发送的是地址
	fmt.Println("发送数据到C,C是多少",c)
}

func say(s string) {
	for i:=0;i<5 ;i++  {
		//GOGHED产生处理器，允许其他GOOTEON运行。它不
		//暂停当前GOODUTE，因此执行自动恢复。
		runtime.Gosched()//表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		fmt.Println("角标 i===",i,s)
	}

	//Go 1.5以前调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，需要在我们的程序中显式调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。
	//var n=runtime.GOMAXPROCS(10)
	//println("shiming nffffffffffffff===",n)
	//var n1=runtime.GOMAXPROCS(10)
	//println("shiming neeeeeeeeeeeeeeeeeeeeee===",n1)
}
func hello(a interface{}, b interface{}, c interface{}) {

}