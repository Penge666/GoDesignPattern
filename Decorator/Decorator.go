package main

import (
	"fmt"
	"time"
)

type ReqI interface {
	Handler(url string) string
}

type Req struct{}

func (r Req) Handler(url string) string {
	fmt.Println("请求 " + url)
	time.Sleep(100 * time.Millisecond)
	return ""
}

type LogReqDecorator struct {
	req ReqI
}

func (l LogReqDecorator) Handler(url string) string {
	fmt.Println("日志记录前")
	res := l.req.Handler(url)
	fmt.Println("日志记录后")
	return res
}

type MonitorDecorator struct {
	req ReqI
}

func (m MonitorDecorator) Handler(url string) string {
	start := time.Now()
	res := m.req.Handler(url)
	fmt.Println("耗时", time.Since(start))
	return res
}

func main() {
	req := Req{}
	req1 := LogReqDecorator{req} // 应用日志装饰器
	req1.Handler("baidu1.com")
	req2 := MonitorDecorator{req} // 应用监控装饰器
	req2.Handler("baidu2.com")
}
