// Package structural @Author:冯铁城 [17615007230@163.com] 2025-09-17 15:43:26
package decorator

// IRoom 抽象房间接口
type IRoom interface {
	Show() string
}

// RoomDecorator 房间装饰器
type RoomDecorator struct {
	IRoom
}

// Show 获取房间名称
func (a *RoomDecorator) Show() string {
	return a.IRoom.Show()
}

// BaseRoom 基础房间类
type BaseRoom struct{}

// NewBaseRoom 创建基础房间
func NewBaseRoom() *BaseRoom {
	return &BaseRoom{}
}

// Show 获取房间名称
func (b *BaseRoom) Show() string {
	return "with four wall"
}

// DoorRoomDecorator 带门房间
type DoorRoomDecorator struct {
	RoomDecorator
}

// NewDoorRoomDecorator 创建带门房间
func NewDoorRoomDecorator(room IRoom) *DoorRoomDecorator {
	return &DoorRoomDecorator{
		RoomDecorator: RoomDecorator{
			IRoom: room,
		},
	}
}

// Show 获取房间名称
func (d *DoorRoomDecorator) Show() string {
	return d.RoomDecorator.IRoom.Show() + " with one door"
}

// WindowRoomDecorator 带窗房间
type WindowRoomDecorator struct {
	RoomDecorator
}

// NewWindowRoomDecorator 创建带窗房间
func NewWindowRoomDecorator(room IRoom) *WindowRoomDecorator {
	return &WindowRoomDecorator{
		RoomDecorator: RoomDecorator{
			IRoom: room,
		},
	}
}

// Show 获取房间名称
func (w *WindowRoomDecorator) Show() string {
	return w.RoomDecorator.IRoom.Show() + " with one window"
}
