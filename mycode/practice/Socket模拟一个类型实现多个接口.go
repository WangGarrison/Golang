package main

import "io"

//Socket结构体
type Socket struct {
}

//Writer接口
type Writer interface {
	Write(p []byte) (n int, err error)
}

//Closer接口
type Closer interface {
	Close() error
}

//接收器Socket的Write方法，该方法实现了Writer接口
func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

//Socket的Close方法，该方法实现了Closer接口
func (s *Socket) Close() error {
	return nil
}

/* ------------------------------------------------------------- */
/* 在代码中使用 Socket 结构实现的 Writer 接口和 Closer 接口代码如下：*/
/* ------------------------------------------------------------- */

func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

func usingCloser(closer io.Closer) {
	closer.Close()
}

func main1() {
	s := new(Socket)
	usingWriter(s) //将对象实例赋值给接口：要求对象实现了接口的所有方法
	usingCloser(s)
}
