# 包简介

Go 语言的源码复用建立在包（package）基础之上。Go 语言的入口 main() 函数所在的包（package）叫 main，main 包想要引用别的代码，必须同样以包的方式进行引用

Go 语言的包与文件夹一一对应，所有与包相关的操作，必须依赖于工作目录（GOPATH）。Go语言的包借助了目录树的组织形式，==一般包的名称就是其源文件所在目录的名称==，虽然Go语言没有强制要求包名必须和其所在的目录名同名，但还是建议包名和所在目录同名，这样结构更清晰

任何源代码文件必须属于某个包，同时源码文件的第一行有效代码必须是`package pacakgeName `语句，通过该语句声明自己所在的包。

包可以定义在很深的目录中，包名的定义是不包括目录路径的，但是包在引用时一般使用全路径引用。比如在`GOPATH/src/a/b/ `下定义一个包 c。**在包 c 的源码中只需声明为`package c`，而不是声明为`package a/b/c`，但是在导入 c 包时，需要带上路径，例如`import "a/b/c"`**。

**注意：==一个文件夹下的所有源码文件只能属于同一个包，同样属于同一个包的源码文件不能放在多个文件夹下==**

# 包的导入

```go
import "包的路径"
```

```go
import (
    "包 1 的路径"
    "包 2 的路径"
)  
/*
导入的包之间可以通过添加空行来分组；通常将来自不同组织的包独自分组。包的导入顺序无关紧要，但是在每个分组中一般会根据字符串顺序排列
*/
```

注意事项：

- import 导入语句通常放在源码文件开头包声明语句的下面；
- 导入的包名需要使用双引号包裹起来；
- 包名是从`GOPATH/src/ `后开始计算的，使用`/ `进行路径分隔。
- ==使用 import 语句导入包时，使用的是包所属文件夹的名称==；
- 自定义包的包名不必与其所在文件夹的名称保持一致，但为了便于维护，建议保持一致；

**可以自定义包别名**

```go
import F "fmt"  //在代码中可以使用：F.Println("hello")
```

**省略引用格式**

```go
//这种格式相当于把 fmt 包直接合并到当前程序中，在使用 fmt 包内的方法是可以不用加前缀fmt.，直接引用
import . "fmt"
```

**匿名引用格式**

```go
//在引用某个包时，如果只是希望执行包初始化的 init 函数，而不使用包内部的数据时，可以使用匿名引用格式
import _ "fmt"
```

**关于`import _ " "`的说明**

我们都知道，go对包导入非常严格，**不允许导入不使用的包**。但是有时候我们导入包只是为了做一些初始化的工作，这样就应该采用`import _ " "`的形式

**注意：**

- 包不能出现环形引用的情况，比如包 a 引用了包 b，包 b 引用了包 c，如果包 c 又引用了包 a，则编译不能通过。
- 包的重复引用是允许的，比如包 a 引用了包 b 和包 c，包 b 和包 c 都引用了包 d。这种场景相当于重复引用了 d，这种情况是允许的，并且 Go 编译器保证包 d 的 init 函数只会执行一次

# 包加载流程

执行 main 包的 main函数之前， Go 引导程序会先对整个程序的包进行初始化。整个执行的流程如下图所示

![image-20210707112517472](img/Go5%EF%BC%9Apackage.img/image-20210707112517472.png)

Go语言包的初始化有如下特点：

- 包初始化程序从 main 函数引用的包开始，逐级查找包的引用，直到找到没有引用其他包的包，最终生成一个包引用的有向无环图。
- Go 编译器会将有向无环图转换为一棵树，然后从树的叶子节点开始逐层向上对包进行初始化。
- 单个包的初始化过程如上图所示，先初始化常量，然后是全局变量，最后执行包的 init 函数。

# 包的init函数

每个包可以有一个init函数，在程序执行前被系统自动调用，调用顺序如下图：

![image-20210707143900307](img/Go5%EF%BC%9A%E5%8C%85package.img/image-20210707143900307.png)

在运行时，被最后导入的包会最先初始化并调用 init() 函数。

例如，假设有这样的包引用关系：main→A→B→C，那么这些包的 init() 函数调用顺序为：

```go
C.init→B.init→A.init→main
```

# Go实现封装：成员名小写

在Go语言中封装就是把抽象出来的字段和对字段的操作封装在一起，数据被保护在内部，程序的其它包只能通过被授权的方法，才能对字段进行操作。

 如何体现封装：

- 变量名==首字母小写则只可以在本包中访问，首字母大写则其他包中也可以访问==
- 通过方法，包，实现封装。

![image-20210707134441243](img/Go5%EF%BC%9A%E5%8C%85package.img/image-20210707134441243.png)

# GOPATH

GOPATH 是 Go语言中使用的一个环境变量，它使用绝对路径提供项目的工作目录

在命令行使用`go env`可以查看相关的环境变量

在 GOPATH 指定的工作目录下：

- 代码总是会保存在 **$GOPATH/src** 目录下
- 在工程经过 go build、go install 或 go get 等指令后，会将产生的二进制可执行文件放在 **$GOPATH/bin** 目录下
- 生成的中间缓存文件会被保存在 **$GOPATH/pkg** 下。

如果需要将整个源码添加到版本管理工具（Version Control System，VCS）中时，只需要添加 $GOPATH/src 目录的源码即可。bin 和 pkg 目录的内容都可以由 src 目录生成。

**设置和使用GOPATH**（Linux下）

1）设置当前目录为GOPATH

```shell
export GOPATH=`pwd`
```

2）建立源码目录

```sh
mkdir -p src/hello  #mkdir 指令的 -p 可以连续创建一个路径。
```

3）在源码目录下编写代码

4）`go install hello`编译源码，编译完成的可执行文件会保存在 $GOPATH/bin 目录下。

# 单例模式

## **懒汉式**

懒汉式就是创建对象时比较懒，先不急着创建对象，在需要加载配置文件的时候再去创建

```go
//懒汉，非线程安全
type Tool struct {
	values int
}

//私有变量
var instance *Tool

func GetInstance() *Tool {
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}
```

**Sync.Mutex 进行加锁保证线程安全**

```go
//线程安全的懒汉

type Tool struct {
	values int
}

var instance *Tool

//锁对象
var lock sync.Mutex

//加锁保证线程安全
func GetInstance() *Tool {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}
```

**锁加双重判断**

```go
//锁对象
var lock sync.Mutex

//第一次判断不加锁，第二次加锁保证线程安全，一旦对象建立后，获取对象就不用加锁了。
func GetInstance() *Tool {
    if instance == nil {
        lock.Lock()
        if instance == nil {
            instance = new(Tool)
        }
        lock.Unlock()
    }
    return instance
}
```

**sync.Once**

通过 sync.Once 来确保创建对象的方法只执行一次（ync.Once 内部本质上也是双重检查的方式，但在写法上会比自己写双重检查更简洁）

```go
var once sync.Once

func GetInstance() *Tool {
    once.Do(func() {
        instance = new(Tool)
    })
    return instance
}
```

## **饿汉式**

饿汉式则是在系统初始化的时候就已经把对象创建好了，需要用的时候直接拿过来用就好了。

Go语言饿汉式可以==使用 init 函数，也可以使用全局变量==。

```go
//使用init
type cfg struct {
}

var cfg *config

func init()  {
   cfg = new(config)
}

// NewConfig 提供获取实例的方法
func NewConfig() *config {
   return cfg
}
```

```go
//使用全局变量
type config struct {  
}

//全局变量
var cfg *config = new(config)

// NewConfig 提供获取实例的方法
func NewConfig() *config {
   return cfg
}
```

