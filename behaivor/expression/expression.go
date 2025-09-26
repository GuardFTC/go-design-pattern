// Package expression @Author:冯铁城 [17615007230@163.com] 2025-09-26 15:33:00
package expression

// Context 上下文
type Context struct {
	Variables map[string]int
}

// SetVariable 设置变量
func (c *Context) SetVariable(key string, value int) {
	c.Variables[key] = value
}

// GetVariable 获取变量
func (c *Context) GetVariable(key string) int {
	if value, exist := c.Variables[key]; exist {
		return value
	} else {
		panic("变量不存在")
	}
}

// NewContext 创建上下文
func NewContext() *Context {
	return &Context{
		Variables: make(map[string]int),
	}
}

// IExpression 解释器接口
type IExpression interface {
	Interpret(context *Context) int
}

// NumberExpression 数字解释器
type NumberExpression struct {
	Value int
}

// Interpret 解释
func (n *NumberExpression) Interpret(context *Context) int {
	return n.Value
}

// NewNumberExpression 创建数字解释器
func NewNumberExpression(value int) *NumberExpression {
	return &NumberExpression{
		Value: value,
	}
}

// AddExpression 加法解释器
type AddExpression struct {
	Left  IExpression
	Right IExpression
}

// Interpret 解释
func (a *AddExpression) Interpret(context *Context) int {
	return a.Left.Interpret(context) + a.Right.Interpret(context)
}

// NewAddExpression 创建加法解释器
func NewAddExpression(left IExpression, right IExpression) *AddExpression {
	return &AddExpression{
		Left:  left,
		Right: right,
	}
}

// SubExpression 减法解释器
type SubExpression struct {
	Left  IExpression
	Right IExpression
}

// Interpret 解释
func (s *SubExpression) Interpret(context *Context) int {
	return s.Left.Interpret(context) - s.Right.Interpret(context)
}

// NewSubExpression 创建减法解释器
func NewSubExpression(left IExpression, right IExpression) *SubExpression {
	return &SubExpression{
		Left:  left,
		Right: right,
	}
}

// MulExpression 乘法解释器
type MulExpression struct {
	Left  IExpression
	Right IExpression
}

// Interpret 解释
func (m *MulExpression) Interpret(context *Context) int {
	return m.Left.Interpret(context) * m.Right.Interpret(context)
}

// NewMulExpression 创建乘法解释器
func NewMulExpression(left IExpression, right IExpression) *MulExpression {
	return &MulExpression{
		Left:  left,
		Right: right,
	}
}

// DivExpression 除法解释器
type DivExpression struct {
	Left  IExpression
	Right IExpression
}

// Interpret 解释
func (d *DivExpression) Interpret(context *Context) int {

	//1.获取右侧表达式计算结果
	rightValue := d.Right.Interpret(context)

	//1.异常判定
	if rightValue == 0 {
		panic("除数不能为0")
	}

	//2.解释返回
	return d.Left.Interpret(context) / rightValue
}

// NewDivExpression 创建除法解释器
func NewDivExpression(left IExpression, right IExpression) *DivExpression {
	return &DivExpression{
		Left:  left,
		Right: right,
	}
}

// BracketExpression 括号解释器
type BracketExpression struct {
	Expression IExpression
}

// Interpret 解释
func (b *BracketExpression) Interpret(context *Context) int {
	return b.Expression.Interpret(context)
}

// NewBracketExpression 创建括号解释器
func NewBracketExpression(expression IExpression) *BracketExpression {
	return &BracketExpression{
		Expression: expression,
	}
}

// VariableExpression 变量解释器
type VariableExpression struct {
	Name string
}

// Interpret 解释
func (v *VariableExpression) Interpret(context *Context) int {
	return context.GetVariable(v.Name)
}

// NewVariableExpression 创建变量解释器
func NewVariableExpression(name string) *VariableExpression {
	return &VariableExpression{
		Name: name,
	}
}
