// Package visitor @Author:冯铁城 [17615007230@163.com] 2025-09-25 10:22:42
package visitor

import "fmt"

// BasketballPlayer 篮球球员接口
type BasketballPlayer interface {
	Accept(visitor IVisitor)
}

// PG 控球后卫
type PG struct {
}

// GetPosition 获取位置
func (p *PG) GetPosition() string {
	return "PG"
}

// Accept 接受访问者
func (p *PG) Accept(visitor IVisitor) {
	visitor.VisitPG(p)
}

// SG 得分后卫
type SG struct {
}

// GetPosition 获取位置
func (s *SG) GetPosition() string {
	return "SG"
}

// Accept 接受访问者
func (s *SG) Accept(visitor IVisitor) {
	visitor.VisitSG(s)
}

// SF 小前锋
type SF struct {
}

// GetPosition 获取位置
func (s *SF) GetPosition() string {
	return "SF"
}

// Accept 接受访问者
func (s *SF) Accept(visitor IVisitor) {
	visitor.VisitSF(s)
}

// PF 大前锋
type PF struct {
}

// GetPosition 获取位置
func (p *PF) GetPosition() string {
	return "PF"
}

// Accept 接受访问者
func (p *PF) Accept(visitor IVisitor) {
	visitor.VisitPF(p)
}

// C 中锋
type C struct {
}

// GetPosition 获取位置
func (c *C) GetPosition() string {
	return "C"
}

// Accept 接受访问者
func (c *C) Accept(visitor IVisitor) {
	visitor.VisitC(c)
}

// IVisitor 访问者接口
type IVisitor interface {
	VisitPG(player *PG)
	VisitSG(player *SG)
	VisitSF(player *SF)
	VisitPF(player *PF)
	VisitC(player *C)
}

// BodyTrain 身体训练
type BodyTrain struct{}

func (b *BodyTrain) VisitPG(player *PG) {
	fmt.Println(player.GetPosition(), " 正在身体训练")
}

func (b *BodyTrain) VisitSG(player *SG) {
	fmt.Println(player.GetPosition(), " 正在身体训练")
}

func (b *BodyTrain) VisitSF(player *SF) {
	fmt.Println(player.GetPosition(), " 正在身体训练")
}

func (b *BodyTrain) VisitPF(player *PF) {
	fmt.Println(player.GetPosition(), " 正在身体训练")
}

func (b *BodyTrain) VisitC(player *C) {
	fmt.Println(player.GetPosition(), " 正在身体训练")
}

// ShootTrain 投篮训练
type ShootTrain struct{}

func (s *ShootTrain) VisitPG(player *PG) {
	fmt.Println(player.GetPosition(), " 正在投篮训练")
}

func (s *ShootTrain) VisitSG(player *SG) {
	fmt.Println(player.GetPosition(), " 正在投篮训练")
}

func (s *ShootTrain) VisitSF(player *SF) {
	fmt.Println(player.GetPosition(), " 正在投篮训练")
}

func (s *ShootTrain) VisitPF(player *PF) {
	fmt.Println(player.GetPosition(), " 正在投篮训练")
}

func (s *ShootTrain) VisitC(player *C) {
	fmt.Println(player.GetPosition(), " 正在投篮训练")
}

// SkillTrain 技巧训练
type SkillTrain struct{}

func (s *SkillTrain) VisitPG(player *PG) {
	fmt.Println(player.GetPosition(), " 正在技巧训练")
}

func (s *SkillTrain) VisitSG(player *SG) {
	fmt.Println(player.GetPosition(), " 正在技巧训练")
}

func (s *SkillTrain) VisitSF(player *SF) {
	fmt.Println(player.GetPosition(), " 正在技巧训练")
}

func (s *SkillTrain) VisitPF(player *PF) {
	fmt.Println(player.GetPosition(), " 正在技巧训练")
}

func (s *SkillTrain) VisitC(player *C) {
	fmt.Println(player.GetPosition(), " 正在技巧训练")
}
