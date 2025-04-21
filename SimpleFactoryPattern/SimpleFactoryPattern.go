package main

import "fmt"

// 消息发送接口
type Messenger interface {
	Send(to string, content string) string
}

// 具体消息类型：短信 和 邮件
type SMS struct{}

func (sms *SMS) Send(to, content string) string {
	return fmt.Sprintf("发送短信到 %s: %s", to, content)
}

type Email struct{}

func (Email) Send(to, content string) string {
	return fmt.Sprintf("发送邮件到 %s: %s", to, content)
}

// 消息类型枚举
type MessageType int8

const (
	SMSMessage   MessageType = 1
	EmailMessage MessageType = 2
)

// 简单工厂
func NewMessenger(msgType MessageType) Messenger {
	switch msgType {
	case SMSMessage:
		return &SMS{} // 假设：这个方法里面会修改结构体字段状态
	case EmailMessage:
		return Email{} // 假设：这个方法里面不会修改结构体字段状态
	default:
		return nil
	}
}

func main() {
	msg := NewMessenger(SMSMessage)
	fmt.Println(msg.Send("13800138000", "您好，验证码是 123456"))

	msg = NewMessenger(EmailMessage)
	fmt.Println(msg.Send("test@example.com", "欢迎注册本服务"))
}
