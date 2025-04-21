package main

import "fmt"

/**
你正在开发一个电商平台的订单处理系统，用户在平台下单后，系统会进行库存扣除、支付操作以及物流发货流程。为此，系统定义了几个核心模块：
- 库存模块（Inventory）：在用户下单后，系统需要从库存中扣除相应的商品数量。
- 支付模块（Pay）：用户支付成功后，系统需要执行支付操作，并记录支付结果。
- 物流模块（Logistics）：支付完成后，系统会将订单发送至用户，并触发物流发货操作。
*/
// 库存管理服务
type Inventory struct {
}

func (i *Inventory) Deduction() {
	fmt.Println("从库存中扣除商品")
}

// 支付服务
type Pay struct {
}

func (p *Pay) Pay() {
	fmt.Println("处理支付操作")
}

// 物流服务
type Logistics struct {
}

func (l *Logistics) SendOutGoods() {
	fmt.Println("发货")
}

// 订单管理
type Order struct {
	inventory *Inventory
	pay       *Pay
	logistics *Logistics
}

// 构造函数，用于创建新的订单实例
func NewOrder() *Order {
	return &Order{
		inventory: &Inventory{}, // 创建库存服务实例
		pay:       &Pay{},       // 创建支付服务实例
		logistics: &Logistics{}, // 创建物流服务实例
	}
}

// 处理下单操作
func (o *Order) Place() {
	o.inventory.Deduction()    // 扣除库存
	o.pay.Pay()                // 进行支付
	o.logistics.SendOutGoods() // 发货
}

func main() {
	// 创建订单实例
	order := NewOrder()

	// 执行下单流程
	order.Place()
}
