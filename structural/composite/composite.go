// Package composite @Author:冯铁城 [17615007230@163.com] 2025-09-23 17:40:06
package composite

import "fmt"

// Component 组件接口
type Component interface {
	ShowDetail()
}

// Level1 一级组件结构体
type Level1 struct {
	Name string
}

// ShowDetail 显示详情
func (l *Level1) ShowDetail() {
	fmt.Printf("the level1 name is %s\n", l.Name)
}

// Level2 二级组件结构体
type Level2 struct {
	Name     string
	Children []Component
}

// ShowDetail 显示详情
func (l *Level2) ShowDetail() {
	fmt.Printf("the level2 name is %s\n", l.Name)
	for _, child := range l.Children {
		child.ShowDetail()
	}
}

// Level3 二级组件结构体
type Level3 struct {
	Name     string
	Children []Component
}

// ShowDetail 显示详情
func (l *Level3) ShowDetail() {
	fmt.Printf("the level3 name is %s\n", l.Name)
	for _, child := range l.Children {
		child.ShowDetail()
	}
}
