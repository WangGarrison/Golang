# 定义结构体

结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”。字段有以下特性：

- 字段拥有自己的类型和值。
- 字段名必须唯一。
- 字段的类型也可以是结构体，甚至是字段所在结构体的类型。

Go 语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。但 Go 语言中结构体的内嵌配合接口比面向对象具有更高的扩展性和灵活性

结构体定义格式如下：（结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存）

```go
type  类型名  struct {
    字段1  类型
    字段2  类型
}

//type 类型名 struct{} 可以理解为将 struct{} 结构体定义为类型名的类型。类似：type NewInt int

type Point struct {
    X int
    Y int
}

type Color struct {
    R, G, B byte
}
```

# 实例化结构体

```go
var p Point
p.X = 10
p.Y = 20
```

Go语言中，还可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成指针类型的结构体。

**使用 new 的实例化如下：**

```go
p := new(Point)  //p的类型位*Point
p.X = 10  //Go语言可以像访问普通结构体一样使用`.`来访问结构体指针的成员。
```

在 C/C++ 语言中，使用 new 实例化类型后，访问其成员变量时必须使用`->`操作符。

在Go语言中，访问结构体指针的成员变量时可以继续使用`.`，这是因为Go语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术，将 ins.Name 形式转换为 (*ins).Name。

**取结构体地址的实例化：**

```go
p := &Point{}  //对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
p.X = 10
```

取地址实例化是最广泛的一种结构体实例化方式：

```go
type Point struct {
	X int
	Y int
}

func main() {
	p := &Point{
		X: 10,
		Y: 20,
	}
	fmt.Println(p.X, p.Y)
}
```

结构体实例化后字段的默认值是字段类型的默认值，例如 ，数值为 0、字符串为 ""（空字符串）、布尔为 false、指针为 nil 等。

# 初始化结构体的成员变量

- 键值对形式的初始化：适合选择性填充字段较多的结构体
- 多个值的列表形式初始化：适合填充字段较少的结构体。

**键值对形式的初始化**

结构体可以使用“键值对”（Key value pair）初始化字段，每个“键”（Key）对应结构体中的一个字段，键的“值”（Value）对应字段需要初始化的值。

键值对的填充是可选的，不需要初始化的字段可以不填入初始化列表中。

注意：==键值之间以`:`分隔，键值对之间以`,`分隔==

```go
p2 := Point{
    X: 10,
    Y: 20,
}

p := &Point{  //对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
    X: 10,
    Y: 20,
}
```

**列表形式初始化**

Go语言可以在“键值对”初始化的基础上忽略“键”，也就是说，可以使用多个值的列表初始化结构体的字段。

- 必须初始化结构体的所有字段。
- 每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
- 键值对与值列表的初始化形式不能混用。

```go
p3 := Point{
    10,
    20,
}
```

**初始化匿名结构体**

```go
ins := struct {
    // 匿名结构体字段定义
    字段1 字段类型1
    字段2 字段类型2
    …
}{
    // 字段值初始化
    初始化字段1: 字段1的值,
    初始化字段2: 字段2的值,
    …
}
```

键值对初始化部分是可选的，不初始化成员时，匿名结构体的格式变为：

```go
ins := struct {
    字段1 字段类型1
    字段2 字段类型2
    …
}
```

# 构造函数

Go语言的类型或结构体没有构造函数的功能，但是我们可以使用结构体初始化的过程来模拟实现构造函数。

```go
package main

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

```

**带有父子关系的结构体的构造和初始化——模拟父级构造调用**

（组合代替继承）

```go
type Cat struct {
    Color string
    Name  string
}
type BlackCat struct {
    Cat  // 嵌入Cat, 类似于派生
}
// “构造基类”
func NewCat(name string) *Cat {
    return &Cat{
        Name: name,
    }
}
// “构造子类”
func NewBlackCat(color string) *BlackCat {
    cat := &BlackCat{}
    cat.Color = color
    return cat
}
```

# 方法和接收器

> 函数：函数是一一映射的关系，给定一个变量，则会出现一个确定的值。 是指一段可以直接被其名称调用的代码块，它可以传入一些参数进行处理并返回一些数据，所有传入函数的数据都是被明确定义。
>
> 方法：类的成员函数。方法是特殊的函数，可以说是函数的子集，==方法的作用对象是接收器，也就是类实例==，而函数没有作用对象
>
> 接收器：类的实例，是方法作用的目标

在Go语言中，结构体就像是类的一种简化形式，那么类的方法在哪里呢？

在Go语言中有一个概念，它和方法有着同样的名字，并且大体上意思相同，==Go 方法是作用在接收器（receiver）上的一个函数==，接收器是某种类型的变量，因此==方法是一种特殊类型的函数==。

在Go语言中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在不同的源文件中，唯一的要求是它们必须是同一个包的。

类型 T（或 T）上的所有方法的集合叫做类型 T（或 T）的方法集。

因为方法是函数，所以同样的，==不允许方法重载，即对于一个类型只能有一个给定名称的方法==，但是如果基于接收器类型，是有重载的：==具有同样名字的方法可以在 2 个或多个不同的接收器类型上存在==，比如在同一个包里这么做是允许的。

在面向对象的语言中，类拥有的方法一般被理解为类可以做的事情。在Go语言中“方法”的概念与其他语言一致，只是Go语言建立的“接收器”强调==方法的作用对象是接收器，也就是类实例，而函数没有作用对象==。

**方法作用于接收器的格式如下：**

```go
func (接收器变量名 接收器类型) 方法名(参数列表)(返回参数) {   //格式就是在函数格式的基础上加了个接收器的名字和类型
    ...
}
//接收器中的参数变量名在命名时，官方建议使用接收器类型名的第一个小写字母，如：s Socket, c Connector
```

**为结构体添加方法**

```go
package main

type Bag struct {
	items []int
}

//Bag对象实例的Insert方法
//Insert(itemid int) 的写法与函数一致，(b*Bag) 表示接收器，即 Insert 作用的对象实例
func (b *Bag) Insert(itemid int) {  
	b.items = append(b.items, itemid)
}

func main() {
	b := new(Bag)
	b.Insert(1001)  //可以愉快地像其他语言一样，用面向对象的方法来调用 b 的 Insert
}

```

每个方法只能有一个接收器，如下图所示。

<img align='left' src="img/Go%EF%BC%9Astruct%E7%BB%93%E6%9E%84%E4%BD%93.img/image-20210620134419740.png" alt="image-20210620134419740" style="zoom:30%;" />

接收器类型可以是（几乎）任何类型，不仅仅是结构体类型，任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型，但是接收器不能是一个接口类型，因为接口是一个抽象定义，而方法却是具体实现，如果这样做了就会引发一个编译错误

接收器根据接收器的类型可以分为指针接收器、非指针接收器，两种接收器在使用时会产生不同的效果，根据效果的不同，两种接收器会被用于不同性能和功能要求的代码中。

**指针类型接收器**

指针类型的接收器由一个结构体的指针组成，更接近于面向对象中的 this，由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的

```go
package main

import "fmt"

type Property struct {
	value int
}

//Property的设置值的方法
func (p *Property) SetValue(v int) {  //指针类型接收器
	p.value = v
}

//Property的获取值的方法
func (p *Property) GetValue() int {
	return p.value
}

func main() {
	p := new(Property)

	p.SetValue(10)

	fmt.Println(p.GetValue())
}
```

**非指针类型的接收器**

当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效。

```go
package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) Add(o Point) Point {  //非指针类型接收器
	return Point{p.X + o.X, p.Y + o.Y}
}

func main() {
	p1 := Point{1, 1}
	p2 := Point{2, 2}

	result := p1.Add(p2)

	fmt.Println(result)
}
```

在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器，大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针

# 为任意类型添加方法

Go语言可以对任何类型添加方法，给一种类型添加方法就像给结构体添加方法一样，因为结构体也是一种类型。

代码示例：将int重命名为MyInt，并为该类型添加判0方法、Add方法

```go
package main

import "fmt"

type MyInt int

//为MyInt添加IsZero方法
func (m MyInt) IsZero() bool {  
	return m == 0
}

//为MyInt添加Add()方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero())

	b = 1
	fmt.Println(b.Add(2))
}
```

# 函数变量保存方法/函数

无论是普通函数还是结构体的方法，只要它们的签名一致，与它们签名一致的函数变量就可以保存普通函数或是结构体方法。

如下：

```go
// 声明一个结构体
type class struct {
}
// 给结构体添加Do方法
func (c *class) Do(v int) {
    fmt.Println("call method do:", v)
}
// 普通函数的Do
func funcDo(v int) {
    fmt.Println("call function do:", v)
}
func main() {
    
    var delegate func(int)  // 声明一个函数变量/回调
    
    c := new(class)  // 创建结构体实例
    
    delegate = c.Do  // 将回调设为c的Do方法
    delegate(100)
    
    delegate = funcDo  // 将回调设为普通函数
    delegate(100)
}
```

# 类型内嵌和结构体内嵌

Go语言中的继承是通过内嵌或组合来实现的，所以可以说，在Go语言中，相比较于继承，组合更受青睐。

结构体==可以包含一个或多个匿名（或内嵌）字段==，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型也就是字段的名字

如下代码：

```go
type innerS struct {
    in1 int
    in2 int
}
type outerS struct {
    b int
    c float32
    int // 匿名字段
    innerS //匿名字段
}
func main() {
    outer := new(outerS)
    outer.b = 6
    outer.c = 7.5
    outer.int = 60  //给匿名字段赋值
    outer.in1 = 5   //直接访问内嵌结构体的成员变量
    outer.in2 = 10
    fmt.Printf("outer.b is: %d\n", outer.b)
    fmt.Printf("outer.c is: %f\n", outer.c)
    fmt.Printf("outer.int is: %d\n", outer.int)
    fmt.Printf("outer.in1 is: %d\n", outer.in1)
    fmt.Printf("outer.in2 is: %d\n", outer.in2)
    // 使用结构体字面量
    outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
    fmt.Printf("outer2 is:", outer2)
}
```

**结构体内嵌的特性：**（上例中outers结构体中的inners就是一个匿名的内嵌结构体）

- 内嵌的匿名结构体可以直接访问其成员变量，也可以通过类型名加点间接访问

```go
package main

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b      int
	c      float32
	int    // 匿名字段
	innerS //匿名字段
}

func main() {
	var outer outerS
	outer.in1 = 5        //直接访问内嵌结构体的成员变量
	outer.innerS.in1 = 5 //类型名加点间接访问
}
```

- 非匿名的内嵌结构体要通过名称加点来访问

```go
package main

type Hands struct {
	leftHand  int
	rightHand int
}

type Body struct {
	h Hands
	l int
}

func main() {
	var body Body
	body.h.leftHand = 1
	body.h.rightHand = 2
}
```

# 内嵌结构体成员名字冲突

当组合的多个匿名基类有重名成员时，访问这个重名成员时，需要通过类型名加点间接访问，不能直接用.访问了

```go
package main

type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

func main() {
	c := &C{}
	c.A.a = 1  //ok
	c.B.a = 2  //ok
    c.a = 2  //error: ambiguous selector模糊不清的选择
}
```

# 结构体内嵌模拟类的继承

Go语言中的==继承是通过内嵌或组合来实现==的，所以可以说，在Go语言中，相比较于继承，组合更受青睐。

示例：人类不能飞行，鸟类可以飞行。人类和鸟类都可以继承自可行走类，但只有鸟类继承自飞行类。

```go
package main

import "fmt"

//可飞行的
type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

//可行走的
type Walkable struct{}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

//人类
type Human struct {
	Walkable //人类能行走
}

//鸟类
type Bird struct {
	Walkable //鸟类既能行走也能飞行
	Flying
}

func main() {
	h := new(Human)
	h.Walk()

	b := new(Bird)
	b.Walk()
	b.Fly()
}
```

# 初始化内嵌结构体

匿名结构体内嵌初始化时，将结构体内嵌的类型作为字段名像普通结构体一样进行初始化

```go
package main

import "fmt"

// 车轮
type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Power int    // 功率
	Type  string // 类型
}

// 车
type Car struct {
	Wheel  //匿名
	eng Engine  //非匿名
}

func main() {
	c := Car{
		// 初始化匿名内嵌
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化非匿名内嵌
		eng: Engine{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}
```

在前面描述车辆和引擎的例子中，有时考虑编写代码的便利性，会==将结构体直接定义在嵌入的结构体中==。也就是说，结构体的定义不会被外部引用到。在初始化这个被嵌入的结构体时，就==需要再次声明结构才能赋予数据==：

```go
package main
import "fmt"
// 车轮
type Wheel struct {
    Size int
}
// 车
type Car struct {
    Wheel
    // 引擎
    Engine struct {
        Power int    // 功率
        Type  string // 类型
    }
}
func main() {
    c := Car{
        // 初始化轮子
        Wheel: Wheel{
            Size: 18,
        },
        // 初始化引擎
        Engine: struct {  //需要再次声明
            Power int
            Type  string
        }{
            Type:  "1.4T",
            Power: 143,
        },
    }
    fmt.Printf("%+v\n", c)
}
```

# Go语言垃圾回收和SetFinalizer

Go语言自带垃圾回收机制（Garbage Collection，GC）。GC 通过独立的进程执行，它会搜索不再使用的变量，并将其释放。需要注意的是，GC 在运行时会占用机器资源。

GC 是自动进行的，如果要手动进行 GC，可以使用 runtime.GC() 函数，显式的执行 GC。显式的进行 GC 只在某些特殊的情况下才有用，比如当内存资源不足时调用 runtime.GC() ，这样会立即释放一大片内存，但是会造成程序短时间的性能下降。

finalizer（终止器）是与对象关联的一个函数，通过 runtime.SetFinalizer 来设置，==如果某个对象定义了 finalizer，当它被 GC 回收的时候，这个 finalizer 就会被调用==，类似于析构函数执行的时机，以完成一些特定的任务，例如发信号或者写日志等。

SetFinalizer函数签名：

```go
func SetFinalizer(x, f interface{})  //将对象x的终止器设置为f，当GC发现 x 不能再直接或间接访问时，GC会清理 x 并调用f(x)
```

- 参数 x 必须是一个指向通过 new 申请的对象的指针，或者通过对复合字面值取址得到的指针
- 参数 f 必须是一个函数，它接受单个可以直接用 x 类型值赋值的参数，也可以有任意个被忽略的返回值

另外，x 的终止器会在 x 不能直接或间接访问后的任意时间被调用执行，不保证终止器会在程序退出前执行，因此一般终止器只用于在长期运行的程序中释放关联到某对象的非内存资源。例如，当一个程序丢弃一个 os.File 对象时没有调用其 Close 方法，该 os.File 对象可以使用终止器去关闭对应的操作系统文件描述符。

此外，我们也可以使用`SetFinalizer(x, nil)`来清理绑定到 x 上的终止器。

注意：终止器只有在对象被 GC 时，才会被执行。其他情况下，都不会被执行，即使程序正常结束或者发生错误。

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

type Road int

func findRoad(r *Road) {
	fmt.Println("road:", *r)
}

func entry() {
	var rd Road = Road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad) //设置局部变量r的终止器为findRoad
}

func main() {
	entry()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}
```



