// Package adapter @Author:冯铁城 [17615007230@163.com] 2025-09-18 10:26:57
package adapter

// ChineseGenerator 中国发电机
type ChineseGenerator struct {
}

// GetVoltage220v 获取220伏特电压
func (c *ChineseGenerator) GetVoltage220v() int {
	return 220
}

// JapaneseGenerator 日式发电机
type JapaneseGenerator interface {

	// GetVoltage110v 获取110伏特电压
	GetVoltage110v() int
}

// VoltageAdapter 电压适配器
type VoltageAdapter struct {
	ChineseGenerator
}

// GetVoltage110v 获取110伏特电压
func (v *VoltageAdapter) GetVoltage110v() int {
	return v.GetVoltage220v() / 2
}
