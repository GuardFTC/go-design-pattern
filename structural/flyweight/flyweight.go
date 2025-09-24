// Package flyweight @Author:冯铁城 [17615007230@163.com] 2025-09-24 10:22:27
package flyweight

import "fmt"

// ChessPieceContext 棋子上下文
type ChessPieceContext struct {
	X int
	Y int
}

// IChessPiece 棋子接口
type IChessPiece interface {
	ShowCoordinates(ctx ChessPieceContext)
}

// ChessPiece 棋子结构体
type ChessPiece struct {
	color string
}

// ShowCoordinates 显示坐标
func (w *ChessPiece) ShowCoordinates(ctx ChessPieceContext) {
	fmt.Printf("the %s chess piece coordinates is (%d,%d)\n", w.color, ctx.X, ctx.Y)
}

// NewChessPiece 创建棋子
func NewChessPiece(color string) *ChessPiece {
	return &ChessPiece{
		color: color,
	}
}

// ChessPieceFactory 棋子工厂结构体
type ChessPieceFactory struct {
	chessPieces map[string]IChessPiece
}

// GetChessPieceByColor 获取棋子
func (c *ChessPieceFactory) GetChessPieceByColor(color string) IChessPiece {

	//1.如果不存在棋子，则创建
	if _, isPresent := c.chessPieces[color]; !isPresent {
		c.chessPieces[color] = NewChessPiece(color)
	}

	//2.返回棋子
	return c.chessPieces[color]
}

// NewChessPieceFactory 创建棋子工厂
func NewChessPieceFactory() *ChessPieceFactory {
	return &ChessPieceFactory{
		chessPieces: make(map[string]IChessPiece),
	}
}
