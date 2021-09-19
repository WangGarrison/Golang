/* URL短链接系统 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strings"
)

// 10进制-64进制对照表
var dec64Map = map[int]string{0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J",
	10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R", 18: "S", 19: "T", 20: "U", 21: "V",
	22: "W", 23: "X", 24: "Y", 25: "Z", 26: "a", 27: "b", 28: "c", 29: "d", 30: "e", 31: "f", 32: "g", 33: "h",
	34: "i", 35: "j", 36: "k", 37: "l", 38: "m", 39: "n", 40: "o", 41: "p", 42: "q", 43: "r", 44: "s", 45: "t",
	46: "u", 47: "v", 48: "w", 49: "x", 50: "y", 51: "z", 52: "0", 53: "1", 54: "2", 55: "3", 56: "4", 57: "5",
	58: "6", 59: "7", 60: "8", 61: "9", 62: "+", 63: "/"}

// 64进制-10进制对照表，即把上表中的键值对调换
var _64decMap = func() map[string]int {
	mp2 := map[string]int{}
	for k, v := range dec64Map {
		mp2[v] = k
	}
	return mp2
}()

// Link 数据库中的一条记录
type Link struct {  //表名默认是结构体的复数
	ID int  //名为`ID`的字段会默认作为表的主键
	Url string
}

// 短连接主机
var host = "http://localhost:8080"

// 长连接转短连接
func longToShort(long string, db *gorm.DB, rd * redis.Client) string {
	//先去查redis，若没找到，再去将长连接插入mysql，并把该长短网址映射添加到redis
	val, _ := rd.Get(long).Result()
	if val != "" {
		//刷新该映射的过期时间---------------------------
		return val
	}
	//将该长连接插入数据表
	lk := Link{Url : long}
	db.Create(&lk)
	// 获取该长连接对应的ID，并根据ID生成对应的短连接
	id := lk.ID
	if id == 0 {
		panic("长连接插入到mysql中失败")
		return ""
	}
	//将十进制id转为64进制字符串，128 => 20 => CA
	short := ""
	for id != 0 {
		short = dec64Map[id % 64] + short
		id /= 64
	}
	//组合短连接域名
	shortUrl := fmt.Sprintf("%s/%s",host,short)
	//将映射添加到redis中
	rd.Set(long, shortUrl, 0)
	return shortUrl
}

//  将sho域名重定向到lon
func redirect(sho string, lon string, g *gin.Engine) {
	// 域名重定向
	sho = strings.TrimPrefix(sho, host)  //去除前面的http://localhost:8080
	g.GET(sho, func(c *gin.Context) {  //指定用户执行sho请求时，重定向到lon
		c.Redirect(http.StatusMovedPermanently, lon)
	})
}

// 短连接转长连接
func shortToLong(short string, db *gorm.DB, rd *redis.Client) string {
	//先去查redis，若没找到再去通过短网址反推出其id，然后去查询后方sql，并把该短长网址映射添加到redis
	val, _ := rd.Get(short).Result()
	if val != "" {
		//刷新该映射的过期时间---------------------------
		return val
	}
	// 根据short反推出其ID，再根据ID获取其长连接
	// 将64进制字符串还原成十进制数字
	short = strings.TrimPrefix(short, host+"/")  //去除前面的http://localhost:8080
	id := 0
	le := len(short)
	sh := []byte(short)  //将字符串转换为切片数组
	for i := 0; i < le; i++ {
		id = id*64 + _64decMap[string(sh[i])] //sh[i]是个byte类型，map中一个字符用的是string，所以这里强转用string构造一下
	}
	//根据id去查数据库获得其对应的长连接
	var lk Link
	db.First(&lk, id)
	//将映射添加到redis中
	rd.Set(short, lk.Url,0)
	return lk.Url
}

func main() {
	//连接MySql： "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", "root:123456@(localhost)/shortlink?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Link{})  //自动迁移
	defer func() {  //主函数结束时，关闭db
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	//连接redis
	rd := redis.NewClient(&redis.Options {
		Addr: 		   "127.0.0.1:6379",
		Password:    "",
		DB:				0,
	})

	//创建一个默认的路由引擎
	g := gin.Default()
	gin.DisableConsoleColor()  //控制台颜色显示乱码，干脆直接禁用颜色

	//显示主界面请求
	g.LoadHTMLFiles("./html/mainPage.html")
	g.GET("./", func(c *gin.Context) {
		c.HTML(200, "mainPage.html", gin.H{})
	})

	//缩短网址请求，注意有提交数据，得用POST
	g.POST("shortenUrl", func(c *gin.Context) {
		longUrl := c.PostForm("longUrl")  //获取submit提交的表单中longUrl对应的value值
		shortUrl := longToShort(longUrl, db, rd)  //转换
		if shortUrl == "" {
			panic("shortUrl为空")
			return
		}
		redirect(shortUrl, longUrl, g)  //重定向
		c.String(200,  "%s缩短后的网址为：%s", longUrl, shortUrl)
	})

	//还原网址请求，注意有提交数据，得用POST
	g.POST("reductionUrl", func(c *gin.Context) {
		shortUrl := c.PostForm("shortUrl") //获取submit提交的表单中shortUrl对应的value值
		longUrl := shortToLong(shortUrl, db, rd)  //转换
		c.String(200, "%s还原后的网址为：%s", shortUrl, longUrl)
	})

	err = g.Run(":8080")
	if err != nil {
		fmt.Println("启动失败")
	}
}





