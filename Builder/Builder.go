package main

import "fmt"

// 目标结构体：数据库连接配置
type DBConfig struct {
	host     string
	port     int
	user     string
	password string
	database string
	timeout  int
}

// 构造函数
func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

// 链式调用方法
func (c *DBConfig) Host(h string) *DBConfig {
	c.host = h
	return c
}

func (c *DBConfig) Port(p int) *DBConfig {
	c.port = p
	return c
}

func (c *DBConfig) User(u string) *DBConfig {
	c.user = u
	return c
}

func (c *DBConfig) Password(pwd string) *DBConfig {
	c.password = pwd
	return c
}

func (c *DBConfig) Database(db string) *DBConfig {
	c.database = db
	return c
}

func (c *DBConfig) Timeout(t int) *DBConfig {
	c.timeout = t
	return c
}

// 打印配置（模拟连接）
func (c *DBConfig) Connect() {
	fmt.Printf("Connecting to DB at %s:%d as user '%s' to database '%s' with timeout %ds\n",
		c.host, c.port, c.user, c.database, c.timeout)
}

func main() {
	cfg := NewDBConfig().
		Host("localhost").
		Port(3306).
		User("root").
		Password("123456").
		Database("test").
		Timeout(30)

	cfg.Connect()
}
