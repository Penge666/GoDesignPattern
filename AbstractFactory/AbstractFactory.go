package main

import "fmt"

// ========== 支付接口定义 ==========
type Pay interface {
	PayPage(price int64) string
}

// ========== 支付实现：支付宝 ==========
type AliPay struct{}

func (AliPay) PayPage(price int64) string {
	return fmt.Sprintf("跳转到【支付宝】支付页面，金额：%d 元", price)
}

// ========== 支付实现：微信 ==========
type WeiXinPay struct{}

func (WeiXinPay) PayPage(price int64) string {
	return fmt.Sprintf("跳转到【微信】支付页面，金额：%d 元", price)
}

// ========== 退款接口定义 ==========
type Refund interface {
	Refund(no string) error
}

// ========== 退款实现：支付宝 ==========
type AliRefund struct{}

func (AliRefund) Refund(no string) error {
	fmt.Printf("【支付宝】订单 %s 已退款成功\n", no)
	return nil
}

// ========== 退款实现：微信 ==========
type WeiXinRefund struct{}

func (WeiXinRefund) Refund(no string) error {
	fmt.Printf("【微信】订单 %s 已退款成功\n", no)
	return nil
}

// ========== 工厂接口 ==========
type PayFactory interface {
	CreatePay() Pay
	CreateRefund() Refund
}

// ========== 支付宝工厂 ==========
type AliPayFactory struct{}

func (AliPayFactory) CreatePay() Pay {
	return AliPay{}
}
func (AliPayFactory) CreateRefund() Refund {
	return AliRefund{}
}

// ========== 微信工厂 ==========
type WeiXinPayFactory struct{}

func (WeiXinPayFactory) CreatePay() Pay {
	return WeiXinPay{}
}
func (WeiXinPayFactory) CreateRefund() Refund {
	return WeiXinRefund{}
}

// ========== 模拟用户下单支付 & 退款 ==========
func main() {
	fmt.Println("用户选择使用支付宝支付：")
	aliFactory := AliPayFactory{}
	aliPay := aliFactory.CreatePay()
	fmt.Println(aliPay.PayPage(88))
	aliFactory.CreateRefund().Refund("ALI123456")

	fmt.Println("\n?? 用户选择使用微信支付：")
	wxFactory := WeiXinPayFactory{}
	wxPay := wxFactory.CreatePay()
	fmt.Println(wxPay.PayPage(66))
	wxFactory.CreateRefund().Refund("WX987654")
}
