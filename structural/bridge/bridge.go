// Package bridge @Author:冯铁城 [17615007230@163.com] 2025-09-23 14:42:40
package bridge

import "fmt"

// Color 颜色接口
type Color interface {
	ShowColor()
}

// Red 红色
type Red struct{}

func (r *Red) ShowColor() {
	fmt.Printf("the color is red\n")
}

// Blue 蓝色
type Blue struct{}

func (b *Blue) ShowColor() {
	fmt.Printf("the color is blue\n")
}

// Green 绿色
type Green struct{}

func (g *Green) ShowColor() {
	fmt.Printf("the color is green\n")
}

// Shape 形状结构体
type Shape struct {
	Color
	Name string
}

// SetColor 设置颜色
func (s *Shape) SetColor(color Color) *Shape {
	s.Color = color
	return s
}

// Show 显示
func (s *Shape) Show() {
	fmt.Printf("the shape is %s\n", s.Name)
	s.Color.ShowColor()
	fmt.Println("----------------------------")
}

// Circle 圆形
type Circle struct {
	Shape
}

// NewCircle 创建圆形
func NewCircle(color Color) *Circle {
	return &Circle{
		Shape: Shape{
			Name:  "circle",
			Color: color,
		},
	}
}

// Square 正方形
type Square struct {
	Shape
}

// NewSquare 创建正方形
func NewSquare(color Color) *Square {
	return &Square{
		Shape: Shape{
			Color: color,
			Name:  "square",
		},
	}
}

// Rectangle 矩形
type Rectangle struct {
	Shape
}

// NewRectangle 创建矩形
func NewRectangle(color Color) *Rectangle {
	return &Rectangle{
		Shape: Shape{
			Color: color,
			Name:  "rectangle",
		},
	}
}
