// Package status @Author:冯铁城 [17615007230@163.com] 2025-09-24 11:40:31
package status

import "fmt"

// OrderStatusContext 订单状态上下文结构体
type OrderStatusContext struct {
	IOrderStatus
}

// SetStatus 设置状态
func (o *OrderStatusContext) SetStatus(status IOrderStatus) {
	o.IOrderStatus = status
}

// ShowStatus 显示状态
func (o *OrderStatusContext) ShowStatus() {
	fmt.Printf("the order status is %s\n", o.GetCurrentStatusAndProcessNext(o))
}

// NewOrderStatusContext 创建订单状态上下文结构体
func NewOrderStatusContext(orderStatus IOrderStatus) *OrderStatusContext {
	return &OrderStatusContext{
		IOrderStatus: orderStatus,
	}
}

// IOrderStatus 订单状态接口 已创建->待支付->待发货->待收货->待评价->完成
type IOrderStatus interface {
	GetCurrentStatusAndProcessNext(o *OrderStatusContext) string
}

// OrderCreated 已创建状态结构体
type OrderCreated struct{}

// GetCurrentStatusAndProcessNext 获取状态
func (o *OrderCreated) GetCurrentStatusAndProcessNext(ctx *OrderStatusContext) string {
	ctx.SetStatus(NewOrderPaid())
	return "已创建"
}

// NewOrderCreated 创建已创建状态结构体
func NewOrderCreated() *OrderCreated {
	return &OrderCreated{}
}

// OrderPaid 待支付状态结构体
type OrderPaid struct{}

// GetCurrentStatusAndProcessNext 获取状态
func (o *OrderPaid) GetCurrentStatusAndProcessNext(ctx *OrderStatusContext) string {
	ctx.SetStatus(NewOrderShipped())
	return "待支付"
}

// NewOrderPaid 创建待支付状态结构体
func NewOrderPaid() *OrderPaid {
	return &OrderPaid{}
}

// OrderShipped 待发货状态结构体
type OrderShipped struct{}

// GetCurrentStatusAndProcessNext 获取状态
func (o *OrderShipped) GetCurrentStatusAndProcessNext(ctx *OrderStatusContext) string {
	ctx.SetStatus(NewOrderReceived())
	return "待发货"
}

// NewOrderShipped 创建待发货状态结构体
func NewOrderShipped() *OrderShipped {
	return &OrderShipped{}
}

// OrderReceived 待收货状态结构体
type OrderReceived struct{}

// GetCurrentStatusAndProcessNext 获取状态
func (o *OrderReceived) GetCurrentStatusAndProcessNext(ctx *OrderStatusContext) string {
	ctx.SetStatus(NewOrderEvaluated())
	return "待收货"
}

// NewOrderReceived 创建待收货状态结构体
func NewOrderReceived() *OrderReceived {
	return &OrderReceived{}
}

// OrderEvaluated 待评价状态结构体
type OrderEvaluated struct{}

// GetCurrentStatusAndProcessNext 获取状态
func (o *OrderEvaluated) GetCurrentStatusAndProcessNext(ctx *OrderStatusContext) string {
	ctx.SetStatus(NewOrderCompleted())
	return "待评价"
}

// NewOrderEvaluated 创建待评价状态结构体
func NewOrderEvaluated() *OrderEvaluated {
	return &OrderEvaluated{}
}

// OrderCompleted 完成状态结构体
type OrderCompleted struct{}

// GetCurrentStatusAndProcessNext 获取状态
func (o *OrderCompleted) GetCurrentStatusAndProcessNext(ctx *OrderStatusContext) string {
	return "完成"
}

// NewOrderCompleted 创建完成状态结构体
func NewOrderCompleted() *OrderCompleted {
	return &OrderCompleted{}
}
