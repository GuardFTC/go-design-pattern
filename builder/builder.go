// Package builder @Author:冯铁城 [17615007230@163.com] 2025-09-17 14:11:02
package builder

// TestStruct 测试结构体
type TestStruct struct {
	id   int
	name string
	age  int
	sex  string
}

// NewTestStructBuilder 创建测试结构体构造器
func NewTestStructBuilder() *TestStructBuilder {
	return &TestStructBuilder{
		TestStruct: TestStruct{},
	}
}

// TestStructBuilder 测试结构体构造器
type TestStructBuilder struct {
	TestStruct
}

// Build 构造方法
func (t *TestStructBuilder) Build() *TestStruct {
	return &t.TestStruct
}

// Reset 重置
func (t *TestStructBuilder) Reset() *TestStructBuilder {
	t.TestStruct = TestStruct{}
	return t
}

// SetId 设置id
func (t *TestStructBuilder) SetId(id int) *TestStructBuilder {
	t.TestStruct.id = id
	return t
}

// SetName 设置name
func (t *TestStructBuilder) SetName(name string) *TestStructBuilder {
	t.TestStruct.name = name
	return t
}

// SetAge 设置age
func (t *TestStructBuilder) SetAge(age int) *TestStructBuilder {
	t.TestStruct.age = age
	return t
}

// SetSex 设置性别
func (t *TestStructBuilder) SetSex(sex string) *TestStructBuilder {
	t.TestStruct.sex = sex
	return t
}
