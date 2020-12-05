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

* 

