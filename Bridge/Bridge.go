package main

import "fmt"

/*
需求：
我们有不同的图形：如 圆形 和 矩形。
每个图形都可以使用不同的绘制方式：如 绘制到屏幕 或 绘制到文件。
*/

// ========== 绘制方式（实现部分） ==========

type DrawMethod interface {
	DrawShape(shapeName string)
}

// 在屏幕上绘制
type ScreenDraw struct{}

func (s *ScreenDraw) DrawShape(shapeName string) {
	fmt.Println("Drawing", shapeName, "on screen")
}

// 在文件中绘制
type FileDraw struct{}

func (f *FileDraw) DrawShape(shapeName string) {
	fmt.Println("Drawing", shapeName, "to file")
}

// ========== 图形（抽象部分） ==========

type Shape interface {
	Draw()
}

// 圆形
type Circle struct {
	drawMethod DrawMethod
}

func (c *Circle) Draw() {
	c.drawMethod.DrawShape("Circle")
}

// 矩形
type Rectangle struct {
	drawMethod DrawMethod
}

func (r *Rectangle) Draw() {
	r.drawMethod.DrawShape("Rectangle")
}
func main() {
	// 使用屏幕绘制方法
	screenDraw := &ScreenDraw{}

	// 创建圆形，使用屏幕绘制
	circle := &Circle{drawMethod: screenDraw}
	circle.Draw()

	// 创建矩形，使用屏幕绘制
	rectangle := &Rectangle{drawMethod: screenDraw}
	rectangle.Draw()

	// 使用文件绘制方法
	fileDraw := &FileDraw{}

	// 创建圆形，使用文件绘制
	circle = &Circle{drawMethod: fileDraw}
	circle.Draw()

	// 创建矩形，使用文件绘制
	rectangle = &Rectangle{drawMethod: fileDraw}
	rectangle.Draw()
}
