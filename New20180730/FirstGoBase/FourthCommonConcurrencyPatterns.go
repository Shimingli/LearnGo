package main

import (
	"fmt"
	"sync"
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
