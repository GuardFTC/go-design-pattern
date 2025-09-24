// Package strategy @Author:冯铁城 [17615007230@163.com] 2025-09-24 11:07:04
package strategy

// IProductPrice 产品价格接口
type IProductPrice interface {
	GetPrice(price float64) float64
}

// OriginalPrice 原价结构体
type OriginalPrice struct{}

// GetPrice 获取价格
func (o *OriginalPrice) GetPrice(price float64) float64 {
	return price
}

// NewOriginalPrice 创建原价结构体
func NewOriginalPrice() *OriginalPrice {
	return &OriginalPrice{}
}

// DiscountPrice 折扣结构体
type DiscountPrice struct {
	Discount float64
}

// GetPrice 获取价格
func (d *DiscountPrice) GetPrice(price float64) float64 {
	return price * d.Discount
}

// NewDiscountPrice 创建折扣结构体
func NewDiscountPrice(discount float64) *DiscountPrice {
	return &DiscountPrice{
		Discount: discount,
	}
}

// FullReductionStrategy 满减结构体
type FullReductionStrategy struct {
	FullPrice float64
	Reduction float64
}

// GetPrice 获取价格
func (f *FullReductionStrategy) GetPrice(price float64) float64 {
	if price >= f.FullPrice {
		return price - f.Reduction
	}
	return price
}

// NewFullReductionStrategy 创建满减结构体
func NewFullReductionStrategy(fullPrice float64, reduction float64) *FullReductionStrategy {
	return &FullReductionStrategy{
		FullPrice: fullPrice,
		Reduction: reduction,
	}
}

// ProductPriceContext 价格上下文结构体
type ProductPriceContext struct {
	price float64
	IProductPrice
}

// GetProductPrice 获取价格
func (p *ProductPriceContext) GetProductPrice() float64 {
	return p.GetPrice(p.price)
}

// SetIProductPrice 设置价格接口
func (p *ProductPriceContext) SetIProductPrice(productPrice IProductPrice) {
	p.IProductPrice = productPrice
}

// NewProductPriceContext 创建价格上下文结构体
func NewProductPriceContext(price float64, productPrice IProductPrice) *ProductPriceContext {
	return &ProductPriceContext{
		price:         price,
		IProductPrice: productPrice,
	}
}
