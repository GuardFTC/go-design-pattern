// Package command @Author:冯铁城 [17615007230@163.com] 2025-09-24 14:44:58
package command

import "fmt"

// Commander 长官
type Commander struct {
	commands []Command
}

// ExecuteCommands 执行命令
func (c *Commander) ExecuteCommands() {
	for _, command := range c.commands {
		command.Execute()
	}
}

// NewCommander 创建长官
func NewCommander(commands []Command) *Commander {
	return &Commander{
		commands: commands,
	}
}

// Soldier 士兵
type Soldier struct {
	name string
}

// Walk 走
func (s *Soldier) Walk() {
	fmt.Printf("Soldier %v is walking...\n", s.name)
}

// Run 跑
func (s *Soldier) Run() {
	fmt.Printf("Soldier %v is running...\n", s.name)
}

// Attention 立正
func (s *Soldier) Attention() {
	fmt.Printf("Soldier %v is attending...\n", s.name)
}

// NewSoldier 创建士兵
func NewSoldier(name string) *Soldier {
	return &Soldier{name: name}
}

// Command 命令接口
type Command interface {
	Execute()
}

// WalkCommand 走命令
type WalkCommand struct {
	soldier *Soldier
}

// Execute 执行
func (w *WalkCommand) Execute() {
	w.soldier.Walk()
}

// NewWalkCommand 创建走命令
func NewWalkCommand(soldier *Soldier) *WalkCommand {
	return &WalkCommand{soldier: soldier}
}

// RunCommand 跑命令
type RunCommand struct {
	soldier *Soldier
}

// Execute 执行
func (r *RunCommand) Execute() {
	r.soldier.Run()
}

// NewRunCommand 创建跑命令
func NewRunCommand(soldier *Soldier) *RunCommand {
	return &RunCommand{soldier: soldier}
}

// AttentionCommand 立正命令
type AttentionCommand struct {
	soldier *Soldier
}

// Execute 执行
func (a *AttentionCommand) Execute() {
	a.soldier.Attention()
}

// NewAttentionCommand 创建立正命令
func NewAttentionCommand(soldier *Soldier) *AttentionCommand {
	return &AttentionCommand{soldier: soldier}
}
