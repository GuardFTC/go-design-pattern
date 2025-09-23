// Package facade @Author:冯铁城 [17615007230@163.com] 2025-09-23 16:02:43
package facade

import "fmt"

// LightSystem 子系统1：灯光控制系统
type LightSystem struct{}

func (l *LightSystem) TurnOn() {
	fmt.Println("灯光已打开")
}

func (l *LightSystem) TurnOff() {
	fmt.Println("灯光已关闭")
}

func (l *LightSystem) SetBrightness(level int) {
	fmt.Printf("灯光亮度已设置为 %d%%\n", level)
}

// AirConditioner 子系统2：空调控制系统
type AirConditioner struct{}

func (a *AirConditioner) TurnOn() {
	fmt.Println("空调已打开")
}

func (a *AirConditioner) TurnOff() {
	fmt.Println("空调已关闭")
}

func (a *AirConditioner) SetTemperature(temp int) {
	fmt.Printf("空调温度已设置为 %d°C\n", temp)
}

// AudioSystem 子系统3：音响控制系统
type AudioSystem struct{}

func (a *AudioSystem) TurnOn() {
	fmt.Println("音响已打开")
}

func (a *AudioSystem) TurnOff() {
	fmt.Println("音响已关闭")
}

func (a *AudioSystem) SetVolume(volume int) {
	fmt.Printf("音响音量已设置为 %d\n", volume)
}

// SmartHomeFacade 外观类：智能家居控制器
type SmartHomeFacade struct {
	lightSystem    *LightSystem
	airConditioner *AirConditioner
	audioSystem    *AudioSystem
}

// NewSmartHomeFacade 创建智能家居控制器
func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		lightSystem:    &LightSystem{},
		airConditioner: &AirConditioner{},
		audioSystem:    &AudioSystem{},
	}
}

// WakeUpMode 起床模式
func (s *SmartHomeFacade) WakeUpMode() {
	fmt.Println("========== 起床模式 ==========")
	s.lightSystem.TurnOn()
	s.lightSystem.SetBrightness(50)
	s.airConditioner.TurnOn()
	s.airConditioner.SetTemperature(22)
	s.audioSystem.TurnOn()
	s.audioSystem.SetVolume(30)
	fmt.Println("起床模式已启动！")
	fmt.Println()
}

// SleepMode 睡眠模式
func (s *SmartHomeFacade) SleepMode() {
	fmt.Println("========== 睡眠模式 ==========")
	s.lightSystem.TurnOff()
	s.airConditioner.SetTemperature(26)
	s.audioSystem.TurnOff()
	fmt.Println("睡眠模式已启动！")
	fmt.Println()
}

// PartyMode 聚会模式
func (s *SmartHomeFacade) PartyMode() {
	fmt.Println("========== 聚会模式 ==========")
	s.lightSystem.TurnOn()
	s.lightSystem.SetBrightness(100)
	s.airConditioner.TurnOn()
	s.airConditioner.SetTemperature(20)
	s.audioSystem.TurnOn()
	s.audioSystem.SetVolume(80)
	fmt.Println("聚会模式已启动！")
	fmt.Println()
}

// LeaveHomeMode 离家模式
func (s *SmartHomeFacade) LeaveHomeMode() {
	fmt.Println("========== 离家模式 ==========")
	s.lightSystem.TurnOff()
	s.airConditioner.TurnOff()
	s.audioSystem.TurnOff()
	fmt.Println("离家模式已启动！")
	fmt.Println()
}
