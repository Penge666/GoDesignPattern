package main

import "fmt"

// 表达式接口
type Expression interface {
	Interpret() int
}

// 数字表达式
type Number struct {
	value int
}

func (n Number) Interpret() int {
	return n.value
}

// 加法表达式
type Add struct {
	left, right Expression
}

func (a Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// 减法表达式
type Subtract struct {
	left, right Expression
}

func (s Subtract) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

func main() {
	// 表达式: 5 + 3 - 2

	// 构建语法树 (5 + 3)
	add := Add{
		left:  Number{5},
		right: Number{3},
	}

	// 构建最终表达式 ((5 + 3) - 2)
	expression := Subtract{
		left:  add,
		right: Number{2},
	}

	// 解释执行
	result := expression.Interpret()
	fmt.Println("结果:", result) // 输出: 6
}
