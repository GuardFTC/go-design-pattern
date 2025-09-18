// Package proxy @Author:冯铁城 [17615007230@163.com] 2025-09-18 10:43:22
package proxy

import "fmt"

// Star 明星
type Star interface {

	// Sing 唱歌
	Sing()

	// Businesses 商业活动
	Businesses()
}

// RealStar 真实明星
type RealStar struct {
	Name string
}

// Sing 唱歌
func (r *RealStar) Sing() {
	fmt.Printf("明星%s正在唱歌...\n", r.Name)
}

// Businesses 商业活动
func (r *RealStar) Businesses() {
	fmt.Printf("明星%s正在商业活动...\n", r.Name)
}

// Agent 经纪人
type Agent struct {
	Star
}

// Sing 唱歌
func (a *Agent) Sing() {
	fmt.Println("经纪人正在洽谈唱歌合同...")
	a.Star.Sing()
}

// Businesses 商业活动
func (a *Agent) Businesses() {
	fmt.Println("经纪人正在洽谈商业合同...")
	a.Star.Businesses()
}
