package main

import "fmt"

// === 需求：项目中统一用这个接口记录日志 ===
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// === 第三方日志库 A ===
type FancyLog struct{}

func (f *FancyLog) PrintInfo(info string) {
	fmt.Println("[INFO]", info)
}

func (f *FancyLog) PrintError(err string) {
	fmt.Println("[ERROR]", err)
}

// === 第三方日志库 B ===
type SimpleLog struct{}

func (s *SimpleLog) Log(level string, message string) {
	fmt.Printf("[%s] %s\n", level, message)
}

// === 适配器 A：适配 FancyLog 为 Logger 接口 ===
type FancyLogAdapter struct {
	fancy *FancyLog
}

func NewFancyLogAdapter(f *FancyLog) Logger {
	return &FancyLogAdapter{fancy: f}
}

func (a *FancyLogAdapter) Info(msg string) {
	a.fancy.PrintInfo(msg)
}

func (a *FancyLogAdapter) Error(msg string) {
	a.fancy.PrintError(msg)
}

// === 适配器 B：适配 SimpleLog 为 Logger 接口 ===
type SimpleLogAdapter struct {
	simple *SimpleLog
}

func NewSimpleLogAdapter(s *SimpleLog) Logger {
	return &SimpleLogAdapter{simple: s}
}

func (a *SimpleLogAdapter) Info(msg string) {
	a.simple.Log("INFO", msg)
}

func (a *SimpleLogAdapter) Error(msg string) {
	a.simple.Log("ERROR", msg)
}

// === 使用统一 Logger 接口 ===
func AppRun(logger Logger) {
	logger.Info("程序启动成功")
	logger.Error("发生了错误")
}

func main() {
	// 使用 FancyLog 适配器
	fancyLogger := NewFancyLogAdapter(&FancyLog{})
	AppRun(fancyLogger)

	// 使用 SimpleLog 适配器
	simpleLogger := NewSimpleLogAdapter(&SimpleLog{})
	AppRun(simpleLogger)
}
