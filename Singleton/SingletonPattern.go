package main

type Config struct {
	Value string
}

// 懒汉模式
// 方式1：sync.Once 本质上就是一种安全的双重检查锁机制。可以点击进去查看源码
//var (
//	instance *Config
//	once     sync.Once
//)
//
//func GetInstance() *Config {
//	once.Do(func() {
//		instance = &Config{Value: "Hello, Singleton!"}
//	})
//	return instance
//}

// 方式2：使用 sync.Mutex 实现
// var (
//
//	instance2 *Config
//	lock      sync.Mutex
//
// )
//
//	func GetInstanceWithLock() *Config {
//		lock.Lock()
//		defer lock.Unlock()
//		if instance2 == nil {
//			instance2 = &Config{Value: "Lock version"}
//		}
//		return instance2
//	}
//
// 方式3：饿汉模式

var instance4 = &Config{Value: "Eager singleton"}

func GetInstanceEager() *Config {
	return instance4
}

func main() {
	// 使用懒汉模式
	// 方式1：sync.Once
	//cfg1 := GetInstance()
	//cfg2 := GetInstance()
	//
	//fmt.Println(cfg1 == cfg2) // true，表示是同一个对象
	//fmt.Println(cfg1.Value)   // Hello, Singleton!

	// 方式2：使用锁机制
	//cfg3 := GetInstanceWithLock()
	//cfg4 := GetInstanceWithLock()
	//
	//println(cfg3 == cfg4) // true，表示是同一个对象
	//println(cfg3.Value)   // Lock version

	// 饿汉模式
	cfg5 := GetInstanceEager()
	cfg6 := GetInstanceEager()
	println(cfg5 == cfg6) // true，表示是同一个对象
	println(cfg5.Value)   // Eager singleton
}
