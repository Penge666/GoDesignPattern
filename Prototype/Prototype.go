package main

import "fmt"

// 原型接口
type Prototype interface {
	Clone() Prototype
}

// UserConfig 是一个复杂结构体
type UserConfig struct {
	Theme     string
	FontSize  int
	Bookmarks []string
}

// Clone 实现 Prototype 接口
func (c *UserConfig) Clone() Prototype {
	// 深拷贝 Bookmarks
	bookmarksCopy := make([]string, len(c.Bookmarks))
	copy(bookmarksCopy, c.Bookmarks)

	return &UserConfig{
		Theme:     c.Theme,
		FontSize:  c.FontSize,
		Bookmarks: bookmarksCopy,
	}
}

func main() {
	// 缓存中已有一个默认配置
	defaultConfig := &UserConfig{
		Theme:     "dark",
		FontSize:  14,
		Bookmarks: []string{"https://golang.org", "https://github.com"},
	}

	// 克隆出一个新用户的配置
	user1Config := defaultConfig.Clone().(*UserConfig)

	// 修改新用户配置，不影响默认缓存
	user1Config.FontSize = 18
	user1Config.Bookmarks = append(user1Config.Bookmarks, "https://chat.openai.com")

	fmt.Println("默认配置：", defaultConfig)
	fmt.Println("用户1配置：", user1Config)
}
