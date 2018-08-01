package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	fmt.Println("面向并发的内存模型")
}

func main() {
	fmt.Println("")

	//在早期，CPU都是以单核的形式顺序执行机器指令。Go语言的祖先C语言正是这种顺序编程语言的代表。顺序编程语言中的顺序是指:所有的指令都是以串行的方式执行，在相同的时刻有且仅有一个CPU在顺序执行程序的指令。

	//随着处理器技术的发展，单核时代以提升处理器频率来提高运行效率的方式遇到了瓶颈，目前各种主流的CPU频率基本被锁定在了3GHZ附近。单核CPU的发展的停滞，给多核CPU的发展带来了机遇。相应地，编程语言也开始逐步向并行化的方向发展。Go语言正是在多核和网络化的时代背景下诞生的原生支持并发的编程语言。

	//常见的并行编程有多种模型，主要有多线程、消息传递等。从理论上来看，多线程和基于消息的并发编程是等价的。由于多线程并发模型可以自然对应到多核的处理器，主流的操作系统因此也都提供了系统级的多线程支持，同时从概念上讲多线程似乎也更直观，因此多线程编程模型逐步被吸纳到主流的编程语言特性或语言扩展库中。而主流编程语言对基于消息的并发编程模型支持则相比较少，Erlang语言是支持基于消息传递并发编程模型的代表者，它的并发体之间不共享内存


	//Go语言是基于消息并发模型的集大成者，它将基于CSP模型的并发编程内置到了语言中，通过一个go关键字就可以轻易地启动一个Goroutine，与Erlang不同的是Go语言的Goroutine之间是共享内存的


	goroutine()

    //原子操作
    atomicOperation()

	//顺序一致性内存模型
	memoryModel()
     //在main.main函数执行之前所有代码都运行在同一个goroutine中，也是运行在程序的主系统线程中。如果某个init函数内部用go关键字启动了新的goroutine的话，新的goroutine只有在进入main.main函数之后才可能被执行到

    //Goroutine的创建 注意也是不能够输出的哦
     creatGoroutine()

	//基于Channel的通信
	channel()

   /*
   主线程休眠了1秒钟，因此这个程序大概率是可以正常输出结果的。因此，很多人会觉得这个程序已经没有问题了。但是这个程序是不稳健的，依然有失败的可能性。我们先假设程序是可以稳定输出结果的。因为Go线程的启动是非阻塞的，main线程显式休眠了1秒钟退出导致程序结束，我们可以近似地认为程序总共执行了1秒多时间。现在假设println函数内部实现休眠的时间大于main线程休眠的时间的话，就会导致矛盾：后台线程既然先于main线程完成打印，那么执行时间肯定是小于main线程执行时间的。当然这是不可能的。
严谨的并发程序的正确性不应该是依赖于CPU的执行速度和休眠时间等不靠谱的因素的。严谨的并发也应该是可以静态推导出结果的：根据线程内顺序一致性，结合Channel或sync同步事件的可排序性来推导，最终完成各个线程各段代码的偏序关系排序。如果两个事件无法根据此规则来排序，那么它们就是并发的，也就是执行先后顺序不可靠
    */
	go println("hello, world-------- 我是睡了一秒的 ")
	time.Sleep(time.Second*1000)
}
/*
Channel通信是在Goroutine之间进行同步的主要方法。在无缓存的Channel上的每一次发送操作都有与其对应的接收操作相配对，发送和接收操作通常发生在不同的Goroutine上（在同一个Goroutine上执行2个操作很容易导致死锁）。无缓存的Channel上的发送操作总在对应的接收操作完成前发生.
 */
 var doneChannel =make(chan bool)
 var msg string
 func aGoroutine(){
 	msg = "nihao shiming gege"
 	//   todo   用close(c)关闭管道代替done <- false依然能保证该程序产生相同的行为
 	//doneChannel <- true
 	close(doneChannel)
 }

func channel() {
    go aGoroutine()
    <-doneChannel
    fmt.Println(msg)


	//对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是Channel的缓存大小。 如果将C设置为0自然就对应无缓存的Channel，也即使第K个接收完成在第K个发送完成之前。因为无缓存的Channel只能同步发1个，也就简化为前面无缓存Channel的规则：对于从无缓冲信道进行的接收，发生在对该信道进行的发送完成之前

	//我们可以根据控制Channel的缓存大小来控制并发执行的Goroutine的最大数目
	var limit = make(chan int, 3)
     fmt.Println(limit)
		//for _, w := range work {
		//	go func() {
		//		limit <- 1
		//		w()
		//		<-limit
		//	}()
		//}
		//select{}


}




var a string

func f1()  {
	fmt.Println(a)
}
func creatGoroutine() {
	//Goroutine的创建
	a="ni hao ya "
	go f1()

}

var str string
var done bool

func setup()  {
	str ="shi ming"
	done=true
}
func setup1()  {
	str ="shi ming 1"
	done=true
}
func memoryModel() {
	//创建了setup线程，用于对字符串str的初始化工作，初始化完成之后设置done标志为true。main函数所在的主线程中，通过for !done {}检测done变为true时，认为字符串初始化工作完成，然后进行字符串的打印工作
    go  setup()
    //但是Go语言并不保证在main函数中观测到的对done的写入操作发生在对字符串a的写入的操作之后，因此程序很可能打印一个空字符串。更糟糕的是，因为两个线程之间没有同步事件，setup线程对done的写入操作甚至无法被main线程看到，main函数有可能陷入死循环中
	for !done{}
	fmt.Println("str =",str)
	go setup1()
	fmt.Println("str end=",str)
    //下面的输出的语句 也可能不执行，也可能执行  如果一个并发程序无法确定事件的偏序关系，那么程序的运行结果往往会有不确定的结果
	//  todo  根据Go语言规范，main函数退出时程序结束，不会等待任何后台线程。因为Goroutine的执行和main函数的返回事件是并发的，谁都有可能先发生，所以什么时候打印，能否打印都是未知的
    go fmt.Println("你好, 世界")

	//用前面的原子操作并不能解决问题，因为我们无法确定两个原子操作之间的顺序。解决问题的办法就是通过同步原语来给两个事件明确排序
    //Go提供了一个很好的通信机制channel
	//  Channel是Go中的一个核心类型,你可以把它看成一个管道, 可以通过 channel 发送和接受值，这些值只能是 channle的类型，定义一个channel时，也需要定义
	//    发送到channel的值的类型，注意必须使用make创建channel
	done1 := make(chan int)
	fmt.Println(done1)
    go func() {
    	fmt.Println("我肯定打印的出来")
		done1<-1
	}()
	<-done1//其实发送的是地址
	fmt.Println(done1)
   //  当<-done执行时，必然要求done <- 1也已经执行。根据同一个Gorouine依然满足顺序一致性规则，我们可以判断当done <- 1执行时，println("你好, 世界")语句必然已经执行完成了。因此，现在的程序确保可以正常打印结果


   // 通过 sync.Mutex 互斥量也是可以实现同步的：

	var  mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("我也是能够输出的哦")
		mu.Unlock()
	}()
	mu.Lock()
	//可以确定后台线程的mu.Unlock()必然在println("你好, 世界")完成后发生（同一个线程满足顺序一致性），main函数的第二个mu.Lock()必然在后台线程的mu.Unlock()之后发生（sync.Mutex保证），此时后台线程的打印工作已经顺利完成了。


}





/*
所谓的原子操作就是并发编程中“最小的且不可并行化”的操作。通常，有多个并发体对一个共享资源的操作是原子操作的话，同一时刻最多只能有一个并发体对该资源进行操作。从线程角度看，在当前线程修改共享资源期间，其它的线程是不能访问该资源的。原子操作对于多线程并发编程模型来说，不会发生有别于单线程的意外情况，共享资源的完整性可以得到保证。

一般情况下，原子操作都是通过“互斥”访问来保证访问的，通常由特殊的CPU指令提供保护。当然，如果仅仅是想模拟下粗粒度的原子操作，我们可以借助于sync.Mutex来实现
 */
func atomicOperation() {
	fmt.Println("atomic")

	var wg  sync.WaitGroup
	//Add方法向内部计数加上delta，delta可以是负数；如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，如果计数器小于0，方法panic。注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。一般来说本方法应在创建新的线程或者其他应等待的事件之前调用
	wg.Add(2)
   go  worker(&wg)
   go  worker(&wg)
	//主线程里可以调用Wait方法阻塞至所有线程结束
	wg.Wait()
	fmt.Println("total.value=",total.value)

	// todo 用互斥锁来保护一个数值型的共享资源，麻烦且效率低下。标准库的sync/atomic包对原子操作提供了丰富的支持。我们可以重新实现上面的例子
	var wgTwo sync.WaitGroup
	wgTwo.Add(2)
	go workerTwo(&wgTwo)
	go workerTwo(&wgTwo)
	wgTwo.Wait()
	fmt.Println("valueTwo 两个go协程=",totalTwo)


    //  todo  原子操作配合互斥锁可以实现非常高效的单件模式。互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。
     Instance()
     Instance()
     Instance()

	//我们可以将通用的代码提取出来，就成了标准库中sync.Once的实现
	//基于标准库实现的单利的模式
	InstanceTwo()
	fmt.Println("InstanceTwo()==",instanceTwo)
	fmt.Println("0InstanceTwo()==",instanceTwo)



	// sync/atomic包对基本的数值类型及复杂对象的读写都提供了原子操作的支持。atomic.Value原子对象提供了Load和Store两个原子方法，分别用于加载和保存数据，返回值和参数都是interface{}类型，因此可以用于任意的自定义复杂类型

	//var config atomic.Value // 保存当前配置信息
	//
	//// 初始化配置信息
	//config.Store(loadConfig())
	//
	//// 启动一个后台线程, 加载更新后的配置信息
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		config.Store(loadConfig())
	//	}
	//}()
	//
	//// 用于处理请求的工作者线程始终采用最新的配置信息
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for r := range requests() {
	//			c := config.Load()
	//			// ...
	//		}
	//	}()
	//}

	//这是一个简化的生产者、消费者模型：后台线程生成最新的配置信息；前台多个工作者线程获取最新的配置信息。所有线程共享配置信息资源。
}

func InstanceTwo() *singletonTwo  {
	// 如果不定义的在变量中 会报错哦
	once.Do(func() {
		instanceTwo=&singletonTwo{}
	})
	return instanceTwo
}
type singletonTwo struct {

}
type singleton struct {

}
var (
	instance *singleton

	instanceTwo *singletonTwo
	once sync.Once

	//在性能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能
	initialized int32
	mu sync.Mutex
)
//原子操作配合互斥锁可以实现非常高效的单件模式。
func Instance() *singleton  {
	//LoadInt32原子性的获取*addr的值。
	if atomic.LoadInt32(&initialized)==1 {
		fmt.Println("单利已经生成了instance=",instance)
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if instance==nil {
		//StoreInt32原子性的将val的值保存到*addr。
		defer atomic.StoreInt32(&initialized,1)
		instance=&singleton{}
		fmt.Println("单利没有被生成 但是初始化了一个单利 instance=",instance)
	}
	return instance
}






var  totalTwo int32

func workerTwo(wg *sync.WaitGroup)  {
	defer wg.Done()
	//var i uint64
	for i:=0;i<=10 ;i++  {
		//atomic.AddInt32函数调用保证了total的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的
		atomic.AddInt32(&totalTwo,int32(i))
	}

}




var total struct{
	//模拟粗粒度的原子操作
	sync.Mutex
	value int
}
/*
WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束
 */
func worker(wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("每次执行前的值是  i=",total.value)
	for i:=0;i<=10 ;i++  {
		//fmt.Println("执行的次数----》  i=",i)
		total.Lock()
		// todo  total.value += i的原子性，我们通过sync.Mutex加锁和解锁来保证该语句在同一时刻只被一个线程访问 对于多线程模型的程序而言，进出临界区前后进行加锁和解锁都是必须的。如果没有锁的保护，total的最终值将由于多线程之间的竞争而可能会不正确
		total.value+=i
		total.Unlock()
	}
	fmt.Println("每次执行完了的值是  i=",total.value)

}








func goroutine() {
	//Goroutine是Go语言特有的并发体，是一种轻量级的线程，由go关键字启动。在真实的Go语言的实现中，goroutine和系统线程也不是等价的。尽管两者的区别实际上只是一个量的区别，但正是这个量变引发了Go语言并发编程质的飞跃

  //首先，每个系统级线程都会有一个固定大小的栈（一般默认可能是2MB），这个栈主要用来保存函数递归调用时参数和局部变量。固定了栈的大小这导致了两个问题：一是对于很多只需要很小的栈空间的线程来说是一个巨大的浪费，二是对于少数需要巨大栈空间的线程来说又面临栈溢出的风险。针对这两个问题的解决方案是：要么降低固定的栈大小，提升空间的利用率;要么增大栈的深度以允许更深的函数递归调用，但这两者是没法同时兼得的。相反，一个Goroutine会以一个很小的栈启动（可能是2KB或4KB），当遇到深度递归导致当前栈空间不足时，Goroutine会根据需要动态地伸缩栈的大小（主流实现中栈的最大值可达到1GB）。因为启动的代价很小，所以我们可以轻易地启动成千上万个Goroutine。



	//Go的运行时还包含了其自己的调度器，这个调度器使用了一些技术手段，可以在n个操作系统线程上多工调度m个Goroutine。Go调度器的工作和内核的调度是相似的，但是这个调度器只关注单独的Go程序中的Goroutine。Goroutine采用的是半抢占式的协作调度，只有在当前Goroutine发生阻塞时才会导致调度；同时发生在用户态，调度器会根据具体函数只保存必要的寄存器，切换的代价要比系统线程低得多。运行时有一个runtime.GOMAXPROCS变量，用于控制当前运行正常非阻塞Goroutine的系统线程数目。
	// todo 运行时有一个runtime.GOMAXPROCS变量，用于控制当前运行正常非阻塞Goroutine的系统线程数目。
//	runtime.GOMAXPROCS(10)

   //在Go语言中启动一个Goroutine不仅和调用函数一样简单，而且Goroutine之间调度代价也很低，这些因素极大地促进了并发编程的流行和发展




}