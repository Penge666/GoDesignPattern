package main

import (
	"fmt"
	"sync"
)

// RedisConnection 模拟一个Redis连接
type RedisConnection struct {
	ID string
}

// 新建一个 Redis 连接
func NewRedisConnection(id string) *RedisConnection {
	return &RedisConnection{ID: id}
}

// RedisPool 享元池，管理多个共享的连接
type RedisPool struct {
	connections map[string]*RedisConnection
	mu          sync.Mutex
}

// 获取一个 Redis 连接，如果连接池中已经存在则复用，否则创建新的连接
func (p *RedisPool) GetConnection(id string) *RedisConnection {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 如果连接池中已有该连接，则直接返回
	if conn, exists := p.connections[id]; exists {
		fmt.Println("复用连接:", id)
		return conn
	}

	// 如果没有该连接，则新建并加入池中
	conn := NewRedisConnection(id)
	p.connections[id] = conn
	fmt.Println("新建连接:", id)
	return conn
}

// 初始化 RedisPool
func NewRedisPool() *RedisPool {
	return &RedisPool{
		connections: make(map[string]*RedisConnection),
	}
}

func main() {
	// 初始化 Redis 连接池
	pool := NewRedisPool()

	// 获取 Redis 连接
	conn1 := pool.GetConnection("redis://localhost:6379")
	fmt.Println("连接 ID:", conn1.ID)

	conn2 := pool.GetConnection("redis://localhost:6379")
	fmt.Println("连接 ID:", conn2.ID)

	conn3 := pool.GetConnection("redis://localhost:6380")
	fmt.Println("连接 ID:", conn3.ID)

	conn4 := pool.GetConnection("redis://localhost:6380")
	fmt.Println("连接 ID:", conn4.ID)
}
