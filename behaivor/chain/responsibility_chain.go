// Package chain @Author:冯铁城 [17615007230@163.com] 2025-09-25 16:04:07
package chain

// IChainNodeV2 链节点接口
type IChainNodeV2 interface {
	SetNextNode(node IChainNodeV2)
	GetNextNode() IChainNodeV2
	CheckNumberV2(number int) bool
}

// ICheckNode 校验接口
type ICheckNode interface {
	Check(number int) bool
}

// CheckNode 校验节点
type CheckNode struct {
	nextNode IChainNodeV2
	ICheckNode
}

// SetNextNode 设置下一个节点
func (c *CheckNode) SetNextNode(node IChainNodeV2) {
	c.nextNode = node
}

// GetNextNode 获取下一个节点
func (c *CheckNode) GetNextNode() IChainNodeV2 {
	return c.nextNode
}

// CheckNumberV2 校验
func (c *CheckNode) CheckNumberV2(number int) bool {

	//1.校验
	if ok := c.Check(number); !ok {
		return false
	}

	//2.获取下一个节点
	nextNode := c.GetNextNode()
	if nextNode == nil {
		return true
	}

	//3.不为空交给下一个节点校验
	return nextNode.CheckNumberV2(number)
}

// NewCheckNode 创建校验节点
func NewCheckNode(checkNode ICheckNode) *CheckNode {
	return &CheckNode{
		ICheckNode: checkNode,
	}
}

// GreaterThan0NodeV2 校验是否>0节点
type GreaterThan0NodeV2 struct {
}

// Check 校验是否>0
func (g *GreaterThan0NodeV2) Check(number int) bool {
	return number > 0
}

// NewGreaterThan0NodeV2 创建>0节点
func NewGreaterThan0NodeV2() *GreaterThan0NodeV2 {
	return &GreaterThan0NodeV2{}
}

// DivisibleBy10NodeV2 校验是否能被10整除节点
type DivisibleBy10NodeV2 struct {
}

// Check 校验是否能被10整除
func (d *DivisibleBy10NodeV2) Check(number int) bool {
	return number%10 == 0
}

// NewDivisibleBy10NodeV2 创建能被10整除节点
func NewDivisibleBy10NodeV2() *DivisibleBy10NodeV2 {
	return &DivisibleBy10NodeV2{}
}

// DivisibleBy4NodeV2 校验是否能被4整除节点
type DivisibleBy4NodeV2 struct {
}

// Check 校验是否能被4整除
func (d *DivisibleBy4NodeV2) Check(number int) bool {
	return number%4 == 0
}

// NewDivisibleBy4NodeV2 创建能被4整除节点
func NewDivisibleBy4NodeV2() *DivisibleBy4NodeV2 {
	return &DivisibleBy4NodeV2{}
}
