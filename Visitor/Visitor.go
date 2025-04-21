package main

import "fmt"

// Expression 是所有表达式的接口
type Expression interface {
	Accept(visitor Visitor)
}

// Visitor 是访问者接口，定义了访问不同节点的操作
type Visitor interface {
	VisitAddExpression(add *AddExpression)
	VisitSubExpression(sub *SubExpression)
}

// AddExpression 是加法表达式类型
type AddExpression struct {
	Left  Expression
	Right Expression
}

// SubExpression 是减法表达式类型
type SubExpression struct {
	Left  Expression
	Right Expression
}

// Implement Accept 方法，让每个表达式接受访问者
func (a *AddExpression) Accept(visitor Visitor) {
	visitor.VisitAddExpression(a)
}

func (s *SubExpression) Accept(visitor Visitor) {
	visitor.VisitSubExpression(s)
}

// ConcreteVisitor 实现了 Visitor 接口，定义了对不同类型节点的具体操作
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitAddExpression(add *AddExpression) {
	fmt.Println("处理加法节点:", add.Left, "+", add.Right)
}

func (v *ConcreteVisitor) VisitSubExpression(sub *SubExpression) {
	fmt.Println("处理减法节点:", sub.Left, "-", sub.Right)
}

func main() {
	// 创建一个简单的表达式树
	expr := &AddExpression{
		Left: &SubExpression{
			Left:  &AddExpression{Left: nil, Right: nil}, // 简单的叶子节点
			Right: nil,
		},
		Right: nil,
	}

	// 创建访问者并访问表达式树
	visitor := &ConcreteVisitor{}
	expr.Accept(visitor)
}
