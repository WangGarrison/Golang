Go 语言提供了另外一种数据类型即接口（interface），它把所有的具有共性的方法定义在一起，这些方法只有函数签名，没有具体的实现代码（类似于Java中的抽象函数），任何其他类型只要实现了接口中定义好的这些方法，那么就说 这个类型实现（implement）了这个接口

接口类型是对其它类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我们可以让我们的函数更加灵活和更具有适应能力

接口是双方约定的一种合作协议。接口实现者不需要关心接口会被怎样使用，调用者也不需要关心接口的实现细节。接口是一种类型，也是一种抽象结构，不会暴露所含数据的格式、类型及结构。

Go语言的接口实现是隐式的，无须让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计。

==在类型中添加与接口签名一致的方法就可以实现该方法==

# 接口声明格式

```go
//每个接口类型由数个方法组成
type 接口类型名 interface {
    方法名1 (参数列表1) 返回值列表1  //没有返回值可以省略返回值列表
    方法名2 (参数列表2) 返回值列表2
    ...
}

/*
接口类型名：使用 type 将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫 Writer，有字符串功能的接口叫 Stringer，有关闭功能的接口叫 Closer 等。

方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。

参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以被忽略，例如：
type writer interface {
	Write ([]byte) error
}
*/
```

# 开发中常见的接口及写法

Go语言提供的很多包中都有接口，例如 io 包中提供的 Writer 接口：

```go
//这个接口可以调用 Write() 方法写入一个字节数组（[]byte），返回值告知写入字节数（n int）和可能发生的错误（err error）。
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

类似的，还有将一个对象以字符串形式展现的接口，只要实现了这个接口的类型，在调用 String() 方法时，都可以获得对象对应的字符串。在 fmt 包中定义如下：

```go
type Stringer interface {
	String() string
}
```

# 接口被实现的条件

- 接口的方法与实现接口的类型方法格式一致

- 接口中所有方法均被实现（当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用）

示例：数据写入器的抽象

![image-20210706094112376](img/Go4%EF%BC%9A%E6%8E%A5%E5%8F%A3interface.img/image-20210706094112376.png)

```go
package main

import "fmt"

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

//实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func main() {
    
	//定义一个f并实例化
	f := new(file)

	//声明一个DataWriter的接口
	var writer DataWriter // 声明接口

	writer = f //将 *file 类型的 f 赋值给 DataWriter 接口的 writer，
	//虽然两个变量类型不一致。但是 writer 是一个接口，
	//且 f 已经完全实现了 DataWriter() 的所有方法，因此赋值是成功的

	//使用接口进行数据写入
	writer.WriteData("data")
}
```

# 类型与接口的关系

在Go语言中类型和接口之间有==一对多和多对一==的关系

- 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。

- 多个类型也可以实现相同的接口

# 接口的nil判断

nil 在 Go语言中只能被赋值给指针和接口。

接口在底层的实现有两个部分：type 和 data。

- 在源码中，显式地将 nil 赋值给接口时，接口的 type 和 data 都将为 nil。此时，接口与 nil 值判断是相等的。
- 但如果将一个带有类型的 nil 赋值给接口时，此时，接口与 nil 判断将不相等。

# 接口赋值

接口是一种类型，接口也可以被赋值：

- 将对象赋值给接口：==只要类实现了该接口的所有方法，即可将该类对象赋值给这个接口==

  ```GO
  //定义接口
  type Testinterface interface{
      Teststring() string
      Testint() int
  }
   
  //定义结构体
  type TestMethod struct{
      name string
      age int
  }
   
  //结构体的两个方法隐式实现接口
  func (t *TestMethod)Teststring() string{
      return t.name
  }
   
  func (t *TestMethod)Testint() int{
      return t.age
  }
   
  func main(){
      T1 := &TestMethod{"ling",34}
      T2 := TestMethod{"gos",43}
      //接口本质是一种类型
      //接口赋值：只要类实现了该接口的所有方法，即可将该类赋值给这个接口
      var Test1 Testinterface  //接口只能是值类型
      Test1 = T1   //TestMethod类的指针类型实例传给接口
      fmt.Println(Test1.Teststring())
      fmt.Println(Test1.Testint())
   
      Test2 := T2   //TestMethod类的值类型实例传给接口
      fmt.Println(Test2.Teststring())
      fmt.Println(Test2.Testint())
  }
  ```

- 将接口赋值给另一个接口：只要两个接口拥有相同的方法列表（与次序无关），即是两个相同的接口，可以相互赋值

  接口赋值只需要接口A的方法列表是接口B的子集（即假设接口A中定义的所有方法，都在接口B中有定义），那么B接口的实例可以赋值给A的对象。反之不成立，即子接口B包含了父接口A，因此可以将子接口的实例赋值给父接口（==多的可以赋值给少的，类似于赋值兼容的切片==）

  （即子接口实例实现了子接口的所有方法，而父接口的方法列表是子接口的子集，则子接口实例自然实现了父接口的所有方法，因此可以将子接口实例赋值给父接口）

- 接口类型作为参数：因为可以将实现接口的类赋值给接口，而将接口类型作为参数很常见，所以：==那些实现接口的实例都能作为接口类型参数传递给函数/方法==

示例：

```GO
//Shaper接口
type Shaper interface {
	Area() float64
}

// Circle struct结构体
type Circle struct {
	radius float64
}

// Circle类型实现Shaper中的方法Area()
func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	c1 := new(Circle)
	c1.radius = 2.5  
	myArea(c1)  //将 Circle的指针类型实例c1传给函数myArea，接收类型为Shaper接口
}

func myArea(n Shaper) {
	fmt.Print(n.Area())
}
```

# 空接口

空接口是指没有定义任何接口方法的接口。**没有定义任何接口方法，意味着Go中的任意对象都已经实现空接口(因为没方法需要实现)，只要实现接口的对象都可以被接口保存，所以==任意对象都可以保存到空接口实例变量中==**。

> 空接口的内部实现保存了对象的类型和指针。使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢。因此在开发中，应在需要的地方使用空接口，而不是在所有地方使用空接口。

空接口的定义方式：

```go
type empty_int interface {}
```

更常见的，会直接使用`interface{}`作为一种类型，表示空接口。例如：

```go
// 声明一个空接口实例
var i interface{}
```

再比如函数使用空接口类型参数：

```go
func myfunc(i interface{})
```

可以定义一个空接口类型的array、slice、map、struct等，这样它们就可以用来存放任意类型的对象，因为任意类型都实现了空接口：

```GO
package main

import "fmt"

func main() {
	any := make([]interface{}, 5)
	any[0] = 11
	any[1] = "hello"
	any[2] = []int{1, 2, 3, 4, 5}
	for _, value := range any {
		fmt.Println(value)
	}
}
```

运行结果：

```go
11
hello
[1 2 3 4 5]
<nil>
<nil>
```

通过空接口类型，Go也能像其它动态语言一样，在数据结构中存储任意类型的数据。

**从空接口获取值**

保存到空接口的值，如果直接取出指定类型的值时，会发生编译错误，代码如下：

```go
// 声明a变量, 类型int, 初始值为1
var a int = 1
// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
var i interface{} = a
// 声明b变量, 尝试赋值i
var b int = i
```

![image-20210706162414363](img/Go4%EF%BC%9A%E6%8E%A5%E5%8F%A3%E3%80%81%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80.img/image-20210706162414363.png)

将 a 的值赋值给 i 时，虽然 i 在赋值完成后的内部值为 int，但 i 还是一个 interface{} 类型的变量。类似于无论集装箱装的是茶叶还是烟草，集装箱依然是金属做的，不会因为所装物的类型改变而改变。

可以使用类型断言来获取空接口的值：

```go
var ii int = i.(int)  
```

**空接口的值比较**

类型不同的空接口间的比较结果不相同

```go
// a保存整型
var a interface{} = 100
// b保存字符串
var b interface{} = "hi"
// 两个空接口不相等
fmt.Println(a == b)

输出：false
```

不能比较空接口中的动态值，当接口中保存有动态类型的值时，运行时将触发错误，代码如下：

```go
// c保存包含10的整型切片
var c interface{} = []int{10}
// d保存包含20的整型切片
var d interface{} = []int{20}
// 这里会发生崩溃
fmt.Println(c == d)  //这是一个运行时错误，提示 []int 是不可比较的类型。
```

下表中列举出了类型及比较的几种情况：

| 类  型          | 说  明                                                       |
| --------------- | ------------------------------------------------------------ |
| map             | 宕机错误，不可比较                                           |
| 切片（[]T）     | 宕机错误，不可比较                                           |
| 通道（channel） | 可比较，必须由同一个 make 生成，也就是同一个通道才会是 true，否则为 false |
| 数组（[容量]T） | 可比较，编译期知道两个数组是否一致                           |
| 结构体          | 可比较，可以逐个比较结构体的值                               |
| 函数            | 可比较                                                       |

# 接口嵌套

一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用。

嵌套的内部接口将属于外部接口，内部接口的方法也将属于外部接口。

另外在类型嵌套时，如果内部类型实现了接口，那么外部类型也会自动实现接口，因为内部属性是属于外部属性的。

```go
type ReadWrite interface {
    Read(b Buffer) bool
    Write(b Buffer) bool
}
  
type Lock interface {
    Lock()
    Unlock()
}

type File interface {　　
    ReadWrite　//ReadWrite为内部接口
    Lock  //Lock为内部接口
    Close()
}
```

# 类型断言

在Go语言中类型断言的语法格式如下：

```GO
value, ok := x.(T)
```

- 判断变量x的类型是否是T
- value值接收变量x的值
- ok接收bool值，根据该布尔值判断 x 是否为 T 类型

示例：

```GO
package main
import (
    "fmt"
)
func main() {
    var x interface{}
    x = 10
    value, ok := x.(int)
    fmt.Print(value, ",", ok)
}

输出：
10,true
```

注意：`value, ok := x.(T)`

- 如果 T 是具体某个类型，类型断言会检查 x 的动态类型是否等于具体类型 T。如果检查成功，类型断言返回的结果是 x 的动态值，其类型是 T。
- 如果 T 是接口类型，类型断言会检查 x 的动态类型是否满足 T，==即类型x是否实现了接口T==。如果检查成功，x 的动态值不会被提取，返回值是一个类型为 T 的接口值。
- 无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败

类型断言还可以配合 switch 使用，示例代码如下：

```GO
func main() {
    var a int
    a = 10
    getType(a)
}

func getType(a interface{}) {  //任意类型都实现了空接口，所以任意类型可以赋值给空接口
    switch a.(type) {
    case int:
        fmt.Println("the type of a is int")
    case string:
        fmt.Println("the type of a is string")
    case float64:
        fmt.Println("the type of a is float")
    default:
        fmt.Println("unknown type")
    }
}
```

**断言实现接口和类型之间的转换**

Go语言中可以使用接口断言（type assertions）将接口转换成另外一个接口，也可以将接口转换为另外的类型。

```go
t := i.(T)  //i 代表接口变量，T 代表转换的目标类型，t 代表转换后的变量。
```

转换有两种情况：

- 如果断言的类型 T 是一个具体类型，然后类型断言检查 i 的动态类型是否和 T 相同。如果这个检查成功了，类型断言的结果是 i 的动态值，当然它的类型是 T
- 如果相反断言的类型 T 是一个接口类型，然后类型断言检查是否 i 的动态类型满足 T。如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同类型和值部分的接口值

**将接口转换为其他类型**

==只要类实现了该接口的所有方法，即可将该类对象赋值给这个接口==

```go
p1 := new(pig)
var a Walker = p1  //Walker是接口类型
p2 := a.(*pig)  //由于 pig 实现了 Walker 接口，因此可以被隐式转换为 Walker 接口类型保存于 a 中
```

# 使用接口实现多态

Golang中的多态是通过接口来实现的，主要用的是如下几条特性来实现多态：

- 多个类型（结构体）可以实现同一个接口
- 一个类型（结构体）可以实现多个接口
- 实现接口的类（结构体）可以赋值给接口

```go
package main

import "fmt"

// Person接口
type Person interface {
	SayHello()
}

type Girl struct {
	Sex string
}

type Boy struct {
	Sex string
}

// Boy类的方法实现了Person的接口
func (this *Boy) SayHello() {
	fmt.Println("Hi, I am a " + this.Sex)
}

// Girl类的方法实现了Person的接口
func (this *Girl) SayHello() {
	fmt.Println("Hi, I am a " + this.Sex)
}

func main() {
	g := &Girl{"girl"}
	b := &Boy{"boy"}

    p := map[int]Person{}  //Girl类和Boy类都实现了Person接口，所以可以赋值给Person{}
	p[0] = g
	p[1] = b

	for _, v := range p {
		v.SayHello()
	}
}
```

Girl类和Boy类都实现了Person接口，所以可以赋值给Person{}，代码运行输出如下：

<img align='left' src="img/Go4%EF%BC%9A%E6%8E%A5%E5%8F%A3%E3%80%81%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80.img/image-20210706151807934.png" alt="image-20210706151807934" style="zoom:67%;" />

# sort包的sort.Interface接口

一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是 sort.Interface 的三个方法：

```go
package sort
type Interface interface {
    Len() int            // 获取元素数量
    Less(i, j int) bool  // 比较两个元素的方法，i，j是序列元素的指数
    Swap(i, j int)       // 交换元素的方法
}
```

对序列进行排序时，我们可以定义一个实现了这三个方法的类型，然后对这个类型的一个实例应用 sort.Sort 函数，来达到特定目的的排序

通过实现 sort.Interface 接口的排序过程具有很强的可定制性，可以根据被排序对象比较复杂的特性进行定制

```go
package main

import (
	"fmt"
	"sort"
)

// 将[]string定义为MyStringList类型
type MyStringList []string

// 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	// 准备一个内容被打乱顺序的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 使用sort包进行排序
	sort.Sort(names)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
```

# 常见类型的便捷排序

需要多种排序逻辑的需求适合使用 sort.Interface 接口进行排序。但大部分情况中，只需要对字符串、整型等进行快速排序。

Go语言中提供了一些固定模式的封装以方便开发者迅速对内容进行排序，即==有一些特定的类型它已经内部实现好了sort.Interface的几种方法==，所以直接定义好该类型，然后调用`sort.Sort(names)`就可以了

sort 包中定义了一些常见类型的排序方法，如下表所示：

| 类  型                | 实现 sort.lnterface 的类型 | 直接排序方法               | 说  明            |
| --------------------- | -------------------------- | -------------------------- | ----------------- |
| 字符串（String）      | StringSlice                | sort.Strings(a [] string)  | 字符 ASCII 值升序 |
| 整型（int）           | IntSlice                   | sort.Ints(a []int)         | 数值升序          |
| 双精度浮点（float64） | Float64Slice               | sort.Float64s(a []float64) | 数值升序          |

**字符串切片的便捷排序**

sort 包中有一个 StringSlice 类型，定义如下：

```go
type StringSlice []string
func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// Sort is a convenience method.
func (p StringSlice) Sort() { Sort(p) }
```

sort 包中的 StringSlice 的代码与 MyStringList 的实现代码几乎一样。因此，只需要使用 sort 包的 StringSlice 就可以更简单快速地进行字符串排序。将代码1中的排序代码简化后如下所示：

```go
names := sort.StringSlice{
    "3. Triple Kill",
    "5. Penta Kill",
    "2. Double Kill",
    "4. Quadra Kill",
    "1. First Blood",
}
sort.Sort(names)
```

或者：

```go
names := []string{
    "3. Triple Kill",
    "5. Penta Kill",
    "2. Double Kill",
    "4. Quadra Kill",
    "1. First Blood",
}
sort.Strings(names)
```

**对整型切片进行排序**

```go
nums := []int{10, 21, 32, 3, 17}
sort.Ints(nums)
```

注意：编程中经常用到的 int32、int64、float32、bool 类型并没有由 sort 包实现，使用时依然需要开发者自己编写

**使用sort.Slice进行切片元素排序**

Go语言在 sort 包中提供了 sort.Slice() 函数进行更为简便的排序方法。sort.Slice() 函数只要求传入需要排序的数据，以及一个排序时对元素的回调函数，类型为 func(i,j int)bool，sort.Slice() 函数的定义如下：

```go
func Slice(slice interface{}, less func(i, j int) bool)
```

该函数类似于C++中sort函数的第三个参数

使用示例：

```go
sort.Slice(heros, func(i, j int) bool {
    if heros[i].Kind != heros[j].Kind {
        return heros[i].Kind < heros[j].Kind
    }
    return heros[i].Name < heros[j].Name
})
```

# error接口

Go语言中引入 error 接口类型作为错误处理的标准模式，如果函数要返回错误，则返回值类型列表中肯定包含 error。error 处理过程类似于C语言中的错误码，可逐层返回，直到被处理

Go语言中返回的 error 类型究竟是什么呢？查看Go语言的源码就会发现 error 类型是一个非常简单的接口类型，如下所示：

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}
```

error 接口有一个签名为 Error() string 的方法，所有实现该接口的类型都可以当作一个错误类型。Error() 方法给出了错误的描述，在使用 fmt.Println 打印错误时，会在内部调用 Error() string 方法来得到该错误的描述

一般情况下，如果函数需要返回错误，就将 error 作为多个返回值中的最后一个（但这并非是强制要求）。

```go
package main
import (
    "errors"
    "fmt"
    "math"
)
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return -1, errors.New("math: square root of negative number")
    }
    return math.Sqrt(f), nil
}
func main() {
    result, err := Sqrt(-13)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}
```

# Go语言接口内部实现

接口的实例化：具体类型实例传递给接口称为接口的实例化

![image-20210707102000988](img/Go4%EF%BC%9A%E6%8E%A5%E5%8F%A3%E3%80%81%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80%E3%80%81sort%E7%9B%B8%E5%85%B3.img/image-20210707102000988.png)

## 非空接口数据结构

非空接口的底层数据结构是 ==iface==，源代码位于：`src/runtime/runtime2.go`

**iface数据结构**

非空接口初始化的过程就是初始化一个 iface 类型的结构，示例如下：

```go
//src/runtime/runtime2.go
type iface struct {
    tab *itab                //itab 存放类型及方法指针信息
    data unsafe.Pointer      //数据信息，指向接口绑定的实例的副本，（接口的初始化也是一种值拷贝）
}
```

iface里面的==itab==结构体是接口内部实现的核心和基础

**itab数据结构**

```go
type itab struct {
    //接口自身的静态类型，（是指向接口类型元信息的指针）
    inter *interfacetype      
    
     //_type 就是接口存放的具体实例的类型（动态类型），iface 里的 data 指针指向的是该类型的值
    _type *_type             
    
    //hash 存放具体类型的 Hash 值，（这里冗余存放主要是为了接口断言或类型查询时快速访问）
     // “copy of _type.hash. Used for type switches.”
    hash uint32
    
    _   [4]byte
    
    //一个函数指针，可以理解为 C++ 对象模型里面的虚拟函数指针，这里虽然只有一个元素，实际上指针数组的大小是可变的，编译器负责填充，运行时使用底层指针进行访问，不会受 struct 类型越界检查的约束，这些指针指向的是具体类型的方法。
    fun [1]uintptr            // variable sized. fun[0]==0 means _type does not implement inter.
}
```

itab 这个数据结构是非空接口实现动态调用的基础，itab 的信息被编译器和链接器保存了下来，存放在可执行文件的只读存储段（==.rodata==）中。itab 存放在静态分配的存储空间中，不受 GC 的限制，其内存不会被回收。

**_type结构体**

Go语言是一种强类型的语言，编译器在编译时会做严格的类型校验。所以 Go 必然为每种类型维护一个类型的元信息，这个元信息在运行和反射时都会用到，Go语言的类型元信息的通用结构是 _type，其他类型都是以 _type 为内嵌字段封装而成的结构体

_type 包含所有类型的共同元信息，编译器和运行时可以根据该元信息解析具体类型、类型名存放位置、类型的 Hash 值等基本信息。

：Go语言类型元信息最初由编译器负责构建，并以表的形式存放在编译后的对象文件中，再由链接器在链接时进行段合并、符号重定向（填充某些值）

**itab里面的interfacetype**

interfacetype元信息包含了_type数据结构，然后又封装了一些接口自己需要的东西

```go
//描述接口的类型
type interfacetype struct {
    typ _type       //类型通用部分
    pkgpath name    //接口所属包的名字信息， name 内存放的不仅有名称，还有描述信息
    mhdr []imethod  //接口的方法
}
//接口方法元信息
type imethod struct {
    name nameOff //方法名在编译后的 section 里面的偏移量
    ityp typeOff //方法类型在编译后的 section 里面的偏移量
}
```

## 接口调用过程分析

```go
//iface.go
package main
type Caler interface {
    Add (a , b int) int
    Sub (a , b int) int
}
type Adder struct ｛id int }

//go:noinline  "//go:noinline"告诉编译器不要内联
func (adder Adder) Add(a, b int) int { return a + b }

//go:noinline
func (adder Adder) Sub(a , b int) int { return a - b }

func main () {
    var m Caler=Adder{id: 1234}
    m.Add(10, 32)
}
```

接口的动态调用分为两个阶段：

- 第一阶段就是构建 iface 动态数据结构，这一阶段是在接口实例化的时候完成的，映射到 Go 语句就是`var m Caler = Adder{id: 1234}`。
- 第二阶段就是通过函数指针间接调用接口绑定的实例方法的过程，映射到 Go 语句就是 m.Add(10, 32)。

这个过程有两部分多余时耗，一个是接口实例化的过程，也就是 iface 结构建立的过程，一旦实例化后，这个接口和具体类型的 itab 数据结构是可以复用的；另一个是接口的方法调用，它是一个函数指针的间接调用

## 空接口数据结构

前面我们了解到==空接口 interface{} 是没有任何方法集的接口==，所以空接口内部不需要维护和动态内存分配相关的数据结构 itab 。空接口只关心存放的具体类型是什么，具体类型的值是什么，所以空接口的底层数据结构也很简单，具体如下：

**eface数据结构**

```go
//go/src/runtime/runtime2.go
//空接口
type eface struct {
    _type *_type
    data unsafe.Pointer
}
```

从 eface 的数据结构可以看出，空接口不是真的为空，其保留了具体实例的类型和值拷贝，即便存放的具体类型是空的，空接口也不是空的。

由于空接口自身没有方法集，所以空接口变量实例化后的真正用途不是接口方法的动态调用。==空接口在Go语言中真正的意义是支持多态==

# Go实现Web服务器

使用 net/http 包能很简单地对 Web 的路由，静态文件，模版，cookie 等数据进行设置和操作

**Web服务器工作方式：**

- 浏览器本身是一个客户端，当在浏览器中输入URL（网址）时，首先浏览器会请求DNS服务器，通过DNS服务器获取域名相应的IP，然后通过IP地址找到对应的服务器后，请求建立TCP连接
- 与服务器建立连接后，浏览器会向服务器发送 HTTP Request （请求）包
- 服务器接收到请求包之后开始处理请求包，并调用自身服务，返回 HTTP Response（响应）包
- 客户端收到来自服务器的响应后开始渲染这个 Response 包里的主体（body），等收到全部的内容后断开与该服务器之间的 TCP 连接
- <img align='left' src="img/Go4%EF%BC%9A%E6%8E%A5%E5%8F%A3%E3%80%81%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80%E3%80%81sort%E7%9B%B8%E5%85%B3.img/image-20210707102627097.png" alt="image-20210707102627097" style="zoom:40%;" />

**搭建一个简单的Web服务器**

```go
package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	//注册处理请求的函数，第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求
	http.HandleFunc("/index", index)

	//开始监听：http.ListenAndServe监听，参数是服务器要监听的主机地址和端口号
	http.ListenAndServe("localhost:8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./index.html")
	w.Write(content)
}
```

index.html文件：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8"
    <title>汪汪</title>
</head>
<body>
    <h1>汪汪</h1>
</body>
</html>
```

