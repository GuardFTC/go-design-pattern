// Package template_method @Author:冯铁城 [17615007230@163.com] 2025-09-26 11:14:44
package template_method

import "fmt"

// ICooker 抽象烹饪接口
type ICooker interface {
	MakeTofuPudding()
	AddSeasoning()
}

// AbstractCooker 抽象烹饪结构体 (golang当中，结构体内嵌接口，能够实现类似于抽象类的功能)
type AbstractCooker struct {
	ICooker
}

// MakeTofuPudding 模版方法 制作豆腐脑
func (a *AbstractCooker) MakeTofuPudding() {
	fmt.Printf("开始做豆腐脑...\n")
	fmt.Printf("添加豆腐\n")
	fmt.Printf("添加水\n")
	fmt.Printf("等待豆腐成型\n")
	a.ICooker.AddSeasoning()
}

// NorthCooker 北方烹饪结构体 (golang当中，结构体内嵌结构体，能够实现类似于继承的功能)
type NorthCooker struct {
	AbstractCooker
}

// MakeTofuPudding 制作豆腐脑
func (n *NorthCooker) MakeTofuPudding() {
	n.AbstractCooker.MakeTofuPudding()
}

// AddSeasoning 添加调料
func (n *NorthCooker) AddSeasoning() {
	fmt.Printf("添加咸味卤子\n")
}

// NewNorthCooker 创建北方烹饪结构体
func NewNorthCooker() *NorthCooker {
	n := &NorthCooker{}
	n.AbstractCooker.ICooker = n
	return n
}

// SouthCooker 南方烹饪结构体
type SouthCooker struct {
	AbstractCooker
}

// AddSeasoning 添加调料
func (s *SouthCooker) AddSeasoning() {
	fmt.Printf("添加白糖\n")
}

// MakeTofuPudding 制作豆腐脑
func (s *SouthCooker) MakeTofuPudding() {
	s.AbstractCooker.MakeTofuPudding()
}

// NewSouthCooker 创建南方烹饪结构体
func NewSouthCooker() *SouthCooker {
	s := &SouthCooker{}
	s.AbstractCooker.ICooker = s
	return s
}
