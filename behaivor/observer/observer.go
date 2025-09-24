// Package observer @Author:冯铁城 [17615007230@163.com] 2025-09-24 17:48:37
package observer

import (
	"fmt"
	"sync"
)

// IPublisher 发布者接口
type IPublisher interface {
	AddSubscriber(name string, subscriber ISubscriber)
	RemoveSubscriber(name string)
	Notify(message string)
}

// Publisher 发布者结构体
type Publisher struct {
	mu          sync.RWMutex
	Subscribers map[string]ISubscriber
}

// AddSubscriber 添加订阅者
func (p *Publisher) AddSubscriber(name string, subscriber ISubscriber) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Subscribers[name] = subscriber
}

// RemoveSubscriber 移除订阅者
func (p *Publisher) RemoveSubscriber(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.Subscribers, name)
}

// Notify 通知
func (p *Publisher) Notify(message string) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, subscriber := range p.Subscribers {
		subscriber.Receive(message)
	}
}

// NewPublisher 创建发布者
func NewPublisher() *Publisher {
	return &Publisher{
		Subscribers: make(map[string]ISubscriber),
	}
}

// ISubscriber 订阅者接口
type ISubscriber interface {
	Receive(message string)
}

// Subscriber 订阅者结构体
type Subscriber struct {
	Name string
}

// Receive 接收
func (s *Subscriber) Receive(message string) {
	fmt.Println(s.Name, "received message:", message)
}

// NewSubscriber 创建订阅者
func NewSubscriber(name string) *Subscriber {
	return &Subscriber{
		Name: name,
	}
}
