concurrency

* 关注嘴,迈开腿: 管好goroutine的生命周期, 什么时候创建,什么时候销毁,一定要在你的掌握之中;

* goroutine不能自动销毁,给使用者判断什么时候开始,什么时候应该结束.

* 并发: 原子  +  可见性.

* happens before

* 老式CPU,锁总线. --> 新版CPU,MESI.

* 内存屏障.

* 熔断->交给一个人重试->CAS

* chan: 在通道两端同时有两个goroutine.

* context: 方便级联取消/超时,元数据的显示传递.

* 程序执行过程:

  * 源码进行编译后,数据和对数据进行操作的指令(函数:一套输入/运算/输出指令集合)分清了,他们的关系也有了.然后程序代码被操纵系统加载到内存中,叫进程,操作系统通过进程来控制程序代码,包括指令和数据(输入数据/输出数据,给数据起名,叫变量),.

  * CPU从内存中加载指令的成套集合(函数)到cache中进行快速运算,CPU是加载成套的指令集合(函数)为单元,进行运算执行的,这些许许多多的函数,在进程中排序是有规律的,哪个函数是入口,哪个函数是子函数,哪个函数是父函数,哪些函数是兄弟函数.这是在一个线程中执行的.

    > 非常类似场景: scene 就是进程, node 就是函数. 一次正常的对话交互,也就是说,这个执行路径walk node就是一个线程thread.  如果是从入口节点开始执行,那么这就是一个主线程,不是从入口节点开始的都是子线程.

  * 线程是CPU执行的. **cpu执行的是线程,而不是进程,进程只是在内存中的对程序的表示.**

  * 主线程是CPU自动执行的.

  * 子线程/子子线程是go 关键字 + node(就是函数)

  * runtime调度器通过一定的算法调度goroutine在不同的CPU上执行.

    * runtime里面有个P(一个数组),所有的G都挂接到P中.

    * P又挂接到M上.

    * GPM关系图.

      ![](F:\陈晓会\GO\极客大学\Go-000\Week03\GPM.bmp)

  * P:逻辑CPU.是操作系统调度单元M的逻辑单元,本来M只能挂一个珠串,但在P中可以挂无数个珠串. 只是这些goroutine珠串是**并发**执行的,不能并行执行. 如果在不同的M中P上挂接珠串,那才是**并行**

  * 逻辑隔离出这么多挂接点的目的: 为了保证隔离goroutine.  

  * 类似硬盘MBR分区,最多分三个主分区,再分主分区的话,就限制了分区了.和一个扩展分区,

* keep yourself busy or do the work yourself.

* 调用os.exit 会导致defer无法执行.

* Never  start a goroutine without knowning when it will stop.

* 什么时候将阻塞(死循环也是阻塞)委派: 当自己goroutine 忙不过来的时候.将阻塞委派出去.

* Any time you start a goroutine you must ask yourself 

  > When will it terminate? 通常是什么时候结束,它告诉调用者,什么时候结束.
  >
  > What could prevent it from terminating? 如何控制其可以中途结束而不是直到最后的terminal;虽然不知道它什么时候结束,但调用者可以控制其结束的判断条件,这里可以通过显示插入一个控制它结束的函数/方法.

* 性能调优挂掉, 应用也必须停掉.!!!!

* only use log.Fatal from main.main or init functions.

* ![](F:\陈晓会\GO\极客大学\Go-000\Week03\knowwhentistop.png)

* leave concurrency to the caller

* 如果函数启动goroutine,则必须向调用方提供显示停止该goroutine的方法. 类似filepath.Walk() 当开启函数,虽然我不知道什么时候它会自动结束时, 但我可以控制其结束.

* 一个函数要等很久才可能进行下一步,要考虑,是否要等,如果不想等,那请go 出去,并且通过超时控制来控制其返回,此外还有出问题了/或正常退出返回.

  ![](F:\陈晓会\GO\极客大学\Go-000\Week03\goroutine超时处理.png)

* 通过sync.WaitGroup来控制goroutine退出.

  ![](F:\陈晓会\GO\极客大学\Go-000\Week03\处理waitgroup超时.png)

* 重要点: 类似4次挥手

  1.   控制子线程什么时候退出:    数据传输完毕,close(chan).
  2.    知道最终子线程什么时候退出了.   子线程发送信号到stop通道.

* `Memory model`  内存模型

  1. 存在公共变量. 

  2. 怎么保证在一个goroutine中看到另一个goruotine修改的变量值:    如果一个goroutine对公共变量有写操作,另一个goroutine对公共变量有读操作,必须将读操作**串行化**.      **Happen_Before**   如果不跟读协调好,可能读到不同的值.

  3. 怎么**串行化**:   同步事件.

     1. **channel**
     2. 同步原语: **sync**或**sync/atomic**.

  4. go内存模型: 当一个goroutine对公共变量修改 时,另一个goroutine看不到其修改,原因,L1缓存.缓存公共变量数据. 

  5. `Memory Reordering`  内存重排

  6. 编译器重排BUG: 当存在多个goroutine时,可能多个goroutine会相互影响.而此时编译器是不清楚这些影响的,如果还要考虑重排,会导致BUG.

  7. cache line: store buffer: 针对单线程是完全没有问题的.

  8. CPU锁: barrier/fence->内存屏障. flush缓存到内存.    atomic compare-and-swap 这是依据CPU锁实现的. 使用标准库即可.  解决内存重排BUG.

  9. atomic操作: 迅速将缓存flush到内存.

  10. **Happens Before**  **先行发生** 这是一种条件,状态   谁先谁后:两个goroutine里面处理不同事件, 谁先谁后是非常必要的.

      > 2个事件发生的顺序导致两种情况:
      >
      > ​	一:  存在先后顺序.  event2发生在event1后.  单goroutine
      >
      > ​	二:  不存在先后顺序.  event1 不发生在event2前,也不发生在event2后,也称它们是**并发的**.  多goroutines. 这种不存在先后顺序的逻辑.在生活中经常发生,因为他们是完全不相干的东西. 比如比较变长等于2的面积和周长.哪个大.这从逻辑上说,是完全没有比较性的. 从计算过程上讲,计算两个完全不搭边的东西.无所谓谁先谁后. 因为就算计算出来了.也没有任何关联性.不知道后面要干什么. 这就是**并发的**.
      >
      > ​	三: 并发会导致问题,要通过happens before解决.也就是染个没有先后顺序的事件,人工强制进行先后排列事件.

  11.  事件分类

      1. 读操作: fmt/ 条件判断
      2. 写操作: 赋值,  声明变量时的零值

  12.  分析读操作&写操作.

      1. 读操作想看到写操作的变量必须满足: **写后面紧挨着读**:那么写是读看到的唯一写.
         1. 读在写后.
         2. 读/写之间没有其他对公共变量的写.
         3. 其他对公共变量的写:要么在写前,要么在读后.  而不能是并发的,无控制状态的.

  13. **将相关事件一定要修改成happens before, 而不能是并发的**

  14. 单线程**先天**存在读写时间顺序,只要代码顺序复合happens before,那么读就可以看到写;  而并发编程**先天**不存在读写时间顺序,要通过channel/同步原语来实现happens before.

  15. 

      

* **竞争检测器**:go build -race/ go test -race  查看data race  

* single machine word 

* 为什么通过channel可以实现happens before:   channel一定是要先发送数据给ch,才能从中读取.



# SYNC

* Share Memory By Communicating

  > Do not communicate by sharing memory, instead, share memory by communicating .

* 死锁: 你在电影院等她,她在电影院等你.

* i++:  查汇编`go tool compile -S 文件名(*.go)`

* 上下文切换: 不同goroutine的代码,执行一半就跳到另一个代码的地方了.

* data race 数据竞争.

* map或指针:做原子替换,比较安全

* 锁技巧: 在代码中  **最晚加锁,最早释放**  **锁里面代码最短** 

* atomic.value

* **读多写少**: 用读写锁  RWMutex

* 互斥锁可以被用来保护一个临界区

* go test -bench=. config_test.go   通过benchmark来判断是否使用Mutex还是使用atom

* _, _ = r, c  如果不想用可以省略

* COPY ON WRITE 

* 锁实现: 

  * Barging   冲撞/乱冲
  * Handsoft   手持式
  * Spinning    自旋

* ERRGROUP用法:

  * 如果有报错,全部取消,通过context实现级联取消.
  * 如果有报错,降级处理.
  * 好处: 不产生data race, 由于使用多个局部变量和闭包,每个goroutine都在处理独立的变量.

* sync.Pool: 用于高频的内存申请,可以减缓重新申请内存的操作,保留前一次的内存结构..例如: request driven











# CONTEXT  上下文

* 请求作用域的上下文(request-scoped context) 

* 跨goroutine共享数据,并控制超时,

* 使用方式:

  1. 在需要开goroutine时候,传入context.

     > The first parameter of a function call
     >
     > Optional config on a request structure 

  2. 仅将context用于跨越进程和API的请求范围的数据，而不是用于向函数传递可选参数。

  3. context要放到函数签名中`func DoSomething(ctx context.Context, arg Arg) error`

     > Do not store Contexts inside a struct type 

  4. context在程序代码中是流动的,贯穿性的.还是说,要在函数中作为参数传递.而不应该使用结构体存储.

  5. Context.Value should inform, not control

  6. The chain of function calls between them must propagate the Context.

* 好处:

  1. 可以输入特定于请求的元数据(request-specific)
  2. 输入取消信号,可以控制所有go出去的goroutine级联取消.
  3. 输入截止日期,可以解决超时问题.

* 坏处:

  1. 所有goroutine都会被污染.

* 核心:   超时处理

* 计算密集型: 耗时非常短,不需要超时处理. 根本管不住生命周期.  `二分查找`

* setreaddeadline 

* Notes

  > Incoming requests to a server should create a Context.
  >
  > Outgoing calls to servers should accept a Context.
  >
  > Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
  >
  > The chain of function calls between them must propagate the Context.
  >
  > Replace a Context using WithCancel, WithDeadline, WithTimeout, or WithValue.
  >
  > When a Context is canceld, all Contexts derived from it are alse canceled.
  >
  > The same Context may be passed to functions running in different goroutines; Context are safe for simultaneous use by multiple goroutines.
  >
  > Do not pass a nil Context,even if a function permits it.Pass a TODO context if you are unsure about which Context to use.
  >
  > Use context values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
  >
  > All blocking/long operations should be cancelable.
  >
  > Context.Value obscures your program's flow.
  >
  > Context.Value should inform, not control.
  >
  > Try not to use context.Value.







# CHAN

>  Do not communicate by sharing memory; instead, share memory by communicating.

* 无缓冲channel原理

  ![](F:\陈晓会\GO\极客大学\Go-000\Week03\unbuffered_channel.png)

* 无缓冲通道的本质是保证同步.

  > Receive 不晚于 Send发生.
  >
  > 好处: 100% 保证能收到.
  >
  > 代价: 延迟时间未知.

* 延迟厉害.





* 有缓冲channel原理

![](F:\陈晓会\GO\极客大学\Go-000\Week03\buffered_channel.png)

* 基本无延迟, reduce blocking latency

> Send 不晚于 Receive 发生

* 不保证数据到达, 越大的buffer, 越小的保障到达. buffer = 1时,给你延迟一个消息的保证.





* 保障率: 1/buffered_size  * 100% (无缓冲时,100%)



* Tming out  超时
* Moving on 放弃数据 select 中的 defalut
* Pipeling  ->ch1 -> ch2 -> ch3 ->ch4 ->ch5
* Fan-out 扇出   -> ch->goroutine1, goroutine2, goroutine3,... ...; Fan-in 扇入   ch1, ch2, ch3, ch4,... ... ->goroutine
* Cancellation
  * close先于receive发生
  * 只有发送者知道什么时候close.
  * 不需要传递数据,或nil
  * 适合去做超时控制
* Context





* 设计哲学

  * If any given Send on a channel CAN cause the sending goroutine to bloack:
  * If any given Send on a cha nnel WON'T cause the sending goroutine to block:
  * Less is more with buffers.

* master-worker

* kafka: 一个partition一个进程,来消费.

  > 程序运行遇到瓶颈:
  >
  > 1. cpu是否达到100%.
  > 2. 程序代码有问题,串行并用/大锁.



