// Package prototype @Author:冯铁城 [17615007230@163.com] 2025-09-17 14:53:35
package prototype

// IPrototype 原型接口
type IPrototype interface {
	Clone() IPrototype
}

// People 人
type People struct {
	Name string
	Age  int
	Ids  []int
}

// Clone 克隆
func (p *People) Clone() IPrototype {

	//1.定义ID切片
	ids := make([]int, len(p.Ids))
	copy(ids, p.Ids)

	//2.创建结构体返回
	return &People{
		Name: p.Name,
		Age:  p.Age,
		Ids:  ids,
	}
}
