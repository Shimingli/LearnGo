package main

import (
	"fmt"
	"sync"
	"os"
	"os/signal"
	"syscall"
	"time"
	"GoDemo/New20180730/FirstGoBase/pubsub"
	"strings"
)

func init() {
	fmt.Println("常见的并发模式")

}

func main() {
	//在并发编程中，对共享资源的正确访问需要精确的控制，在目前的绝大多数语言中，都是通过加锁等线程同步方案来解决这一困难问题，而Go语言却另辟蹊径，它将共享的值通过信道传递(实际上多个独立执行的线程很少主动共享资源)。在任意给定的时刻，最好只有一个Goroutine能够拥有该资源。数据竞争从设计层面上就被杜绝了。为了提倡这种思考方式，Go语言将其并发编程哲学化为一句口号
	//Do not communicate by sharing memory; instead, share memory by communicating.
	//不要通过共享内存来通信，而应通过通信来共享内存。

	//并发版本的Hello world

	concurrencyHelloWorld()

	//生产者和消费者的模型
	//producerAndConsumerModel()


	//发布订阅模型
	publishAndSubscribe()




}
/*
发布／订阅（publish-and-subscribe）模型通常被简写为pub／sub模型。在这个模型中，消息生产者成为发布者（publisher），而消息消费者则称对应订阅者（subscriber），生产者和消费者是M：N的关系。在传统生产者和消费者模型中，成果是将消息发送到一个队列中，而发布/订阅模型则是将消息发布给一个主题
 */
func publishAndSubscribe() {
	//  构建一个发布者对象，可以设置发布的超时的时间和缓存队列的的长度
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()
     //添加一个新的订阅者，订阅全部的主题
	all := p.SubscibeAll()
	// 添加一个新的订阅者，订阅过滤器筛选后的主题
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "信息")
		}
		return false
	})

	p.Publish("我发布了一个信息 谁能收到")
	p.Publish("我没有关键字 看谁能收到")

	go func() {
		for  msg := range all {
			fmt.Println("全部接收到的信息:", msg)
		}
	} ()

	go func() {
		for  msg := range golang {
			fmt.Println("包含了关键字的`信息`:", msg)
		}
	} ()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
	//在发布订阅模型中，每条消息都会传送给多个订阅者。发布者通常不会知道、也不关心哪一个订阅者正在接收主题消息。订阅者和发布者可以在运行时动态添加是一种松散的耦合关心，这使得系统的复杂性可以随时间的推移而增长。在现实生活中，不同城市的象天气预报之类的应用就可以应用这个并发模式。

}
/*
并发编程中最常见的例子就是生产者/消费者模式，该模式主要通过平衡生产线程和消费线程的工作能力来提高程序的整体处理数据的速度。简单地说，就是生产者生产一些数据，然后放到成果队列中，同时消费者从成果队列中来取这些数据。这样就让生产消费变成了异步的两个过程。当成果队列中没有数据时，消费者就进入饥饿的等待中；而当成果队列中数据已满时，生产者则面临因产品挤压导致CPU被剥夺的下岗问题。
 */
func producerAndConsumerModel() {
   ch :=make(chan int,64)
    //启了2个Producer生产流水线，分别用于生成3和5的倍数的序列
   go Producer(3,ch)
   go Producer(5,ch)

   go Consumer(ch)
   //这种靠休眠方式是无法保证稳定的输出结果的

   //time.Sleep(time.Second*5)

   //E:\new_code\GoDemo\New20180730\FirstGoBase>go build -gcflags "-N -l" FourthCommo nConcurrencyPatterns.go
	//todo 让main函数保存阻塞状态不退出，只有当用户输入Ctrl-C时才真正退出程序
   sig:=make(chan os.Signal,1)
   signal.Notify(sig,syscall.SIGINT,syscall.SIGTERM)
   fmt.Printf("quit (%v)\n",<-sig)

   //有2个生产者，并且2个生产者之间并无同步事件可参考，它们是并发的。因此，消费者输出的结果序列的顺序是不确定的，这并没有问题，生产者和消费者依然可以相互配合工作

}
//生产者 生产factor 的倍数
func Producer(factor int,out chan <-int)  {
	for i:=0;;i++  {
		out<-i*factor
	}
}

func Consumer(in <-chan int)  {
	i:=0
	for v :=range in{

		fmt.Println("打印的 Value=",v,"in=",in)
		i++
		fmt.Println("输入了第几次了",i,"宝宝我爱你----")
	}
}






func concurrencyHelloWorld() {
    //下面的方式 会出现异常
	//var   mu sync.Mutex
	//go func() {
		//fmt.Println("你会输出么？？？其实我知道 你可能不会输出")
		////
		//mu.Lock()
	//}()
	////我们不能直接对一个未加锁状态的sync.Mutex进行解锁，这会导致运行时异常。mu.Lock()和mu.Unlock()并不在同一个Goroutine中，所以也就不满足顺序一致性内存模型
	//mu.Unlock()
	//// 因为可能是并发的事件，所以main函数中的mu.Unlock()很有可能先发生，而这个时刻mu互斥对象还处于未加锁的状态，从而会导致运行时异常


	   //修复后的代码

	   //修复的方式是在main函数所在线程中执行两次mu.Lock()，当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，main函数的阻塞状态驱动后台线程继续向前执行。当后台线程执行到mu.Unlock()时解锁，此时打印工作已经完成了，解锁会导致main函数中的第二个mu.Lock()阻塞状态取消，此时后台线程和主线程再没有其它的同步事件参考，它们退出的事件将是并发的：在main函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的。
		var mu sync.Mutex

		mu.Lock()
		go func(){
			fmt.Println("你好, 世界")
			mu.Unlock()
		}()

		mu.Lock()



	//sync.Mutex互斥锁同步是比较低级的做法。我们现在改用无缓存的管道来实现同步
    //根据Go语言内存模型规范，对于从无缓冲信道进行的接收，发生在对该信道进行的发送完成之前。因此，后台线程<-done接收操作完成之后，main线程的done <- 1发生操作才可能完成（从而退出main、退出程序），而此时打印工作已经完成了
       done:= make(chan int)

       go func() {
       	fmt.Println("你好啊 ！ shiming")

       	//  <-  是对chan类型来说的。chan类型类似于一个数组。
       	//当<- chan 的时候是对chan中的数据读取；
		 //相反 chan <- value 是对chan赋值。
       	<-done
	   }()
       done<-1

       //上面的代码虽然可以正确同步，但是对管道的缓存大小太敏感：如果管道有缓存的话，就无法保证能main退出之前后台线程能正常打印了。更好的做法是将管道的发送和接收方向调换一下，这样可以避免同步事件受管道缓存大小的影响

    //带缓存的管道
    done1 := make(chan int,1)
    //对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是Channel的缓存大小。虽然管道是带缓存的，main线程接收完成是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的
    go func() {
    	fmt.Println("你好啊 小司机")
		done1<-1
	}()
    <-done1


	//基于带缓存的管道，我们可以很容易将打印线程扩展到N个
	//带10个缓存
	doneTwo:=make(chan int,10)
	for i:=0;i<cap(doneTwo) ;i++  {
		//这里有个很 有趣的现象-- 就是i=10 又有，才开始执行输出的语句
		go func() {
			fmt.Println("shiming woaini i=",i)
			doneTwo<-1
		}()
	}
	//等待N个后代线程的完成
	for i:=0;i<cap(doneTwo) ;i++  {
		<-doneTwo
	}



	//对于这种要等待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用sync.WaitGroup来等待一组事件
	var   wg sync.WaitGroup
	// 开启N个后台的打印的线程
	for i:=0; i<10;i++  {
		//其中wg.Add(1)用于增加等待事件的个数，必须确保在后台线程启动之前执行（如果放到后台线程之中执行则不能保证被正常执行到）。当后台线程完成打印工作之后，调用wg.Done()表示完成一个事件。main函数的wg.Wait()是等待全部的事件完成。
		wg.Add(1)
		go func() {
			//todo 这里的区别是  每次打印的时候 都会 执行正确的数字哦 i=1  i=2
			fmt.Println("nihao xiaojiejie  i=",i)
			wg.Done()
		}()
		// 等待后台线程完成
		wg.Wait()
	}



}
