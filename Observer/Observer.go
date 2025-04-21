package main

import "fmt"

// 观察者接口
type Observer interface {
	Update(config string)
}

// 发布者（主题）
type ConfigCenter struct {
	observers []Observer
}

func (c *ConfigCenter) Register(o Observer) {
	c.observers = append(c.observers, o)
}

func (c *ConfigCenter) Notify(config string) {
	for _, o := range c.observers {
		o.Update(config)
	}
}

// 观察者实现A
type Logger struct{}

func (l Logger) Update(config string) {
	fmt.Println("Logger 收到配置更新：", config)
}

// 观察者实现B
type Metrics struct{}

func (m Metrics) Update(config string) {
	fmt.Println("Metrics 收到配置更新：", config)
}

func main() {
	center := &ConfigCenter{}

	center.Register(Logger{})
	center.Register(Metrics{})

	center.Notify("log_level=debug")
}
