
# GoDemo <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>
### Old目录-->以前学习的Demo，停止更新

#### 2018.7.26 学习Deployment部署和Maintenance维护
* 1、日记应用 尼玛逼啊！还要一个依赖：https://github.com/golang/sys 
* 2、seelog的Demo完成
* 3、错误的处理：panic和recover是针对自己开发package里面实现的逻辑，针对一些特殊情况来设计。
* 5、错误的处理
*  6、网站错误处理：数据库错误（连接错误、查询错误、数据错误）；应用运行时错误（文件系统和权限、第三方应用和接口错误）；HTTP错误；操作系统出错；网络出错
*  7、错误处理的目标：通知访问用户出现错误了；记录错误；回滚当前的请求操作；保证现有程序可运行可服务
*  8、如何处理错误
*  9、应用部署： daemon：Go程序还不能实现daemon，详细的见这个Go语言的bug：<http://code.google.com/p/go/issues/detail?id=227>，大概的意思说很难从现有的使用的线程中fork一个出来，因为没有一种简单的方法来确保所有已经使用的线程的状态一致性问题
* 10、Supervisord可惜啊，不支持window系统啊日了狗 
* 11、备份和恢复

### New20180730目录,重新学习了一本书来了解更深刻的原理，
#### 说明一：如果见着 `todo  image`的字样，那么在`image`目录下就有图片的说明 
#### 说明二：要有一定经验才能学习这个目录的东西，没有经验学习`Old`的目录

* 2018.7.30 
  *  数组Array
* 2018.7.31
  *  字符串、切片（slice），切片内存技巧，避免切片内存的泄露！切片类型的强制转换，sort.Ints对转换后的[]int排序的性能要比用sort.Float64s排序的性能好一点
  *  Go语言函数的递归调用深度逻辑上没有限制，函数调用的栈是不会出现溢出错误的，因为Go语言运行时会根据需要动态地调整函数栈的大小。每个goroutine刚启动时只会分配很小的栈（4或8KB，具体依赖实现），根据需要动态调整栈的大小，栈最大可以达到GB级（依赖具体实现）。在Go1.4以前，Go的动态栈采用的是分段式的动态栈，通俗地说就是采用一个链表来实现动态栈，每个链表的节点内存位置不会发生变化。但是链表实现的动态栈对某些导致跨越链表不同节点的热点调用的性能影响较大，因为相邻的链表节点它们在内存位置一般不是相邻的，这会增加CPU高速缓存命中失败的几率。为了解决热点调用的CPU缓存命中率问题，Go1.4之后改用连续的动态栈实现，也就是采用一个类似动态数组的结构来表示栈。不过连续动态栈也带来了新的问题：当连续栈动态增长时，需要将之前的数据移动到新的内存空间，这会导致之前栈中全部变量的地址发生变化。虽然Go语言运行时会自动更新引用了地址变化的栈变量的指针，但最重要的一点是要明白Go语言中指针不再是固定不变的了（因此不能随意将指针保持到数值变量中，Go语言的地址也不能随意保存到不在GC控制的环境中，因此使用CGO时不能在C语言中长期持有Go语言对象的地址）
  *  Go语言的函数，导包方法执行的先后的顺序：通过日记输出，可以看出，先执行所有的pkg的init的方法，然后执行mian.init ，然后执行main.main 最后执行方法导入到方法
  * 我们无法知道函数参数或局部变量到底是保存在栈中还是堆中，我们只需要知道它们能够正常工作就可以了

* 2018.8.1
  * 方法：oop(面向对象的程序设计)  1、封装 2、继承 3、多态 4、抽象
  * 方法一般是面向对象编程(OOP)的一个特性 
  * 一般静态编程语言都有着严格的类型系统。过于严格的编译系统，会导致编程的效率过低 ---  go在其中取得平衡 
  * 鸭子类型：当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子
  * 面向并发的内心模型:Go语言是基于消息并发模型的集大成者，它将基于CSP模型的并发编程内置到了语言中，通过一个go关键字就可以轻易地启动一个Goroutine，与Erlang不同的是Go语言的Goroutine之间是共享内存的
  * Goroutine是Go语言特有的并发体，是一种轻量级的线程，由go关键字启动。在真实的Go语言的实现中，goroutine和系统线程也不是等价的。尽管两者的区别实际上只是一个量的区别，但正是这个量变引发了Go语言并发编程质的飞跃
  * 原子操作：所谓的原子操作就是并发编程中“最小的且不可并行化”的操作。通常，有多个并发体对一个共享资源的操作是原子操作的话，同一时刻最多只能有一个并发体对该资源进行操作！自己的话来讲：原子性及时一个操作或者是多个操作，要么全部执行并且执行的过程中不会被任何因素打断，要么就不执行
  * "原子操作(atomic operation)是不需要synchronized"，这是多线程编程的老生常谈了。所谓原子操作是指不会被线程调度机制打断的操作；这种操作一旦开始，就一直运行到结束，中间不会有任何 context switch （切换到另一个线程）
  * 在Java中有个关键字 volatile 就是关系到原子性：并发编程中，我们通常会遇到以下三个问题：原子性问题，可见性问题，有序性问题
  * https://www.jianshu.com/p/9080483bac91 （自己写过的Glide图片架构的封装有提到 volatile关键字）
  * 可以使用 `sync.Once`实现单利，也可以使用原子操作配合互斥锁可以实现非常高效的单利事件 
  * 所有的init函数和main函数都是在主线程完成，它们也是满足顺序一致性模型的,在main.main函数执行之前所有代码都运行在同一个goroutine中，也是运行在程序的主系统线程中。如果某个init函数内部用go关键字启动了新的goroutine的话，新的goroutine只有在进入main.main函数之后才可能被执行到
  * 在main方法中创建协程，有时候是不能执行的，需要使用channle，如果一个并发程序无法确定事件的偏序关系，那么程序的运行结果往往会有不确定的结果
  * 基于Channel的通信
  * 解决同步问题的思路是相同的：使用显式的同步
  
* 2018.8.2
  *  常见的并发的模式 
  *  `<-`  是对chan类型来说的。chan类型类似于一个数组。当`<- chan` 的时候是对chan中的数据读取；相反 `chan <- value` 是对chan赋值。
  *  `sync.WaitGroup`来等待一组事件，能够执行每一条语句完成了后才执行下一句，如果通过管道来等待一组事件，是等所有事件都来了，才开始执行，注意使用 `for` 来观察下这种的结果  
  * 生产者和消费者的模型:主要通过平衡生产线程和消费线程的工作能力来提高程序的整体处理数据的速度
  * 通过命令`go build -gcflags "-N -l" FourthCommo nConcurrencyPatterns.go` 可以把`go`的文件变成 windows可以执行的 `exe`的文件
  * 在发布订阅模型中，每条消息都会传送给多个订阅者。发布者通常不会知道、也不关心哪一个订阅者正在接收主题消息。订阅者和发布者可以在运行时动态添加是一种松散的耦合关心，这使得系统的复杂性可以随时间的推移而增长。在现实生活中，不同城市的象天气预报之类的应用就可以应用这个并发模式。
  * `PubAndSub`这个Demo 非常的有意思,基本上后续的思想都要基于这个Demo的原理
  * 多个 `channel `的情况，Go提供了一个关键字 `select` ，通过select可以监听 `channel` 上的数据流动
  * 写代码有个误区，不是说你的代码运行快！就好，要给用户的反馈的速度最快，才是最好的！记住这一点，哈哈！想起来我自己写的按个手写体，还是很强的
  * 并发版本的素数筛是一个经典的并发例子，通过它我们可以更深刻地理解Go语言的并发特性：质数（prime number）又称素数，有无限个。 质数定义为在大于1的自然数中，除了1和它本身以外不再有其他因数。
  * `chan` 不能乱结束啊 草 你结束了这个管道 ，后面的输入的就会有问题啊,素数就是质数
  * 素数筛展示了一种优雅的并发程序结构。但是因为每个并发体处理的任务粒度太细微，程序整体的性能并不理想。对于细力度的并发程序，CSP模型中固有的消息传递的代价太高了（多线程并发模型同样要面临线程启动的代价）
  * 并发的安全退出，使用 `select`关键字 
  * 需要结合到结合`sync.WaitGroup`来改进,让每个工作者并发体的创建、运行、暂停和退出都是在main函数的安全控制之下了
  * Go1.7发布时，标准库增加了一个context包，用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作
  * `context`包下的应用 :可以用context包来重新实现线程安全退出或超时的控制
* 2018.8.3
   * 错误和异常：错误处理是每个编程语言都要考虑的一个重要话题。在Go语言的错误处理中，错误是软件包API和应用程序用户界面的一个重要组成部分。
   * `recover()`函数的使用，说到底就是，及时发生了异常，我们也得继续运行，不能让用户不能使用了，这个很关键！
   * CGO编程：C语言作为一个通用语言，很多库会选择提供一个C兼容的API，然后用其他不同的编程语言实现。Go语言通过自带的一个叫CGO的工具来支持C语言函数调用，同时我们可以用Go语言导出C动态库接口给其它语言使用。开发区块链，就是通过Go生成so库给安卓使用，有点意思
   * cgo 是让 Go 程序在 Android 和 iOS 上运行的关键。
   * 安装必要的环境：`https://blog.csdn.net/mecho/article/details/24305369` 通过运行一个，第一个真正的`CGO`的程序！！
   * `C` 代码的模块化，抽象和模块化是将复杂文件简化的通用手段。模块化编程的核心是面向程序接口编程 
   * 用Go重新实现C函数：C.h 有个函数，其中的方法，交个Go去实现，有个关键的调用 `   //export SayThirdHello `  这个标记也是起作用的哦！
   ```
   //export SayThirdHello
   func SayThirdHello(s *C.char)  {
   	fmt.Println("不要管我  我肯定会执行 ")
      fmt.Println(C.GoString(s))
   }
   ``` 
   * 在Go1.10中CGO新增加了一个_GoString_预定义的C语言类型，用来表示Go语言字符串，但是由于我的是1.9.2 ，这个Demo没有成功的跑起来，所以说，这是个遗憾 
   * 面向C接口的Go编程：主要是两个Demo: ` FourthGoProgrammingForCInterface` 和 `FourthGoProgrammingForCInterfaceBetter`
   *  类型转换:最初CGO是为了达到方便从Go语言函数调用C语言函数以复用C语言资源这一目的而出现的(因为C语言还会涉及回调函数，自然也会涉及到从C语言函数调用Go语言函数)。现在，它已经演变为C语言和Go语言双向通讯的桥梁。要想利用好CGO特性，自然需要了解此二语言类型之间的转换规则
   
   * Go语言中数值类型和C语言数据类型基本上是相似的，以下是它们的对应关系表。
   
   C语言类型               | CGO类型      | Go语言类型
   ---------------------- | ----------- | ---------
   char                   | C.char      | byte
   singed char            | C.schar     | int8
   unsigned char          | C.uchar     | uint8
   short                  | C.short     | int16
   unsigned short         | C.ushort     | uint16
   int                    | C.int       | int32
   unsigned int           | C.uint      | uint32
   long                   | C.long      | int32
   unsigned long          | C.ulong     | uint32
   long long int          | C.longlong  | int64
   unsigned long long int | C.ulonglong | uint64
   float                  | C.float     | float32
   double                 | C.double    | float64
   size_t                 | C.size_t    | uint
* 2018.8.6
   * 除了`GoInt`和`GoUint`之外，我们并不推荐直接访问`GoInt32`、`GoInt64`等类型。更好的做法是通过C语言的C99标准引入的`<stdint.h>`头文件。为了提高C语言的可移植性，在`<stdint.h>`文件中，不但每个数值类型都提供了明确内存大小，而且和Go语言的类型命名更加一致
   
   C语言类型 | CGO类型     | Go语言类型
   -------- | ---------- | ---------
   int8_t   | C.int8_t   | int8
   uint8_t  | C.uint8_t  | uint8
   int16_t  | C.int16_t  | int16
   uint16_t | C.uint16_t | uint16
   int32_t  | C.int32_t  | int32
   uint32_t | C.uint32_t | uint32
   int64_t  | C.int64_t  | int64
   uint64_t | C.uint64_t | uint64
    *  Go 字符串和切片
    * 在CGO生成的`_cgo_export.h`头文件中还会为Go语言的字符串、切片、字典、接口和管道等特有的数据类型生成对应的C语言类型：
      
      ```c
      typedef struct { const char *p; GoInt n; } GoString;
      typedef void *GoMap;
      typedef void *GoChan;
      typedef struct { void *t; void *v; } GoInterface;
      typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;
      ```
      
      不过需要注意的是，其中只有字符串和切片在CGO中有一定的使用价值，因为此二者可以在Go调用C语言函数时马上使用;而CGO并未针对其他的类型提供相关的辅助函数，且Go语言特有的内存模型导致我们无法保持这些由Go语言管理的内存指针，所以它们C语言环境并无使用的价值。
      
      在导出的C语言函数中我们可以直接使用Go字符串和切片。假设有以下两个导出函数：
      
      ```go
      //export helloString
      func helloString(s string) {}
      
      //export helloSlice
      func helloSlice(s []byte) {}
      ```
      
      CGO生成的`_cgo_export.h`头文件会包含以下的函数声明：
      
      ```c
      extern void helloString(GoString p0);
      extern void helloSlice(GoSlice p0);
      ```
      
      不过需要注意的是，如果使用了GoString类型则会对`_cgo_export.h`头文件产生依赖，而这个头文件是动态输出的。
      
      Go1.10针对Go字符串增加了一个`_GoString_`预定义类型，可以降低在cgo代码中可能对`_cgo_export.h`头文件产生的循环依赖的风险。我们可以调整helloString函数的C语言声明为：
      
      ```c
      extern void helloString(_GoString_ p0);
      ```
      
      因为`_GoString_`是预定义类型，我们无法通过此类型直接访问字符串的长度和指针等信息。Go1.10同时也增加了以下两个函数用于获取字符串结构中的长度和指针信息：
      
      ```c
      size_t _GoStringLen(_GoString_ s);
      const char *_GoStringPtr(_GoString_ s);
      ```
      
      更严谨的做法是为C语言函数接口定义严格的头文件，然后基于稳定的头文件实现代码。
  * 结构体、联合、枚举类型
   * 如果结构体的成员名字中碰巧是Go语言的关键字，可以通过在成员名开头添加下划线来访问
           ```
           /*
           struct A {
           	int type; // type 是 Go 语言的关键字
           };
           */
           import "C"
           import "fmt"
           
           func main() {
           	var a C.struct_A
           	fmt.Println(a._type) // _type 对应 type
           }
           ```
    *   一个是以Go语言关键字命名，另一个刚好是以下划线和Go语言关键字命名，那么以Go语言关键字命名的成员将无法访问（被屏蔽）      
    *  C语言结构体中位字段对应的成员无法在Go语言中访问，如果需要操作位字段成员，需要通过在C语言中定义辅助函数来完成。对应零长数组的成员，无法在Go语言中直接访问数组的元素，但其中零长的数组成员所在位置的偏移量依然可以通过unsafe.Offsetof(a.arr)来访问     
    *  对于联合类型，我们可以通过C.union_xxx来访问C语言中定义的union xxx类型。但是Go语言中并不支持C语言联合类型，它们会被转为对应大小的字节数组
    *  需要操作C语言的联合类型变量，一般有三种方法：第一种是在C语言中定义辅助函数；第二种是通过Go语言的"encoding/binary"手工解码成员(需要注意大端小端问题)；第三种是使用unsafe包强制转型为对应类型(这是性能最好的方式)
      
    
       union B {
       	int i;
      	float f;
        };
        var b C.union_B;
     	fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&b)))
     	fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&b)))
     	

      
  *  对于枚举类型，我们可以通过C.enum_xxx来访问C语言中定义的enum xxx结构体类型
   ```go
   /*
   enum C {
   	ONE,
   	TWO,
   };
   */
   import "C"
   import "fmt"
   
   func main() {
   	var c C.enum_C = C.TWO
   	fmt.Println(c)
   	fmt.Println(C.ONE)
   	fmt.Println(C.TWO)
   }
   ```
  
   *  Go语言的字符串是只读的，用户需要自己保证Go字符串在使用期间，底层对应的C字符串内容不会发生变化、内存不会被提前释放掉

      