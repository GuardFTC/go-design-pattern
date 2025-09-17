// Package factory @Author:冯铁城 [17615007230@163.com] 2025-09-17 11:04:31
package factory

// IProduct 工厂模式产品接口
type IProduct interface {
	GetName() string
}

// Product1 工厂模式产品1
type Product1 struct {
	name string
}

// GetName 获取产品名称
func (p *Product1) GetName() string {
	return p.name
}

// Product2 工厂模式产品2
type Product2 struct {
	name string
}

// GetName 获取产品名称
func (p *Product2) GetName() string {
	return p.name
}

// Product3 工厂模式产品3
type Product3 struct {
	name string
}

// GetName 获取产品名称
func (p *Product3) GetName() string {
	return p.name
}

// ProductSimpleFactory 简单工厂模式工厂
type ProductSimpleFactory struct {
}

// getProduct 获取产品
func (p *ProductSimpleFactory) GetProduct(productType int) IProduct {
	switch productType {
	case 1:
		return &Product1{name: "product1"}

	case 2:
		return &Product2{name: "product2"}

	case 3:
		return &Product3{name: "product3"}
	default:
		return nil
	}
}
