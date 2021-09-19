# net/http包实现Web开发

一个请求对应一个响应

<img align='left' src="img/Go8%EF%BC%9A%E9%95%BF%E8%BF%9E%E6%8E%A5%E8%BD%AC%E7%9F%AD%E8%BF%9E%E6%8E%A5%E9%A1%B9%E7%9B%AE.img/image-20210715125442429.png" alt="image-20210715125442429" style="zoom:50%;" />

使用go自带的net/http包简单启动一个web服务器：

```go
package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_,_ = fmt.Fprintln(w, "<h1>Hello</h1>")

}

func main() {
	http.HandleFunc("/hello", sayHello) //一次hello的请求，对应一次hello的响应
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve failed, err:%v\n", err)
	}
}
```

![image-20210715132126254](img/Go8%EF%BC%9A%E9%95%BF%E8%BF%9E%E6%8E%A5%E8%BD%AC%E7%9F%AD%E8%BF%9E%E6%8E%A5%E9%A1%B9%E7%9B%AE.img/image-20210715132126254.png)

也可以把html语句`<h1>Hello</h1>`写在文本里：

![image-20210715132756049](img/Go8%EF%BC%9A%E9%95%BF%E8%BF%9E%E6%8E%A5%E8%BD%AC%E7%9F%AD%E8%BF%9E%E6%8E%A5%E9%A1%B9%E7%9B%AE.img/image-20210715132756049.png)

hello.txt如下：

![image-20210715133308025](img/Go8%EF%BC%9A%E9%95%BF%E8%BF%9E%E6%8E%A5%E8%BD%AC%E7%9F%AD%E8%BF%9E%E6%8E%A5%E9%A1%B9%E7%9B%AE.img/image-20210715133308025.png)

这时候启动服务器，显示如下：

<img src="img/Go8%EF%BC%9A%E9%95%BF%E8%BF%9E%E6%8E%A5%E8%BD%AC%E7%9F%AD%E8%BF%9E%E6%8E%A5%E9%A1%B9%E7%9B%AE.img/image-20210715133346548.png" alt="image-20210715133346548" style="zoom:50%;" />

注意：不结束程序，直接修改hello.txt文件，刷新浏览器就能直接看到效果

# Gin

> Gin：Go最流行的Web开发框架，[github地址](https://github.com/gin-gonic/gin)

使用Gin框架搭建简易的Web服务器示例：

```go
package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{     //返回给前端一个json格式的消息
		"message" : "Hello",
	})
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 指定用户使用GET请求访问/hello时，执行sayHello这个函数
	r.GET("/hello", sayHello)

	// 启动服务
    r.Run()  //默认在8080端口，也可以指定端口r.Run(":9090")
}
```

# RESTful风格的API

REST与技术无关，代表的是一种软件架构风格，REST是Representational State Transfer的简称，中文翻译为“表征状态转移”或“表现层状态转化”。

简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同的动作。

- `GET`用来获取资源
- `POST`用来新建资源
- `PUT`用来更新资源
- `DELETE`用来删除资源。

==即用4个http请求的动词来表示4个动作==（4个请求方法返回的是json格式的数据）

只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

例如，我们现在要编写一个管理书籍的系统，我们可以查询对一本书进行查询、创建、更新和删除等操作，我们在编写程序的时候就要设计客户端浏览器与我们Web服务端交互的方式和路径。按照经验我们通常会设计成如下模式：

| 请求方法 |     URL      |     含义     |
| :------: | :----------: | :----------: |
|   GET    |    /book     | 查询书籍信息 |
|   POST   | /create_book | 创建书籍记录 |
|   POST   | /update_book | 更新书籍信息 |
|   POST   | /delete_book | 删除书籍信息 |

同样的需求我们按照RESTful API设计如下：

| 请求方法 |  URL  |     含义     |
| :------: | :---: | :----------: |
|   GET    | /book | 查询书籍信息 |
|   POST   | /book | 创建书籍记录 |
|   PUT    | /book | 更新书籍信息 |
|  DELETE  | /book | 删除书籍信息 |

Gin框架支持开发RESTful API的开发。

![image-20210716095753951](img/Go8%EF%BC%9AGin%E3%80%81GORM.img/image-20210716095753951.png)

**使用postman工具可以测试接口：**

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{     //返回给前端一个json格式的消息
		"message" : "Hello",
	})
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 指定用户使用GET请求访问/hello时，执行sayHello这个函数
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"GET",
		})
	})
	
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message":"POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message":"PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message":"DELETE",
		})
	})

	// 启动服务
	r.Run()
}
```

<img align='left' src="img/Go8%EF%BC%9AGin%E3%80%81GORM.img/image-20210716102626434.png" alt="image-20210716102626434" style="zoom:50%;" />

# Gin渲染、http/template

## Go模板引擎

模板可以理解为事先定义好的HTML文档文件，模板渲染的作用机制可以简单理解为文本替换操作–==使用相应的数据去替换HTML文档中事先准备好的标记==。

`html/template`包实现了数据驱动的模板，用于生成可防止代码注入的安全的HTML内容。它提供了和`text/template`包相同的接口，Go语言中输出HTML的场景都应使用`html/template`这个包。

Go语言内置了文本模板引擎`text/template`和用于HTML文档的`html/template`。它们的作用机制可以简单归纳如下：

1. 模板文件通常定义为`.tmpl`和`.tpl`为后缀（也可以使用其他的后缀），必须使用`UTF8`编码。
2. 模板文件中使用`{{`和`}}`包裹和标识需要传入的数据。
3. 传给模板这样的数据就可以通过点号（`.`）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
4. 除`{{`和`}}`包裹的内容外，其他内容均不做修改原样输出。

**模板引擎的使用：**

- 定义模板文件
- 解析模板文件
- 模板渲染 

**使用示例：**

```go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	// 3.渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("return template failed, err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err: %v", err)
		return
	}
}
```

hello.tmpl

```go
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>Hello</title>
</head>
<body>
    <p>Hello {{ . }}</p>
</body>
</html>
```

![image-20210716110636946](img/Go8%EF%BC%9AGin%E3%80%81GORM.img/image-20210716110636946.png)

## T模板语法

![image-20210716111804912](img/Go8%EF%BC%9AGin%E3%80%81GORM.img/image-20210716111804912.png)





 



















# GORM

go操作mysql的orm

```go
db.Create(&User{
  Name:     "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// INSERT INTO `users` (`name`,`location`) VALUES "jinzhu",ST_PointFromText("POINT(100 100)"))
```

