package main

import "fmt"

// 假设我们有一个菜单系统，它有 菜单（组合对象）和 菜单项（叶子对象）。
// 组件接口
type MenuComponent interface {
	Show() // 展示菜单或菜单项
}

// 菜单项：叶子节点
type MenuItem struct {
	name string
}

func (m *MenuItem) Show() {
	fmt.Println("Menu Item:", m.name)
}

// 菜单：组合节点，包含多个菜单项或子菜单
type Menu struct {
	name  string
	items []MenuComponent
}

// 添加菜单项或子菜单
func (m *Menu) Add(item MenuComponent) {
	m.items = append(m.items, item)
}

// 展示菜单及其内容
func (m *Menu) Show() {
	fmt.Println("Menu:", m.name)
	for _, item := range m.items {
		item.Show() // 递归展示所有子菜单和菜单项
	}
}
func main() {
	// 创建菜单项
	item1 := &MenuItem{name: "Burger"}
	item2 := &MenuItem{name: "Fries"}
	item3 := &MenuItem{name: "Coke"}

	// 创建菜单
	mainMenu := &Menu{name: "Main Menu"}
	drinksMenu := &Menu{name: "Drinks Menu"}

	// 向主菜单和子菜单添加菜单项
	mainMenu.Add(item1)
	mainMenu.Add(item2)
	drinksMenu.Add(item3)

	// 将子菜单添加到主菜单
	mainMenu.Add(drinksMenu)

	// 展示主菜单及其所有内容
	mainMenu.Show()
}
