// Package singleton @Author:冯铁城 [17615007230@163.com] 2025-09-17 10:33:54
package singleton

import "sync"

// ISingleton 单例接口
type ISingleton interface {
	GetName() string
}

// singleton 单例结构体
type singleton struct {
	name string
}

// GetName 获取名称
func (s *singleton) GetName() string {
	return s.name
}

//--------------------------------饿汉式--------------------------------//

// Instance 单例实例
var eagerInstance *singleton = &singleton{
	name: "eager singleton",
}

// GetInstanceEager 获取单例实例
func GetInstanceEager() ISingleton {
	return eagerInstance
}

//--------------------------------懒汉式--------------------------------//

// once 锁
var once sync.Once

// lazyInstance 单例实例
var lazyInstance *singleton

// GetInstanceLazy 获取单例实例
func GetInstanceLazy() ISingleton {
	once.Do(func() {
		lazyInstance = &singleton{
			name: "lazy singleton",
		}
	})
	return lazyInstance
}
