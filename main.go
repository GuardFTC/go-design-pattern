// @Author:冯铁城 [17615007230@163.com] 2025-09-17 10:32:44
package main

import (
	"degisn-pettern/create/builder"
	"degisn-pettern/create/factory"
	"degisn-pettern/create/prototype"
	"degisn-pettern/create/singleton"
	"degisn-pettern/structural/adapter"
	"degisn-pettern/structural/bridge"
	"degisn-pettern/structural/decorator"
	"degisn-pettern/structural/facade"
	"degisn-pettern/structural/proxy"
	"fmt"
)

func main() {

	//1.单例模式测试
	//singletonPetternTest()

	//2.工厂模式测试
	//factoryPatternTest()

	//3.建造者模式测试
	//builderTest()

	//4.原型模式测试
	//prototypeTest()

	//5.装饰器模式测试
	//decoratorTest()

	//6.适配器模式测试
	//adapterTest()

	//7.代理模式测试
	//proxyTest()

	//8.桥接模式测试
	//bridgeTest()

	//9.外观模式测试
	facadeTest()
}

// 单例模式测试
func singletonPetternTest() {

	//1.获取懒汉式实例
	lazy := singleton.GetInstanceLazy()
	fmt.Printf("lazy singleton name: %s\n", lazy.GetName())

	//2.获取饿汉式实例
	eager := singleton.GetInstanceEager()
	fmt.Printf("eager singleton name: %s\n", eager.GetName())

	//3.再次获取懒汉式实例,比较内存地址是否一致
	lazy2 := singleton.GetInstanceLazy()
	fmt.Printf("lazy address: %p\nlazy2 address: %p\n", lazy, lazy2)
	fmt.Printf("lazy == lazy2: %t\n", lazy == lazy2)

	//4.再次获取饿汉式实例,比较内存地址是否一致
	eager2 := singleton.GetInstanceEager()
	fmt.Printf("eager address: %p\neager2 address: %p\n", eager, eager2)
	fmt.Printf("eager == eager2: %t\n", eager == eager2)
}

// 工厂模式测试
func factoryPatternTest() {

	//1.简单工厂模式
	simpleFactory := new(factory.ProductSimpleFactory)
	product1 := simpleFactory.GetProduct(1)
	fmt.Printf("product1 name: %s\n", product1.GetName())
	product2 := simpleFactory.GetProduct(2)
	fmt.Printf("product2 name: %s\n", product2.GetName())
	product3 := simpleFactory.GetProduct(3)
	fmt.Printf("product3 name: %s\n", product3.GetName())

	//2.工厂方法模式
	factory1 := new(factory.ProductFactory1)
	product1 = factory1.CreateProduct()
	fmt.Printf("product1 name: %s\n", product1.GetName())
	factory2 := new(factory.ProductFactory2)
	product2 = factory2.CreateProduct()
	fmt.Printf("product2 name: %s\n", product2.GetName())
	factory3 := new(factory.ProductFactory3)
	product3 = factory3.CreateProduct()
	fmt.Printf("product3 name: %s\n", product3.GetName())
}

// 建造者模式测试
func builderTest() {

	//1.通过链式调用创建结构体对象
	testStruct := builder.NewTestStructBuilder().
		SetId(1).
		SetName("test").
		SetAge(18).
		SetSex("male").
		Build()

	//2.输出信息
	fmt.Printf("testStruct: %+v\n", testStruct)
}

// 原型模式测试
func prototypeTest() {

	//1.创建结构体
	people := &prototype.People{
		Name: "test",
		Ids:  []int{1, 2, 3},
		Age:  18,
	}

	//2.克隆
	people2 := people.Clone().(*prototype.People)

	//3.输出结构体信息
	fmt.Printf("people: %+v\npeople2: %+v\n", people, people2)

	//4.内存地址比较
	fmt.Printf("people == people2: %t\n", people == people2)

	//5.切片属性内存地址比较
	fmt.Printf("people.Ids address %p\npeople2.Ids address %p\n", people.Ids, people2.Ids)
}

// 装饰器模式测试
func decoratorTest() {

	//1.创建基础房间
	room := decorator.NewBaseRoom()
	fmt.Printf("base room: %+v\n", room.Show())

	//2.创建带门房间
	doorRoom := decorator.NewDoorRoomDecorator(room)
	fmt.Printf("door room: %+v\n", doorRoom.Show())

	//3.创建带窗格房间
	windowRoom := decorator.NewWindowRoomDecorator(room)
	fmt.Printf("window room: %+v\n", windowRoom.Show())

	//4.创建带门窗格房间
	doorWindowRoom := decorator.NewWindowRoomDecorator(doorRoom)
	fmt.Printf("door window room: %+v\n", doorWindowRoom.Show())
}

// 适配器模式测试
func adapterTest() {

	//1.创建中国发电机
	chineseGenerator := adapter.ChineseGenerator{}
	fmt.Printf("chinese generator voltage: %d\n", chineseGenerator.GetVoltage220v())

	//2.创建发电机适配器
	voltageAdapter := new(adapter.VoltageAdapter)
	voltageAdapter.ChineseGenerator = chineseGenerator
	fmt.Printf("voltageAdapter voltage: %d\n", voltageAdapter.GetVoltage110v())
}

// 代理模式测试
func proxyTest() {

	//1.创建明星
	readStar := proxy.RealStar{Name: "ftc"}

	//2.创建经纪人代理
	agent := proxy.Agent{
		Star: &readStar,
	}

	//3.唱歌
	agent.Sing()

	//4.商业活动
	agent.Businesses()
}

// 桥接模式测试
func bridgeTest() {

	//1.创建颜色
	blue := new(bridge.Blue)
	green := new(bridge.Green)
	red := new(bridge.Red)

	//2.创建图形，桥接颜色
	circle := bridge.NewCircle(blue)
	square := bridge.NewSquare(green)
	rectangle := bridge.NewRectangle(red)

	//3.展示
	circle.Show()
	square.Show()
	rectangle.Show()
	rectangle.SetColor(blue)
	rectangle.Show()
	rectangle.SetColor(green)
	rectangle.Show()
}

// 外观模式测试
func facadeTest() {

	//1.创建智能家居控制器
	smartHome := facade.NewSmartHomeFacade()

	//2.测试起床模式
	smartHome.WakeUpMode()

	//3.测试聚会模式
	smartHome.PartyMode()

	//4.测试睡眠模式
	smartHome.SleepMode()

	//5.测试离家模式
	smartHome.LeaveHomeMode()
}
