// Package factory @Author:冯铁城 [17615007230@163.com] 2025-09-17 11:24:13
package factory

// IProductFactory 工厂接口
type IProductFactory interface {
	CreateProduct() IProduct
}

// ProductFactory1 工厂1
type ProductFactory1 struct {
}

// CreateProduct 创建产品1
func (f *ProductFactory1) CreateProduct() IProduct {
	return &Product1{
		name: "product1",
	}
}

// ProductFactory2 工厂2
type ProductFactory2 struct {
}

// CreateProduct 创建产品2
func (f *ProductFactory2) CreateProduct() IProduct {
	return &Product2{
		name: "product2",
	}
}

// ProductFactory3 工厂3
type ProductFactory3 struct {
}

// CreateProduct 创建产品3
func (f *ProductFactory3) CreateProduct() IProduct {
	return &Product3{
		name: "product3",
	}
}
