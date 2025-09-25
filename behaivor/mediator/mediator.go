// Package mediator @Author:冯铁城 [17615007230@163.com] 2025-09-25 11:16:09
package mediator

import "fmt"

// IMediator 中介者接口
type IMediator interface {
	Notify(colleague IColleague, message string)
}

// IColleague 同事接口
type IColleague interface {
	SetMediator(mediator IMediator)
}

// Boy 男同事结构体
type Boy struct {
	mediator IMediator
}

// SetMediator 设置中介者
func (b *Boy) SetMediator(mediator IMediator) {
	b.mediator = mediator
}

// FindGirl 找妹子
func (b *Boy) FindGirl() {
	b.mediator.Notify(b, "find a girl")
}

// PlayWithGirl 和妹子一起玩
func (b *Boy) PlayWithGirl() {
	fmt.Printf("Boy play with a girl\n")
}

// NewBoy 创建男同事
func NewBoy(mediator IMediator) *Boy {
	b := &Boy{}
	b.SetMediator(mediator)
	return b
}

// Girl 女同事结构体
type Girl struct {
	mediator IMediator
}

// SetMediator 设置中介者
func (g *Girl) SetMediator(mediator IMediator) {
	g.mediator = mediator
}

// FindBoy 找男生
func (g *Girl) FindBoy() {
	g.mediator.Notify(g, "find a boy")
}

// PlayWithBoy 和男生一起玩
func (g *Girl) PlayWithBoy() {
	fmt.Printf("Girl play with a boy\n")
}

// NewGirl 创建女同事
func NewGirl(mediator IMediator) *Girl {
	g := &Girl{}
	g.SetMediator(mediator)
	return g
}

// Matchmaker 媒婆
type Matchmaker struct {
	Boy  *Boy
	Girl *Girl
}

// Notify 通知
func (m *Matchmaker) Notify(colleague IColleague, message string) {

	//.1.判断对象是否为空
	if m.Boy == nil || m.Girl == nil {
		return
	}

	//2.根据同事类型以及消息进行匹配
	switch colleague.(type) {
	case *Boy:
		if message == "find a girl" {
			fmt.Printf("Matchmaker find a girl\n")
			m.Girl.PlayWithBoy()
		}
	case *Girl:
		if message == "find a boy" {
			fmt.Printf("Matchmaker find a boy\n")
			m.Boy.PlayWithGirl()
		}
	}
}

// SetBoy 设置男同事
func (m *Matchmaker) SetBoy(boy *Boy) {
	m.Boy = boy
}

// SetGirl 添加女同事
func (m *Matchmaker) SetGirl(girl *Girl) {
	m.Girl = girl
}

// NewMatchmaker 创建媒婆
func NewMatchmaker() *Matchmaker {
	return &Matchmaker{}
}
