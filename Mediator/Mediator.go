package main

import "fmt"

// 中介者接口
type Mediator interface {
	SendMessage(sender string, message string)
	Register(user *User)
}

// 具体中介者：聊天室
type ChatRoom struct {
	users map[string]*User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: make(map[string]*User),
	}
}

func (c *ChatRoom) Register(user *User) {
	c.users[user.name] = user
	user.mediator = c
}

func (c *ChatRoom) SendMessage(sender string, message string) {
	for name, user := range c.users {
		if name != sender {
			user.Receive(sender, message)
		}
	}
}

// 同事对象：用户
type User struct {
	name     string
	mediator Mediator
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) Send(message string) {
	fmt.Printf("[%s] 发送消息: %s\n", u.name, message)
	u.mediator.SendMessage(u.name, message)
}

func (u *User) Receive(from, message string) {
	fmt.Printf("[%s] 收到来自 [%s] 的消息: %s\n", u.name, from, message)
}

func main() {
	chatRoom := NewChatRoom()

	alice := NewUser("Alice")
	bob := NewUser("Bob")
	tom := NewUser("Tom")

	chatRoom.Register(alice)
	chatRoom.Register(bob)
	chatRoom.Register(tom)

	alice.Send("Hi, 大家好！")
	tom.Send("欢迎 Alice！")
}
