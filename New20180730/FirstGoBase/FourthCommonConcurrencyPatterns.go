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
	"context"
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




	//发布订阅模型
	publishAndSubscribe()


	//控制并发数
	ControlConcurrencyNumber()



	//赢者为王
	winnerisKing()


   // 素数筛  质数（prime number）又称素数，有无限个。 质数定义为在大于1的自然数中，除了1和它本身以外不再有其他因数。
   primeNumberScreen()
	//todo image


	//生产者和消费者的模型   为了更好地演示Demo 把这个给注释掉了
	//producerAndConsumerModel()
    //并发的安全退出
	concurrentSecurityExit()
	//我们通过close来关闭cancel管道向多个Goroutine广播退出的指令。不过这个程序依然不够稳健：当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，但是退出的清理工作并不能保证被完成，因为main线程并没有等待各个工作Goroutine退出工作完成的机制。我们可以结合sync.WaitGroup来改进
  //现在每个工作者并发体的创建、运行、暂停和退出都是在main函数的安全控制之下了
	//concurrentSecurityExitMoreGood()


	//可以用context包来重新实现前面的线程安全退出或超时的控制
	//Go1.7发布时，标准库增加了一个context包，用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作
	contextDemo()


       //Go语言是带内存自动回收的特性，因此内存一般不会泄漏。在前面素数筛的例子中，GenerateNatural和PrimeFilter函数内部都启动了新的Goroutine，当main函数不再使用管道时后台Goroutine有泄漏的风险。我们可以通过context包来避免这个问题，下面是改进的素数筛实现
    contextDemo2()

}

func contextDemo2() {
	// 通过 Context 控制后台Goroutine状态
	ctx, cancel := context.WithCancel(context.Background())
     //当main函数完成工作前，通过调用cancel()来通知后台Goroutine退出，这样就避免了Goroutine的泄漏
	ch := GenerateNaturalTwo(ctx) // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilterTwo(ctx, ch, prime) // 基于新素数构造的过滤器
	}
	cancel()
}

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNaturalTwo(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <- ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilterTwo(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <- ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func contextDemo() {
	//当并发体超时或main主动停止工作者Goroutine时，每个工作者都可以安全退出
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerContext(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}

func workerContext(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello  workerContext  ")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}





func concurrentSecurityExitMoreGood() {
	cancel := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerMoreGood(&wg, cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
func workerMoreGood(wg *sync.WaitGroup, cannel chan bool) {
	defer wg.Done()
     i:=0
	for {
		select {
		default:
			fmt.Println("我在工作哦")
		case <-cannel:
			i++
			fmt.Println("我没有工作了哦      --",i)
			return
		}
	}
}


/*
有时候我们需要通知goroutine停止它正在干的事情，特别是当它工作在错误的方向上的时候。Go语言并没有提供在一个直接终止Goroutine的方法，由于这样会导致goroutine之间的共享变量落在未定义的状态上。但是如果我们想要退出两个或者任意多个Goroutine怎么办呢？
Go语言中不同Goroutine之间主要依靠管道进行通信和同步。要同时处理多个管道的发送或接收操作，我们需要使用select关键字（这个关键字和网络编程中的select函数的行为类似）。当select有多个分支时，会随机选择一个可用的管道分支，如果没有可用的管道分支则选择default分支，否则会一直保存阻塞状态
 */
func concurrentSecurityExit() {
	//基于select实现的管道的超时判断
	//select {
	//case v := <-in:
	//	fmt.Println(v)
	//case <-time.After(time.Second):
	//	return // 超时
	//}


	//通过select的default分支实现非阻塞的管道发送或接收操作：

	//select {
	//case v := <-in:
	//	fmt.Println(v)
	//default:
	//	// 没有数据
	//}

	//通过select来阻止main函数退出：
		// do some thins
		//select{}

	//当有多个管道均可操作时，select会随机选择一个管道。基于该特性我们可以用select实现一个生成随机数列的程序：
	//ch := make(chan int)
	//go func() {
	//	for {
	//		select {
	//
	//		case ch <- 0: //从c中接收数据，并赋值给v  如果成功的话 就打印出来 v := <- c
	//			fmt.Println()
	//		case ch <- 1:
	//		}
	//	}
	//}()
	////安全的退出并发---》 0
	////安全的退出并发---》 1
	//for v := range ch {
	//	fmt.Println("安全的退出并发---》",v)
	//}




	//cannel1 := make(chan bool)
	//go workerd(cannel1)
	//
	//time.Sleep(time.Second)
	//cannel1 <- true

	//但是管道的发送操作和接收操作是一一对应的，如果要停止多个Goroutine那么可能需要创建同样数量的管道，这个代价太大了。其实我们可以通过close关闭一个管道来实现广播的效果，所有从关闭管道接收的操作均会收到一个零值和一个可选的失败标志。
	//cancelTwoTwo := make(chan bool)
	//
	//for i := 0; i < 10; i++ {
	//	go workerTwoT(cancelTwoTwo)
	//}
	//time.Sleep(time.Second)
	//close(cancelTwoTwo)
}

func workerTwoT(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("我在工作哦")
			// 正常工作
		case <-cannel:
			fmt.Println("让我好好的退出下哈哈")
			// 退出
		}
	}
}

//我们通过select和default分支可以很容易实现一个Goroutine的退出控制:
func workerd(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			// 正常工作
		case <-cannel:
			// 退出
		}
	}
}


//并发版本的素数筛是一个经典的并发例子，通过它我们可以更深刻地理解Go语言的并发特性
func primeNumberScreen() {
	var primeInt [10000]int
	// 返回生成自然数序列的管道: 2, 3, 4, ...
	ch:=GenerateNatural()
	//不能乱结束啊 草 你结束了这个管道 ，后面的输入的就会有问题啊
    //fmt.Println("ch=",<-ch)
	for i := 0; i < 10000; i++ {
		prime := <-ch // 新出现的素数
		fmt.Println("第",i+1,"次循环,prime的值是=",prime)
		//str:=strconv.Itoa(i+1)
		primeInt[i]=prime
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}


	fmt.Println(primeInt)


}

//然后是为每个素数构造一个筛子：将输入序列中是素数倍数的数提出，并返回新的序列，是一个新的管道
//管道过滤，删除能被素数整除的数
func PrimeFilter(in <-chan int,prime int)  chan int{
    out :=make(chan int)
    go func() {
    	for  {
    		//第一次进来 i=2
    		if i:= <-in;i%prime!=0{
    			out<-i
			}
		}
	}()
    //fmt.Println("返回回去的值====",<-out)
    return out

}



//需要先生成最初的自然数序列，不包含 0 和 1
//GenerateNatural函数内部启动一个Goroutine生产序列，返回对应的管道。
func GenerateNatural() chan int{
	ch:=make(chan int)
	go func() {
		for i:=2;;i++{
			ch <-i
		}
	}()
	return ch
}




/*
采用并发编程的动机有很多：并发编程可以简化问题，比如一类问题对应一个处理线程会更简单；并发编程还可以提升性能，在一个多核CPU上开2个线程一般会比开1个线程快一些。其实对于提升性能而言，程序并不是简单地运行速度快就表示用户体验好的；很多时候程序能快速响应用户请求才是最重要的，当没有用户请求需要处理的时候才合适处理一些低优先级的后台任务。
 */
func winnerisKing() {
   //假设我们想快速地检索“golang”相关的主题，我们可能会同时打开Bing、Google或百度等多个检索引擎。当某个检索最先返回结果后，就可以关闭其它检索页面了。因为受限于网络环境和检索引擎算法的影响，某些检索引擎可能很快返回检索结果，某些检索引擎也可能遇到等到他们公司倒闭也没有完成检索的情况。我们可以采用类似的策略来编写这个程序


	ch := make(chan string, 32)

	go func() {
		fmt.Println("golang")
		ch<- "bing"
		//ch <- searchByBing("golang")
	}()
	go func() {
		fmt.Println("Google")
		ch<- "Google"
	//	ch <- searchByGoogle("golang")
	}()
	go func() {
		fmt.Println("Baidu")
		ch<- "Baidu"
		//ch <- searchByBaidu("golang")
	}()


	/*
	创建了一个带缓存的管道，管道的缓存数目要足够大，保证不会因为缓存的容量引起不必要的阻塞。然后我们开启了多个后台线程，分别向不同的检索引擎提交检索请求。当任意一个检索引擎最先有结果之后，都会马上将结果发到管道中（因为管道带了足够的缓存，这个过程不会阻塞）。但是最终我们只从管道取第一个结果，也就是最先返回的结果。
通过适当开启一些冗余的线程，尝试用不同途径去解决同样的问题，最终以赢者为王的方式提升了程序的相应性能
	 */
	fmt.Println(<-ch)

}
/*
很多用户在适应了Go语言强大的并发特性之后，都倾向于编写最大并发的程序，因为这样似乎可以提供最大的性能。在现实中我们行色匆匆，但有时却需要我们放慢脚步享受生活，并发的程序也是一样：有时候我们需要适当地控制并发的程度，因为这样不仅仅可给其它的应用/任务让出/预留一定的CPU资源，也可以适当降低功耗缓解电池的压力。
 */
func ControlConcurrencyNumber() {
	//gatefs子包的目的就是为了控制访问该虚拟文件系统的最大并发数
	// 其中vfs.OS("/path")基于本地文件系统构造一个虚拟的文件系统，然后gatefs.New基于现有的虚拟文件系统构造一个并发受控的虚拟文件系统
	//fs := gatefs.New(vfs.OS("/path"), make(chan bool, 8))
	//select 默认是阻塞的，只有当监听的 channel 中有发送或者接受可以进行时才会运动，当多个channel 都准备好了，select是随机选择一个执行的

	//var limit = make(chan int, 3)
	//
	//func main() {
	//	for _, w := range work {
	//		go func() {
	//			limit <- 1
	//			w()
	//			<-limit
	//		}()
	//	}
	//	select{}
	//}

  //我们不仅可以控制最大的并发数目，而且可以通过带缓存Channel的使用量和最大容量比例来判断程序运行的并发率。当管道为空的时候可以认为是空闲状态，当管道满了时任务是繁忙状态，这对于后台一些低级任务的运行是有参考价值的




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
