Go语言中，string、int、bool、数组、stuct都属于非引用数据类型。

Go语言中，==指针、Slice切片、map、chan都是引用数据类型==，引用的时候也是类似指针地址

# 数组

```go
nums := []int{0, 1, 2, 3, 7, 8, 9, 10}
var arr [10]int  //C: int arr[10]; go就是把类型放后面了
len(arr)  //内置函数 len() 可以返回数组中元素的个数
```

数组的长度是数组类型的一个组成部分，因此 [3]int 和 [4]int 是两种不同的数组类型，数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定

```go
var a [3]int
fmt.Println(a[0])  //打印第一个元素
fmt.Println(a[len(a)-1])  //打印最后一个元素

//打印索引和元素
for i, v := range a {  //基于范围的for循环：for key,value := range oldmap { newmap[key]=value }
    fmt.Printf("%d %d\n", i, v) 
}
```

默认情况下，数组的每个元素都会被初始化为元素类型对应的零值，对于数字类型来说就是 0，同时也可以使用数组字面值语法，用一组值来初始化数组：

```go
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) // "0"
```

在数组定义中，如果在数组长度的位置出现...省略号，则表示数组的长度是根据初始化值的个数来计算，如下：

```go
q := [...]int{1,2,3}
```

**数组比较是否相等**

如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（`==`和`!=`）来判断两个数组是否相等，只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。

**遍历数组**

```go
var arr [10]int
for k, v := range arr {
    fmt.Println(k, v)
}
```

**多维数组**

```go
var arr [4][2]int  //C++: int arr[4][2]
arr = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

array = [4][2]int{1: {20, 21}, 3: {40, 41}}  // 声明并初始化数组中索引为 1 和 3 的元素
array = [4][2]int{1: {0: 20}, 3: {1: 41}}  // 声明并初始化数组中指定的元素
```

```go
// 声明一个 2×2 的二维整型数组
var array [2][2]int
// 设置每个元素的整型值
array[0][0] = 10
array[0][1] = 20
array[1][0] = 30
array[1][1] = 40
```

# 切片slice

“动态数组”

Go语言中切片的内部结构包含地址、大小和容量，切片一般用于快速地操作一块数据集合，如果将数据集合比作切糕的话，切片就是你要的“那一块”，切的过程包含从哪里开始（切片的起始位置）及切多大（切片的大小），容量可以理解为装切片的口袋大小，如下图所示。



<img align='left' src="img/Golang%E5%AE%B9%E5%99%A8.img/1-1PQ3154340Y9.jpg" alt="img" style="zoom:70%;" />





切片默认指向一段连续内存区域，可以是数组，也可以是切片本身。从连续内存区域生成切片是常见的操作，格式如下：

```go
切的对象[起始索引:结束索引]
arr[开始位置:结束位置]
slice[开始位置:结束位置]
```

从数组或切片生成新的切片拥有如下特性：

```go
取出的元素数量为：结束位置 - 开始位置；
取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
当缺省开始位置时，表示从连续区域开头到结束位置；
当缺省结束位置时，表示从开始位置到整个连续区域末尾；
两者同时缺省时，与切片本身等效；
两者同时为 0 时，等效于空切片，一般用于切片复位。
```

```go
var arr [3]int = [3]int{0, 1, 2}
fmt.Print(arr, arr[0:1])
输出：[0 1 2] [0]
```

根据索引位置取切片 slice 元素值时，取值范围是（0～len(slice)-1），超界会报运行时错误，生成切片时，结束位置可以填写 len(slice) 但不会报错。

```go
a := []int{1, 2, 3}
fmt.Println(a[:])
输出：[1 2 3]

a := []int{1, 2, 3}
fmt.Println(a[0:0])
输出：[]
```

**直接声明新的切片**

除了可以从原有的数组或者切片中生成切片外，也可以声明一个新的切片，每一种类型都可以拥有其切片类型，表示多个相同类型元素的连续集合，因此切片类型也可以被声明，切片类型声明格式如下：

```go
// 声明字符串切片
var strList []string
// 声明整型切片
var numList []int
// 声明一个空切片
var numListEmpty = []int{}
```

要注意的是slice类型的变量s和数组类型的变量a的初始化语法的差异。slice和数组的字面值语法很类似，它们都是用花括弧包含一系列的初始化元素，但是对于==slice并没有指明序列的长度==。这会隐式地创建一个合适大小的数组，然后slice的指针指向底层的数组

**make()函数构造切片**

内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分可以省略，在这种情况下，容量将等于长度。

```Go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```

在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量。在第一种语句中，slice是整个数组的view。在第二个语句中，slice只引用了底层数组的前len个元素，但是容量将包含整个的数组。额外的元素是留给未来的增长用的。

```go
make([]int, size, cap)
/*其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。*/
```

```go
a := make([]int, 2)
b := make([]int, 2, 10)
fmt.Println(a, b)
fmt.Println(len(a), len(b))

输出：
[0 0] [0 0]
2 2
其中 a 和 b 均是预分配 2 个元素的切片，只是 b 的内部存储空间已经分配了 10 个，但实际使用了 2 个元素。
容量不会影响当前的元素个数，因此 a 和 b 取 len 都是 2。
```

注意：使用 make() 函数生成的切片一定发生了内存分配操作，但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作。

**append()为切片动态添加元素**

append会返回新切片

```go
append(a, val)  // 给第一个参数后面追加第二个参数
```

```go
var a []int
a = append(a, 1) // 给切片a追加1个元素
a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
```

在使用 append() 函数为切片动态添加元素时，如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……

```go
var arr []int
for i := 0; i < 10; i++ {
    arr = append(arr, i)
    fmt.Printf("i:%d, len:%d, cap:%d\n", i, len(arr), cap(arr))
}
```

输出如下：

![image-20210618174629650](img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618174629650.png)

在切片开头添加元素：(相当于把切片追加到元素后面)

```go
var a = []int{0, 1, 2, 3, 4}
a = append([]int{-1}, a...)
fmt.Print(a)

输出：[-1 0 1 2 3 4]
```

append的链式操作：

```go
var a []int
a = append(a[:i], append([]int{x}, a[i:]...)...) // 相当于在第i个位置插入x
/*
先把a[i:]追加到元素x后面，再把该新切片追加到a[:i]后面
*/
```

<img align='left' src="img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618180413937.png" alt="image-20210618180413937" style="zoom:67%;" />

```go
a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 相当于在第i个位置插入切片
/*
先把a[i:]追加到{1,2,3}后面，再把该新切片追加到a[:i]后面
*/
每个添加操作中的第二个 append 调用都会创建一个临时切片，并将 a[i:] 的内容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:i] 中
```

**copy切片的复制**

```go
copy( destSlice, srcSlice []T) int  // 将 srcSlice 复制到 destSlice，copy() 函数的返回值表示实际发生复制的元素个数。
```

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{5, 4, 3}
copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
```

**删除切片元素**

Go语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素，根据要删除元素的位置有三种情况，分别是从开头位置删除、从中间位置删除和从尾部删除，其中删除切片尾部的元素速度最快。

删除开头元素可以直接移动指针：

```go
a = []int{1, 2, 3}
a = a[1:] // 删除开头1个元素
a = a[N:] // 删除开头N个元素
```

也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）

```go
a = []int{1, 2, 3}
a = append(a[:0], a[1:]...) // 删除开头1个元素
a = append(a[:0], a[N:]...) // 删除开头N个元素
```

还可以用copy函数来删除开头的元素：

```go
a = []int{1, 2, 3}
a = a[:copy(a, a[1:])] // 删除开头1个元素
a = a[:copy(a, a[N:])] // 删除开头N个元素
```

删除中间位置元素：

```go
a = []int{1, 2, 3, ...}
a = append(a[:i], a[i+1:]...) // 删除中间1个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素
a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素
```

删除尾部元素：

```go
a = []int{1, 2, 3}
a = a[:len(a)-1] // 删除尾部1个元素
a = a[:len(a)-N] // 删除尾部N个元素
```

Go语言中删除切片元素的本质是，==以被删除元素为分界点，将前后两个部分的内存重新连接起来==

**多维切片**

```go
var slice [][]int  //声明一个二维切片
slice = [][]int{{10}, {100, 200}}  //为二维切片赋值

//简写如下：
slice := [][]int{{10}, {100,200}}
```

该二维切片示意图如下：通过下图可以看到外层的切片包括两个元素，每个元素都是一个切片，第一个元素中的切片使用单个整数 10 来初始化，第二个元素中的切片包括两个整数，即 100 和 200。

![image-20210618182549242](img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618182549242.png)

**切片判空**

随然切片可以与nil是否比较，支持s == nil，但是判空使用==len(s) == 0==来判断





# map

map：映射/哈希表/关联数组/字典

基本用法：

```go
//错误写法：
var p0 map[int]int
p0[0] = 1          //报错，map是引用类型，未经初始化则是nil，不能直接将值赋给nil
fmt.Println(p0[0]) //但是可以打印nil值，ok，int类型nil值打印0，字符串类型打印空字符

//正确写法1
var p1 = map[int]int{0: 1}
p1[0] = 1
fmt.Println(p1[0])

//正确写法2
pp := map[int]int{}
pp[0] = 1
fmt.Println(pp[0])

//正确写法3
p3 := make(map[int]int)
p3[0] = 1
fmt.Println(p3[0])
```

```go
var mapname map[keytype]valuetype  // [keytype] 和 valuetype 之间允许有空格。
var mapLit map[string]int
```

**遍历map**

```go
var mapLit map[string]int = map[string]int{"one": 1, "two": 2}
for k, v := range mapLit {
    fmt.Print(k, v)
}
```

注意：将一个map赋值给另一个map是浅复制，新map是旧map的引用

```go
var map1 map[string]int = map[string]int{"one": 1, "two": 2}
for k, v := range map1 {
    fmt.Println(k, v)
}

var map2 = map1
map1["one"] = 10

for k, v := range map2 {
    fmt.Println(k, v)
}
```

![image-20210618185132735](img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618185132735.png)

指定map容量：

```go
map2 := make(map[string]float, 100)  //100是cap
```

**切片作为map的value值**

既然一个 key 只能对应一个 value，而 value 又是一个原始类型，那么如果一个 key 要对应多个值怎么办？例如，当我们要处理 unix 机器上的所有进程，以父进程（pid 为整形）作为 key，所有的子进程（以所有子进程的 pid 组成的切片）作为 value。通过将 value 定义为 []int 类型或者其他类型的切片，就可以优雅的解决这个问题，示例代码如下所示：

```go
mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)
```

**遍历map**

注意：遍历输出元素的顺序与填充顺序无关，不能期望 map 在遍历时返回某种期望顺序的结果。

```go
for k, v := range map2 {
    fmt.Println(k, v)
}
```

遍历对于Go语言的很多对象来说都是差不多的，直接使用 for range 语法即可，遍历时，可以同时获得键和值，如只遍历值，可以使用下面的形式, 将不需要的键使用`_`改为匿名变量形式：

```go
for _, v := range map1 {}
```

只遍历键时，使用下面的形式：(无须将值改为匿名变量形式，忽略值即可。)

```go
for k := range map1 {}
```

**map元素的删除和清空**

删除map1中的元素：

```go
delete(map1, key)
```

清空map中所有的元素：有意思的是，Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。

**并发安全的sync.Map**

map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。

需要并发读写时，一般的做法是加锁，但这样性能并不高，Go语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map，sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构类型。

sync.Map 有以下特性：

- 无须初始化，直接声明即可。
- sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
- 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map

	//set key-value
	scene.Store("one", 1)
	scene.Store("two", 2)
	scene.Store("Three", 3)

	//get
	fmt.Println(scene.Load("one"))

	//delete
	scene.Delete("two")

	//遍历，Range() 方法可以遍历 sync.Map，遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate: ", k, v)
		return true
	})
}
```

sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。

# list

列表是一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表、双链表等。

在Go语言中，列表使用 container/list 包来实现，内部的实现原理是==双链表==，列表能够高效地进行任意位置的元素插入和删除操作。

**初始化列表**

```go
//通过 container/list 包的 New() 函数初始化 list
变量名 := list.New()

//通过 var 关键字声明初始化 list
var 变量名 list.List
```

**插入元素**

列表插入元素的方法如下表所示。

| 方  法                                                | 功  能                                            |
| ----------------------------------------------------- | ------------------------------------------------- |
| PushFront("hh")                                       | 在列表前面插入                                    |
| PushBack(88)                                          | 在列表末尾插入                                    |
| InsertAfter(v interface {}, mark * Element) * Element | 在 mark 点之后插入元素，mark 点由其他插入函数提供 |
| InsertBefore(v interface {}, mark * Element) *Element | 在 mark 点之前插入元素，mark 点由其他插入函数提供 |
| PushBackList(other *List)                             | 添加 other 列表元素到尾部                         |
| PushFrontList(other *List)                            | 添加 other 列表元素到头部                         |

```go
package main

import (
	"container/list"
)

func main() {
	li := list.New()
	li.PushBack("first")
	li.PushFront(67)
}
```

**删除元素**

列表插入函数的返回值会提供一个 *list.Element 结构，这个结构记录着列表元素的值以及与其他节点之间的关系等信息，从列表中删除元素时，需要用到这个结构进行快速删除。

```go
package main

import "container/list"

func main() {
	l := list.New()

	l.PushBack("hhh")
	l.PushBack(67)

	e := l.PushBack("first")

	//在first之后添加元素second
	l.InsertAfter("second", e)

	//在first之前添加元素zero
	l.InsertBefore("zero", e)

	//删除first
	l.Remove(e)
}
```

**遍历链表**

```go
for i := li.Front(); i != nil; i=i.Next() {
    fmt.Print(i.Value)
}
```

# 空值nil

在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串`""`，而指针、切片、映射、通道、函数和接口的零值则是 nil

零值是Go语言中变量在声明之后但是未初始化被赋予的该类型的一个默认值。

nil 是Go语言中一个预定义好的标识符，有过其他编程语言开发经验的开发者也许会把 nil 看作其他语言中的 null（NULL），其实这并不是完全正确的，因为Go语言中的 nil 和其他语言中的 null 有很多不同点：

- nil 标识符之间是不能比较的

  <img align='left' src="img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618210930075.png" alt="image-20210618210930075" style="zoom:50%;" />

- nil 并不是Go语言的关键字或者保留字，也就是说我们可以定义一个名称为 nil 的变量

- nil没有默认类型

- 不同类型nil的指针是一样的，通过运行结果可以看出 arr 和 num 的指针都是 0x0。

  ![image-20210618211135658](img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618211135658.png)

- 不同类型的nil是不能比较的

  <img align='left' src="img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618211309016.png" alt="image-20210618211309016" style="zoom:50%;" />

- 两个相同类型的nil值也可能无法比较

  <img src="img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618211505463.png" alt="image-20210618211505463" style="zoom:50%;" />

- 在Go语言中 map、slice 和 function 类型的 nil 值不能比较，比较两个无法比较类型的值是非法的

- 不可比较类型的空值==可以直接与 nil 标识符进行比较==

  ![image-20210618211704518](img/Golang%E5%AE%B9%E5%99%A8.img/image-20210618211704518.png)

- 不同类型的nil值占用的内存大小可能是不一样的

  ```go
  package main
  
  import (
  	"fmt"
  	"unsafe"
  )
  
  func main() {
      //以下结果是在64位平台，32位平台减半
  	var p *struct{}
  	fmt.Println(unsafe.Sizeof(p)) // 8
  	var s []int
  	fmt.Println(unsafe.Sizeof(s)) // 24
  	var m map[int]bool
  	fmt.Println(unsafe.Sizeof(m)) // 8
  	var c chan string
  	fmt.Println(unsafe.Sizeof(c)) // 8
  	var f func()
  	fmt.Println(unsafe.Sizeof(f)) // 8
  	var i interface{}
  	fmt.Println(unsafe.Sizeof(i)) // 16
  }
  ```

# make和new关键字的区别

Go语言中 new 和 make 是两个内置函数，主要用来创建并分配类型的内存

- new：分配内存并放零值，new 函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针。同时 new 函数会把分配的内存置为零，也就是类型的零值。

  ```go
  var sum *int
  sum = new(int) //分配空间
  *sum = 98
  fmt.Println(*sum)
  ```

  ```go
  type Student struct {
     name string
     age int
  }
  var s *Student
  s = new(Student) //分配空间
  s.name ="dequan"
  fmt.Println(s)
  ```

  

- make： make 只能用于 slice、map 和 channel 的初始化。make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了

   make 函数的 t 参数必须是 chan（通道）、map（字典）、slice（切片）中的一个，并且返回值也是类型本身。

  ```go
  a := make([]int, 2)
  fmt.Println(a, b)  //[0,0]
  ```

**make和new区别小结：**

- new可以分配任意类型的内存并置零，make 只能用来分配及初始化类型为 slice、map、chan 的数据
- 返回值：new返回一个指向该类型内存地址的指针，make 返回引用，即 该类型
- new 分配的空间被清零。make 分配空间后，会进行初始化

# make和new原理

**make**

在编译期的类型检查阶段，Go语言其实就将代表 make 关键字的 OMAKE 节点根据参数类型的不同转换成了 OMAKESLICE、OMAKEMAP 和 OMAKECHAN 三种不同类型的节点，这些节点最终也会调用不同的运行时函数来初始化数据结构。

<img align='left' src="img/Golang%E5%AE%B9%E5%99%A8%E3%80%81nil%E3%80%81new%E3%80%81make.img/image-20210618214548889.png" alt="image-20210618214548889" style="zoom:30%;" />

**new**

内置函数 new 会在编译期的 SSA 代码生成阶段经过 callnew 函数的处理，如果请求创建的类型大小是 0，那么就会返回一个表示空指针的 zerobase 变量，在遇到其他情况时会将关键字转换成 newobject.

需要提到的是，哪怕当前变量是使用 var 进行初始化，在这一阶段也可能会被转换成 newobject 的函数调用并在堆上申请内存

当然这也不是绝对的，如果当前声明的变量或者参数不需要在当前作用域外生存，那么其实就不会被初始化在堆上，而是会初始化在当前函数的栈中并随着函数调用的结束而被销毁。

newobject 函数的工作就是获取传入类型的大小并调用 mallocgc 在堆上申请一片大小合适的内存空间并返回指向这片内存空间的指针

# 练习：快排

```go
// go语言实现快排

package main

import "fmt"

// 一次快排
func Partition(arp *[]int, left int, right int) int {
	base := (*arp)[left]

	for left < right {
		//从后往前遍历，找比基准小的数字放在前面
		for left < right && (*arp)[right] >= base {
			right--
		}
		if left < right {
			(*arp)[left] = (*arp)[right]
		}

		//从前往后遍历，找比基准大的数字放在后面
		for left < right && (*arp)[left] <= base {
			left++
		}
		if left < right {
			(*arp)[right] = (*arp)[left]
		}
	}
	(*arp)[left] = base
	return left
}

func quickSort(arp *[]int, left int, right int) {
	pos := Partition(arp, left, right)
	if pos-left > 1 {
		quickSort(arp, left, pos-1)
	}
	if right-pos > 1 {
		quickSort(arp, pos+1, right)
	}
}

func main() {
	arr := []int{1, 5, 0, 2, 6, 7, 12, 40, -1, 30}
	fmt.Println("排序前：", arr)

	quickSort(&arr, 0, len(arr)-1)
	fmt.Println("排序后：", arr)
}
```



