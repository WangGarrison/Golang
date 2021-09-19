# Go并发

Go 从语言层面就支持并发。同时实现了自动垃圾回收机制，和其他编程语言相比更加轻量。

Go 语言通过编译器运行时（runtime），从语言上支持了并发的特性。Go 语言的并发通过 **goroutine** 特性完成。goroutine 类似于线程，可以根据需要创建多个 goroutine 并发工作。goroutine 是由 Go 语言的运行时调度完成，而线程是由操作系统调度完成。

Go 语言还提供 **channel** 在多个 goroutine 间进行通信。goroutine 和 channel 是 Go 语言秉承的 CSP（Communicating Sequential Process）并发模式的重要实现基础

- ==goroutine：轻量级线程/协程==
- ==channel：goroutine间通信，（引用类型）==

<font color="blue">**Go 程序从 main 包的 main() 函数开始，在程序启动时，Go 程序就会为 main() 函数创建一个默认的 goroutine，所有 goroutine 在 main() 函数结束时会一同结束**</font>

**进程/线程**

进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。

线程是进程的一个执行实体，是 CPU 调度和分派的基本单位，它是比进程更小的能独立运行的基本单位。

一个进程可以创建和撤销多个线程，同一个进程中的多个线程之间可以并发执行。

**并发/并行**

多线程程序在单核心的 cpu 上运行，称为并发；多线程程序在多核心的 cpu 上运行，称为并行。

并发与并行并不相同，并发主要由切换时间片来实现“同时”运行，并行则是直接利用多核实现多线程的运行，Go程序可以设置使用核心数，以发挥多核计算机的能力。

**线程/协程**

线程：一个线程上可以跑多个协程，协程是轻量级的线程。

协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。

# goroutine

goroutine是协程的go语言实现，相当于把别的语言的类库的功能内置到语言里。从调度上看，goroutine的调度开销远远小于线程调度开销

```go
//go 关键字放在方法调用前新建一个 goroutine 并执行方法体
go GetThingDone(param1, param2);

//新建一个匿名方法并执行
go func(param1, param2) {
}(val1, val2)

//直接新建一个 goroutine 并在 goroutine 中执行代码块
go {
    //do someting...
}

//注意：goroutine 在多核 cpu 环境下是并行的
```

**普通函数创建协程**

一个函数可以被创建多个 goroutine，一个 goroutine 必定对应一个函数。

使用 go 关键字就可以创建 goroutine，将 go 声明放到一个需调用的函数之前，在相同地址空间调用运行这个函数，这样该函数执行时便会作为一个独立的并发线程，这种线程在Go语言中则被称为 goroutine，代码示例如下：

```go
go 函数名( 参数列表 )  //使用 go 关键字创建 goroutine 时，被调用函数的返回值会被忽略，若要返回数据，使用channel
```

实例：

```go
package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

func main() {
	go running()  //创建一个协程，与main协程并发运行

	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
}
```

上述代码示意图如下：

<img align='left' src="img/Go6%EF%BC%9A%E5%B9%B6%E5%8F%91.img/image-20210710094947865.png" alt="image-20210710094947865" style="zoom:40%;" />

**匿名函数创建协程**

```go
//使用匿名函数或闭包创建 goroutine 时，除了将函数定义部分写在 go 的后面之外，还需要加上匿名函数的调用参数
go func( 参数列表 ){
    函数体
}( 调用参数列表 )
```

实例：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}() //记得加调用的参数，这里是空参

	var input string
	fmt.Scanln(&input)
}
```

**代码块创建协程**

```go
//直接新建一个 goroutine 并在 goroutine 中执行代码块
go {
    //do someting...
}
```

# 协程间通信

在工程上，有两种最常见的并发通信模型：

- <font color="red">**共享数据**</font>：共享数据是指多个并发单元分别保持对同一个数据的引用，实现对该数据的共享。被共享的数据可能有多种形式，比如内存数据块、磁盘文件、网络数据等。在实际工程应用中最常见的无疑是内存了，也就是常说的共享内存
- <font color="red">**消息机制**</font>：消息机制认为每个并发单元是自包含的、独立的个体，并且都有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。这有点类似于进程的概念，每个进程不会被其他进程打扰，它只做好自己的工作就可以了。不同进程间靠消息来通信，它们不会共享内存。

**共享数据的缺点：**

如下代码，通过共享全局变量来达到两个协程通信：

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

func Count(mtx *sync.Mutex) {
	mtx.Lock()
    
	counter++
	fmt.Println(counter)
    
	mtx.Unlock()
}

func main() {
	mtx := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(mtx) //创建10个协程，共享一把锁与counter
	}

	for {
		mtx.Lock()
		c := counter
		mtx.Unlock()
		runtime.Gosched() //等待其他协程执行完毕，主协程再继续向下执行。
		if c >= 10 {
			break
		}
	}
}

```

实现一个如此简单的功能，却写出如此臃肿而且难以理解的代码。想象一下，在一个大的系统中具有无数的锁、无数的共享变量、无数的业务逻辑与错误处理分支，那将是一场噩梦。==这噩梦就是众多 C/C++开发者正在经历的==

那怎么办？答：使用==消息机制——channel==来作为协程间通信方式

# channel

Go语言提倡使用通信的方法代替共享内存来完成协程间通信，当一个资源需要在 goroutine 之间共享时，通道在 goroutine 之间架起了一个管道，并提供了确保同步交换数据的机制。

channel是协程间通信，是进程内部的通信，如果需要跨进程通信，建议用分布式系统的方法来解决，比如使用 Socket 或者 HTTP 等通信协议

channel 是类型相关的，也就是说，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定（类似于Unix中的类型安全的管道）	

声明通道：

```go
var name chan 类型  //chan 类型的空值是 nil，声明后需要配合 make 后才能使用
```

channel创建示例代码：(需使用 make 创建 channel)

```go
ch1 := make(chan int)                 // 创建一个整型类型的通道
ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
type Equip struct{ /* 一些字段 */ }
ch2 := make(chan *Equip)             // 创建Equip指针类型的通道, 可以存放*Equip
```

**通道的特性**

- 在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据
- 通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序

**使用通道发送数据**

==使用<-向通道发送数据==

```go
ch1 <- 值  //将值发送到ch1通道
```

```go
// 创建一个空接口通道
ch := make(chan interface{})

// 将0放入通道中
ch <- 0

// 将hello字符串放入通道中
ch <- "hello"
```

<font color='red'>**把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞**</font>（Go 程序运行时能智能地发现一些永远无法发送成功的语句并做出提示），如下示例：

```go
package main

func main() {
	ch := make(chan int)
	ch <- 0
}

运行报错：
fatal error: all goroutines are asleep - deadlock!
运行时发现所有的 goroutine（包括main）都处于等待 goroutine。也就是说所有 goroutine 中的 channel 并没有形成发送和接收对应的代码
```

**使用通道接收数据**

==通道接收同样使用`<-`操作符==

阻塞接收数据：阻塞模式接收数据时，将接收变量作为`<-`操作符的左值

```go
data := <-ch  //执行该语句时将会阻塞，直到接收到数据并赋值给 data 变量
```

非阻塞接收数据：使用非阻塞方式从通道接收数据时，语句不会发生阻塞

```go
data,ok := <-ch  //ok：表示是否接收到数据，非阻塞的通道接收方法可能造成高的CPU占用，因此使用非常少
```

接收任意数据，忽略从通道返回的数据

```go
<-ch  //执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步
```

循环接收：通道的数据接收可以借用 for range 语句进行多个元素的接收操作

```go
for data := range ch {  //通道 ch 是可以进行遍历的，遍历的结果就是接收到的数据
}
```

**通道收发数据示例：**

```go
package main
import (
    "fmt"
    "time"
)
func main() {
    // 构建一个通道
    ch := make(chan int)
    
    // 开启一个并发匿名函数
    go func() {
        // 从3循环到0
        for i := 3; i >= 0; i-- {
            // 发送3到0之间的数值
            ch <- i
            // 每次发送完时等待
            time.Sleep(time.Second)
        }
    }()
    
    // 遍历接收通道数据
    for data := range ch {
        // 打印通道数据
        fmt.Println(data)
        // 当遇到数据0时, 退出接收循环
        if data == 0 {
                break
        }
    }
}
```

**通道使用注意事项：**

- 通道的收发操作在不同的两个 goroutine 间进行：由于通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行
- 接收将持续阻塞直到发送方发送数据：如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止
- 每次接收一个元素。通道一次只能接收一个数据元素

# channel示例：并发打印

```go
package main

import "fmt"

func printer(c chan int) {
	for {
		data := <-c    //从chan中获取一个数据
		if data == 0 { //0视为数据结束
			break
		}
		fmt.Println(data)
	}
	c <- 0 //通知main结束循坏
}

func main() {
	c := make(chan int)
	
    //启动printer协程，阻塞等待获取数据（消费者）
	go printer(c)

	//主协程往chan中发送数据（生产者）
	for i := 1; i <= 10; i++ {
		c <- i
	}
	c <- 0 //通知printer协程结束

	<-c //等待printer结束
}
```

# 单向通道

我们在将一个 channel 变量传递到一个函数时，可以通过将其指定为单向 channel 变量，从而限制该函数中可以对此 channel 的操作，比如只能往这个 channel 中写入数据，或者只能从这个 channel 读取数据。即传递单方向的 channel 类型

只能写入数据的通道类型为`chan<-`，只能读取数据的通道类型为`<-chan`，格式如下：

```go
var 通道实例 chan<- 元素类型    // 只能写入数据的通道
var 通道实例 <-chan 元素类型    // 只能读取数据的通道
```

代码示例：

```go
//创建一个管道
ch := make(chan int)

// 声明一个只能写入数据的通道类型, 并赋值为ch
var chSendOnly chan<- int = ch
//声明一个只能读取数据的通道类型, 并赋值为ch
var chRecvOnly <-chan int = ch
```

使用 make 创建通道时，也可以创建一个只写入或只读取的通道：

```go
ch := make(<-chan int)
var chReadOnly <-chan int = ch
```

单向通道有利于代码的严谨性，比如限制某个通道只能读取数据，这样就避免了对数据的修改，例如time包中的单向通道：time 包中的计时器会返回一个 timer 实例，代码如下：

```go
timer := time.NewTimer(time.Second)
```

timer的Timer类型定义如下：

```go
type Timer struct {    
    C <-chan Time    
    r runtimeTimer
}
```

如果此处不进行通道方向约束，一旦外部向通道写入数据，将会造成其他使用到计时器的地方逻辑产生混乱。

因此，单向通道有利于代码接口的严谨性。

**关闭channel**

```go
close(ch)
```

如何判断一个 channel 是否已经被关闭？

```go
x, ok := <-ch  //如果ok值是 false 则表示 ch 已经被关闭
```

# 资源竞争-线程不安全问题

如下代码，两个协程共享的是全局变量count，并且没有加锁：

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()  //让当前 goroutine 暂停的意思，退回执行队列，让其他等待的 goroutine 运行，放在这目的是为了使资源竞争的结果更明显。
		value++
		count = value
	}
}
```

多运行几次，会发现结果可能是 2，也可以是 3，还可能是 4。这是因为 count 变量没有任何同步保护，所以两个 goroutine 都会对其进行读写，会导致对已经计算好的结果被覆盖，以至于产生错误结果。

两个 goroutine 分别假设为 g1 和 g2：

- g1 读取到 count 的值为 0；
- 然后 g1 暂停了，切换到 g2 运行，g2 读取到 count 的值也为 0；
- g2 暂停，切换到 g1，g1 对 count+1，count 的值变为 1；
- g1 暂停，切换到 g2，g2 刚刚已经获取到值 0，对其 +1，最后赋值给 count，其结果还是 1；
- 可以看出 g1 对 count+1 的结果被 g2 给覆盖了，两个 goroutine 都 +1 而结果还是 1。

之所以出现上面的问题，是因为+1操作并不是原子的，导致线程不安全

并发环境中，对于同一个资源的读写必须是原子化的，也就是说，同一时间只能允许有一个 goroutine 对共享资源进行读写操作

**Go 为我们提供了一个工具帮助我们检查线程安全问题：`go build -race`**

在项目目录下执行这个命令，生成一个可以执行文件，然后再运行这个可执行文件，就可以看到打印出的检测信息

![image-20210710104532675](img/Go6%EF%BC%9A%E5%B9%B6%E5%8F%91.img/image-20210710104532675.png)

![image-20210710104454679](img/Go6%EF%BC%9A%E5%B9%B6%E5%8F%91.img/image-20210710104454679.png)

通过运行结果可以看出 goroutine 8 在代码 25 行读取共享资源`value := count`，而这时 goroutine 7 在代码 28 行修改共享资源`count = value`，而这两个 goroutine 都是从 main 函数的 16、17 行通过 go 关键字启动的。

# 锁住共享资源

atomic 和 sync 包里的一些函数就可以对共享的资源进行加锁操作

**原子函数**

原子函数能够以很底层的加锁机制来同步访问整型变量和指针，示例代码如下所示：

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println(counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) //安全原子的对counter加1,这个函数会同步整型值的加法，方法是强制同一时刻只能有一个 gorountie 运行并完成这个加法操作
		runtime.Gosched()
	}
}
```

**互斥锁**

互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界代码。

```go
package main
import (
    "fmt"
    "runtime"
    "sync"
)

var (
    counter int64
    wg      sync.WaitGroup
    mutex   sync.Mutex
)

func main() {
    wg.Add(2)
    go incCounter(1)
    go incCounter(2)
    wg.Wait()
    fmt.Println(counter)
}

func incCounter(id int) {
    defer wg.Done()
    for count := 0; count < 2; count++ {
        //同一时刻只允许一个goroutine进入这个临界区
        mutex.Lock()
        {
            value := counter
            runtime.Gosched()
            value++
            counter = value
        }
        mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
    }
}
```

# WaitGroup的使用

sync.WaitGroup可以解决同步阻塞等待的问题。一个人等待一堆人干完活的问题得到优雅解决。

等待组内部拥有一个计数器，计数器的值可以通过方法调用实现计数器的增加和减少。当我们添加了 N 个并发任务进行工作时，就将等待组的计数器值增加 N。每个任务完成时，这个值减 1。同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时，表示所有任务已经完成。

`WaitGroup` 对象内部有一个计数器，最初从0开始，它有三个方法：`Add(), Done(), Wait()` 用来控制计数器的数量。`Add(n)` 把计数器设置为`n` ，`Done()` 每次把计数器`-1` ，`wait()` 会阻塞代码的运行，直到计数器地值减为`0`。

- Add()：计数器+
- Done()：计数器-1
- wait()：阻塞代码的运行，直到计数器地值减为`0`

```go
func main() {
    wg := sync.WaitGroup{}
    wg.Add(100)  //计数器+100
    for i := 0; i < 100; i++ {
        go func(i int) {
            fmt.Println(i)
            wg.Done()  //一个干完活，计数器减一
        }(i)
    }
    wg.Wait()  //阻塞代码的运行，直到计数器地值减为0，即等待所有协程执行完毕
}
```

**wg.Done()经常与defer配合使用：**在当前协程执行结束时，执行计数-1

```go
defer wg.Done()
```

# 设置GOMAXPROCS

传统逻辑中，开发者需要维护线程池中线程与 CPU 核心数量的对应关系。同样的，Go 中也可以通过 runtime.GOMAXPROCS() 函数做到，格式为：

```go
runtime.GOMAXPROCS(逻辑CPU数量)
```

这里的逻辑CPU数量可以有如下几种数值：

- <1：不修改任何数值。
- =1：单核心执行。
- \>1：多核并发执行。

一般情况下，可以使用 runtime.NumCPU() 查询 CPU 数量，并使用 runtime.GOMAXPROCS() 函数进行设置，例如：

```go
runtime.GOMAXPROCS(runtime.NumCPU())
```

执行上面语句以便让代码并发执行，最大效率地利用 CPU。

GOMAXPROCS 同时也是一个环境变量，在应用程序启动前设置环境变量也可以起到相同的作用。

注意：不推荐盲目修改语言运行时对逻辑处理器的默认设置。如果真的认为修改逻辑处理器的数量可以改进性能，也可以对语言运行时的参数进行细微调整。

# 无缓冲的管道

```go
ch1 := make(chan int)
```

==以最简单方式调用make函数创建的是一个无缓存的channel，但是我们也可以指定第二个整型参数，对应channel的容量==。如果channel的容量大于零，那么该channel就是带缓存的channel。

Go语言中无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。这种类型的通道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作。

如果两个 goroutine 没有同时准备好，通道会导致先执行发送或接收操作的 goroutine 阻塞等待。这种对通道进行发送和接收的交互行为本身就是同步的。其中任意一个操作都无法离开另一个操作单独存在

- 阻塞指的是由于某种原因数据没有到达，当前协程（线程）持续处于等待状态，直到条件满足才解除阻塞

- 同步指的是在两个或多个协程（线程）之间，保持数据内容一致性的机制。

无缓冲通道可以看作是长度永远为 0 的带缓冲通道

# 带缓冲的通道

Go语言中有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。这种类型的通道并不强制要求 goroutine 之间必须同时完成发送和接收。通道会阻塞发送和接收动作的条件也会不同。

阻塞条件：只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞。

因此：无缓冲的通道保证进行发送和接收的 goroutine 会在同一时间进行数据交换；有缓冲的通道没有这种保证。

> 无缓冲通道保证收发过程同步。无缓冲收发过程类似于快递员给你电话让你下楼取快递，整个递交快递的过程是同步发生的，你和快递员不见不散。但这样做快递员就必须等待所有人下楼完成操作后才能完成所有投递工作。如果快递员将快递放入快递柜中，并通知用户来取，快递员和用户就成了异步收发过程，效率可以有明显的提升。带缓冲的通道就是这样的一个“快递柜”。

**创建带缓冲的通道**

```go
ch1 := make(chan 类型, 缓冲大小)  //缓冲大小：决定通道最多可以保存的元素数量。
```

代码示例：

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	fmt.Println(len(ch))

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(len(ch))
}
```

**为什么Go语言对通道要限制长度而不提供无限长度的通道？**

我们知道通道（channel）是在两个 goroutine 间通信的桥梁。使用 goroutine 的代码必然有一方提供数据，一方消费数据。当提供数据一方的数据供给速度大于消费方的数据处理速度时，如果通道不限制长度，那么内存将不断膨胀直到应用崩溃。因此，限制通道的长度有利于约束数据提供方的供给速度，供给数据量必须在消费方处理量+通道长度的范围内，才能正常地处理数据。

# select实现超时机制

Go语言没有提供直接的超时处理机制，所谓超时可以理解为当我们上网浏览一些网站时，如果一段时间之后不作操作，就需要重新登录。	

我们可以使用 select 来设置超时

select特点：只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况

select 的用法与 switch 语言非常类似，由 select 开始一个新的选择块，每个选择条件由 case 语句来描述。

与 switch 语句相比，select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作，大致的结构如下：

```go
select{
    case 操作1:
        响应操作1
    case 操作2:
        响应操作2
    …
    default:
        没有操作情况
}
```

```go
select {
    case <-chan1:
    // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
    // 如果成功向chan2写入数据，则进行该case处理语句
    default:
    // 如果上面都没有成功，则进入default处理流程
}
```

在一个 select 语句中，Go语言会按顺序从头至尾评估每一个发送和接收的语句。

如果其中的任意一语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。

如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有如下两种可能的情况：

- 如果给出了 default 语句，那么就会执行 default 语句，同时程序的执行会从 select 语句后的语句中恢复；
- 如果没有 default 语句，那么 select 语句将被阻塞，直到至少有一个通信可以进行下去。

**select实现超时示例代码：**

一个协程死循坏从通道中读取数据，超过3秒没读到则结束

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for { //死循坏从ch通道中读取数据，超过3秒阻塞则结束
			select {
			case num := <-ch:
				fmt.Println("num=", num)
			case <-time.After(3 * time.Second): //设置超时时间为3s
				fmt.Println("超时")
				quit <- true
			}
		}
	}() //记得参数

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("程序结束")
}
```

<img align='left' src="img/Go6%EF%BC%9A%E5%B9%B6%E5%8F%91.img/image-20210710171907292.png" alt="image-20210710171907292" style="zoom:50%;" />



# go模拟远程过程调用

使用通道代替 Socket 实现 RPC 的过程。客户端与服务器运行在同一个进程的两个goroutine

```go
package main

import (
	"errors"
	"fmt"
	"time"
)

// 模拟RPC客户端的请求和接收消息封装
func RPCClient(ch chan string, req string) (string, error) {
	// 向服务器发送请求
	ch <- req

	// 等待服务器返回
	select {
	case ack := <-ch: // 接收到服务器返回的数据
		return ack, nil
	case <-time.After(time.Second): // 超时
		return "", errors.New("Time out")
	}
}

// 模拟RPC服务器端接收客户端请求和回应
func RPCServer(ch chan string) {
	for { // 接收客户端的请求
		data := <-ch

		fmt.Println("server received:", data)

		ch <- "roger" //返回给客户端
	}
}

func main() {
	ch := make(chan string)

	//启动服务器
	go RPCServer(ch)

	//启动客户端
	recv, err := RPCClient(ch, "hi")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received", recv)
	}
}
```

# 超时事件

Go语言中的 time 包提供了计时器的封装。由于 Go语言中的通道和 goroutine 的设计，定时任务可以在 goroutine 中通过同步的方式完成，也可以通过在 goroutine 中异步回调完成。

**超时回调**

time.AfterFunc() 函数是在 time.After 基础上增加了到时的回调，方便使用

```go
 // 过1秒后, 调用匿名函数
time.AfterFunc(time.Second, func() {
    // 1秒后, 打印结果
    fmt.Println("one second after")
    // 通知main()的goroutine已经结束
    exit <- 0
})
```

**定时打点**

```go
package main
import (
    "fmt"
    "time"
)
func main() {
    // 创建一个打点器, 每500毫秒触发一次
    ticker := time.NewTicker(time.Millisecond * 500)
    // 创建一个计时器, 2秒后触发
    stopper := time.NewTimer(time.Second * 2)
    // 声明计数变量
    var i int
    // 不断地检查通道情况
    for {
        // 多路复用通道
        select {
        case <-stopper.C:  // 计时器到时了
            fmt.Println("stop")
            // 跳出循环
            goto StopHere
        case <-ticker.C:  // 打点器触发了
            // 记录触发了多少次
            i++
            fmt.Println("tick", i)
        }
    }
// 退出的标签, 使用goto跳转
StopHere:
    fmt.Println("done")
}
```

# 互斥锁和读写锁

Go语言包中的 sync 包提供了两种锁类型：

- sync.Mutex：互斥锁，一个上锁，其他的只能等待
- sync.RWMutex：读写锁，单写多读模型。读锁占用的情况下，会阻止写，但不阻止读；写锁占用，读写都会组织

> 读写锁其实是通过组合互斥锁来实现的，RWMutex结构体如下：
>
> ```go
> type RWMutex struct {
>     w Mutex
>     writerSem uint32
>     readerSem uint32
>     readerCount int32
>     readerWait int32
> }
> ```

互斥锁的使用示例：

```go
// 定义互斥锁
var countGuard sync.Mutex

// 锁定
countGuard.Lock()

// 解除锁定
countGuard.Unlock()  // 前面可以加defer，在函数退出时解除锁定
```

读写锁的使用示例：

```go
// 与变量对应的使用互斥锁
var countGuard sync.RWMutex

// 读加锁
countGuard.RLock()

// 读解锁
defer countGuard.RUnlock()
```

# 死锁、活锁、饥饿

**死锁**：死锁是指两个或两个以上的进程（或线程）在执行过程中，因争夺资源而造成的一种互相等待的现象，若无外力作用，它们都将无法推进下去。

死锁的四个必要条件：

- 互斥
- 请求和保持

- 不可剥夺
- 环路等待

死锁解决办法：

- 如果并发查询多个表，约定访问顺序；
- 在同一个事务中，尽可能做到一次锁定获取所需要的资源；
- 对于容易产生死锁的业务场景，尝试升级锁颗粒度，使用表级锁；
- 采用分布式事务锁或者使用乐观锁。

**活锁**：当多个相互协作的线程都对彼此进行相应而修改自己的状态，并使得任何一个线程都无法继续执行时，就导致了活锁。这就像两个过于礼貌的人在路上相遇，他们彼此让路，然后在另一条路上相遇，然后他们就一直这样避让下去。

例如：线程 1 可以使用资源，但它很礼貌，让其他线程先使用资源，线程 2 也可以使用资源，但它同样很绅士，也让其他线程先使用资源。就这样你让我，我让你，最后两个线程都无法使用资源。

活锁不会阻塞线程，但也不能继续执行，因为线程将不断重复同样的操作，而且总会失败。

活锁通常发生在处理事务消息中，如果不能成功处理某个消息，那么消息处理机制将回滚事务，并将它重新放到队列的开头。这样，错误的事务被一直回滚重复执行，这种形式的活锁通常是由过度的错误恢复代码造成的，因为它错误地将不可修复的错误认为是可修复的错误。

**解决活锁：引入随机性**。在重试机制中引入随机性。例如在网络上发送数据包，如果检测到冲突，都要停止并在一段时间后重发，如果都在 1 秒后重发，还是会冲突，所以引入随机性可以解决该类问题。

**饥饿**：是指一个可运行的进程尽管能继续执行，但被调度器无限期地忽视，而不能被调度执行的情况。

# Go语言CSP：通信顺序进程

 CSP（communicating sequential processes）：通信顺序进程，即两个独立的并发实体通过共享 channel（管道）进行通信的并发模型

- process：对应go中的goroutine
- communicating sequential：对应go中的通过chan完成协程间通信