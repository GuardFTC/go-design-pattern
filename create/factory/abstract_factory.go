// Package factory @Author:冯铁城 [17615007230@163.com] 2025-09-26 17:40:57
package factory

import "fmt"

// ICar 汽车接口
type ICar interface {
	Run()
}

// BMW 宝马
type BMW struct{}

// Run 启动
func (b *BMW) Run() {
	fmt.Printf("BMW is running\n")
}

// NewBMW 创建宝马
func NewBMW() *BMW {
	return &BMW{}
}

// Audi 奥迪
type Audi struct{}

// Run 启动
func (a *Audi) Run() {
	fmt.Printf("Audi is running\n")
}

// NewAudi 创建奥迪
func NewAudi() *Audi {
	return &Audi{}
}

// IEngine 引擎接口
type IEngine interface {
	Start()
}

// BMWEngine 宝马引擎
type BMWEngine struct{}

// Start 启动
func (b *BMWEngine) Start() {
	fmt.Printf("BMW engine is started\n")
}

// NewBMWEngine 创建宝马引擎
func NewBMWEngine() *BMWEngine {
	return &BMWEngine{}
}

// AudiEngine 奥迪引擎
type AudiEngine struct{}

// Start 启动
func (a *AudiEngine) Start() {
	fmt.Printf("Audi engine is started\n")
}

// NewAudiEngine 创建奥迪引擎
func NewAudiEngine() *AudiEngine {
	return &AudiEngine{}
}

// ITire 轮胎接口
type ITire interface {
	DisplayTread()
}

// BMWTire 宝马轮胎
type BMWTire struct{}

// DisplayTread 显示 tire 状态
func (b *BMWTire) DisplayTread() {
	fmt.Printf("BMW tire tread is <<<\n")
}

// NewBMWTire 创建宝马轮胎
func NewBMWTire() *BMWTire {
	return &BMWTire{}
}

// AudiTire 奥迪轮胎
type AudiTire struct{}

// DisplayTread 显示 tire 状态
func (a *AudiTire) DisplayTread() {
	fmt.Printf("Audi tire tread is >>>\n")
}

// NewAudiTire 创建奥迪轮胎
func NewAudiTire() *AudiTire {
	return &AudiTire{}
}

// ICarFactory 汽车工厂接口
type ICarFactory interface {
	CreateCar() ICar
	CreateEngine() IEngine
	CreateTire() ITire
}

// BMWCarFactory 宝马工厂
type BMWCarFactory struct{}

// CreateCar 创建宝马车
func (b *BMWCarFactory) CreateCar() ICar {
	return NewBMW()
}

// CreateEngine 创建宝马发动机
func (b *BMWCarFactory) CreateEngine() IEngine {
	return NewBMWEngine()
}

// CreateTire 创建宝马轮胎
func (b *BMWCarFactory) CreateTire() ITire {
	return NewBMWTire()
}

// NewBMWCarFactory 创建宝马工厂
func NewBMWCarFactory() ICarFactory {
	return &BMWCarFactory{}
}

// AudiCarFactory 奥迪工厂
type AudiCarFactory struct{}

// CreateCar 创建奥迪车
func (a *AudiCarFactory) CreateCar() ICar {
	return NewAudi()
}

// CreateEngine 创建奥迪引擎
func (a *AudiCarFactory) CreateEngine() IEngine {
	return NewAudiEngine()
}

// CreateTire 创建奥迪轮胎
func (a *AudiCarFactory) CreateTire() ITire {
	return NewAudiTire()
}

// NewAudiCarFactory 创建奥迪工厂
func NewAudiCarFactory() ICarFactory {
	return &AudiCarFactory{}
}
