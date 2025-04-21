package main

import "fmt"

// Beverage 抽象类，定义了制作饮品的接口
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
	PrepareRecipe() // 每个具体饮品类型实现此方法
}

// Coffee 咖啡类
type Coffee struct{}

func (Coffee) BoilWater() {
	fmt.Println("Boiling water for coffee")
}

func (Coffee) Brew() {
	fmt.Println("Brewing coffee grounds")
}

func (Coffee) PourInCup() {
	fmt.Println("Pouring coffee into cup")
}

func (Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk to coffee")
}
func (Coffee) PrepareRecipe() {
	// 具体实现
	fmt.Println("Preparing coffee recipe")
}

// Tea 茶类
type Tea struct{}

func (Tea) BoilWater() {
	fmt.Println("Boiling water for tea")
}

func (Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (Tea) PourInCup() {
	fmt.Println("Pouring tea into cup")
}

func (Tea) AddCondiments() {
	fmt.Println("Adding lemon to tea")
}

func (Tea) PrepareRecipe() {
	fmt.Println("Preparing tea recipe")
}

// PrepareRecipe 定义了制作饮品的骨架
func PrepareRecipe(beverage Beverage) {
	beverage.BoilWater()
	beverage.Brew()
	beverage.PourInCup()
	beverage.AddCondiments()
}

func main() {
	// 创建并使用 Coffee
	coffee := Coffee{}
	tea := Tea{}
	PrepareRecipe(coffee)
	PrepareRecipe(tea)

}
