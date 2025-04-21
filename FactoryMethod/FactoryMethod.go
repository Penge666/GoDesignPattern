package main

import "fmt"

// 消息发送接口
type Messenger interface {
	Send(to string, content string) string
}

// 具体消息类型
type SMS struct{}

func (SMS) Send(to, content string) string {
	return fmt.Sprintf("发送短信到 %s: %s", to, content)
}

type Email struct{}

func (Email) Send(to, content string) string {
	return fmt.Sprintf("发送邮件到 %s: %s", to, content)
}

type WeChat struct{}

func (WeChat) Send(to, content string) string {
	return fmt.Sprintf("发送微信消息到 %s: %s", to, content)
}

// 工厂接口
type MessengerFactory interface {
	CreateMessenger() Messenger
}

// 每种消息一个工厂
type SMSFactory struct{}

func (SMSFactory) CreateMessenger() Messenger {
	return SMS{}
}

type EmailFactory struct{}

func (EmailFactory) CreateMessenger() Messenger {
	return Email{}
}

type WeChatFactory struct{}

func (WeChatFactory) CreateMessenger() Messenger {
	return WeChat{}
}

func main() {
	smsFactory := SMSFactory{}
	sms := smsFactory.CreateMessenger()
	fmt.Println(sms.Send("13800138000", "您的验证码是 123456"))

	emailFactory := EmailFactory{}
	email := emailFactory.CreateMessenger()
	fmt.Println(email.Send("test@example.com", "欢迎注册"))

	wechatFactory := WeChatFactory{}
	wechat := wechatFactory.CreateMessenger()
	fmt.Println(wechat.Send("wxid_abc123", "你有一条新消息"))
}
