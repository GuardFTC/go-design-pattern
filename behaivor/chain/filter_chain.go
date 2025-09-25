// Package chain @Author:冯铁城 [17615007230@163.com] 2025-09-25 15:28:09
package chain

// IChainNode 链节点接口
type IChainNode interface {
	CheckNumber(number int) bool
}

// GreaterThan0Node 校验是否>0节点
type GreaterThan0Node struct{}

// CheckNumber 校验是否>0
func (g *GreaterThan0Node) CheckNumber(number int) bool {
	return number > 0
}

// NewGreaterThan0Node 创建>0节点
func NewGreaterThan0Node() *GreaterThan0Node {
	return &GreaterThan0Node{}
}

// DivisibleBy10Node 校验是否能被10整除节点
type DivisibleBy10Node struct{}

// CheckNumber 校验是否能被10整除
func (d *DivisibleBy10Node) CheckNumber(number int) bool {
	return number%10 == 0
}

// NewDivisibleBy10Node 创建能被10整除节点
func NewDivisibleBy10Node() *DivisibleBy10Node {
	return &DivisibleBy10Node{}
}

// DivisibleBy4Node 校验是否能被4整除节点
type DivisibleBy4Node struct{}

// CheckNumber 校验是否能被4整除
func (d *DivisibleBy4Node) CheckNumber(number int) bool {
	return number%4 == 0
}

// NewDivisibleBy4Node 创建能被4整除节点
func NewDivisibleBy4Node() *DivisibleBy4Node {
	return &DivisibleBy4Node{}
}

// VerifyChain 验证链
type VerifyChain struct {
	nodes []IChainNode
}

// CheckNumber 验证数字
func (v *VerifyChain) CheckNumber(number int) bool {
	for _, node := range v.nodes {
		if !node.CheckNumber(number) {
			return false
		}
	}
	return true
}

// NewVerifyChain 创建验证链
func NewVerifyChain(nodes ...IChainNode) *VerifyChain {
	return &VerifyChain{
		nodes: nodes,
	}
}
