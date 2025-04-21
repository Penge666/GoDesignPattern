package main

import "fmt"

// 备忘录：保存状态
type Memento struct {
	content string
}

// 原发器：文本编辑器
type Editor struct {
	content string
}

func (e *Editor) Write(text string) {
	e.content += text
}

func (e *Editor) Save() Memento {
	return Memento{content: e.content}
}

func (e *Editor) Restore(m Memento) {
	e.content = m.content
}

func (e *Editor) Show() {
	fmt.Println("当前内容：", e.content)
}

// 管理者：历史记录管理
type History struct {
	snapshots []Memento
}

func (h *History) Push(m Memento) {
	h.snapshots = append(h.snapshots, m)
}

func (h *History) Pop() Memento {
	if len(h.snapshots) == 0 {
		return Memento{}
	}
	last := h.snapshots[len(h.snapshots)-1]
	h.snapshots = h.snapshots[:len(h.snapshots)-1]
	return last
}

func main() {
	editor := &Editor{}
	history := &History{}

	editor.Write("Hello, ")
	history.Push(editor.Save()) // 保存状态

	editor.Write("world!")
	history.Push(editor.Save()) // 再保存状态

	editor.Write(" GO语言！")
	editor.Show() // 当前内容：Hello, world! GO语言！

	// 撤销一次
	editor.Restore(history.Pop())
	editor.Show() // 当前内容：Hello, world!

	// 再撤销一次
	editor.Restore(history.Pop())
	editor.Show() // 当前内容：Hello,
}
