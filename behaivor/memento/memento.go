// Package memento @Author:冯铁城 [17615007230@163.com] 2025-09-26 10:30:40
package memento

import "fmt"

// LocationMemento 位置备忘录
type LocationMemento struct {
	Location string
}

// NewLocationMemento 创建位置备忘录
func NewLocationMemento(location string) *LocationMemento {
	return &LocationMemento{
		Location: location,
	}
}

// LocationMementoManager 位置备忘录管理器
type LocationMementoManager struct {
	mementoList map[int]*LocationMemento
}

// SetMemento 设置位置备忘录
func (l *LocationMementoManager) SetMemento(index int, memento *LocationMemento) {
	l.mementoList[index] = memento
}

// GetMemento 获取位置备忘录
func (l *LocationMementoManager) GetMemento(index int) *LocationMemento {
	return l.mementoList[index]
}

// NewLocationMementoManager 创建位置备忘录管理器
func NewLocationMementoManager() *LocationMementoManager {
	return &LocationMementoManager{
		mementoList: make(map[int]*LocationMemento),
	}
}

// Ike 艾克（英雄联盟英雄）
type Ike struct {
	Location string
}

// Move 移动
func (i *Ike) Move(location string) {
	i.Location = location
}

// ShowLocation 显示位置
func (i *Ike) ShowLocation() {
	fmt.Printf("Ike is in %s\n", i.Location)
}

// CreateLocationMemento 创建位置备忘录
func (i *Ike) CreateLocationMemento() *LocationMemento {
	return NewLocationMemento(i.Location)
}

// RestoreLocation 恢复位置
func (i *Ike) RestoreLocation(locationMemento *LocationMemento) {
	i.Location = locationMemento.Location
}

// NewIke 创建艾克
func NewIke() *Ike {
	return &Ike{
		Location: "泉水",
	}
}
