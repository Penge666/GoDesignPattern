package main

import "fmt"

// Command 命令接口
type Command interface {
	Execute()
}

// Receiver：具体执行者
type TV struct{}

func (t *TV) On() {
	fmt.Println("电视开了")
}

func (t *TV) Off() {
	fmt.Println("电视关了")
}

// ConcreteCommand：打开命令
type TurnOnCommand struct {
	tv *TV
}

func (c *TurnOnCommand) Execute() {
	c.tv.On()
}

// ConcreteCommand：关闭命令
type TurnOffCommand struct {
	tv *TV
}

func (c *TurnOffCommand) Execute() {
	c.tv.Off()
}

// Invoker：命令调用者
type RemoteControl struct {
	commands []Command
}

func (r *RemoteControl) AddCommand(cmd Command) {
	r.commands = append(r.commands, cmd)
}

func (r *RemoteControl) ExecuteCommands() {
	for _, cmd := range r.commands {
		cmd.Execute()
	}
}

func main() {
	tv := &TV{}

	onCmd := &TurnOnCommand{tv: tv}
	offCmd := &TurnOffCommand{tv: tv}

	remote := &RemoteControl{}
	remote.AddCommand(onCmd)
	remote.AddCommand(offCmd)

	remote.ExecuteCommands()
}
