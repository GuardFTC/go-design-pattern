// @Author:冯铁城 [17615007230@163.com] 2025-09-17 10:32:44
package main

import (
	"degisn-pettern/behaivor/chain"
	"degisn-pettern/behaivor/command"
	"degisn-pettern/behaivor/mediator"
	"degisn-pettern/behaivor/observer"
	"degisn-pettern/behaivor/status"
	"degisn-pettern/behaivor/strategy"
	"degisn-pettern/behaivor/visitor"
	"degisn-pettern/create/builder"
	"degisn-pettern/create/factory"
	"degisn-pettern/create/prototype"
	"degisn-pettern/create/singleton"
	"degisn-pettern/structural/adapter"
	"degisn-pettern/structural/bridge"
	"degisn-pettern/structural/composite"
	"degisn-pettern/structural/decorator"
	"degisn-pettern/structural/facade"
	"degisn-pettern/structural/flyweight"
	"degisn-pettern/structural/proxy"
	"fmt"
)

func main() {

	//------------------------------------------创建型-------------------------------------------------//
	//1.单例模式测试
	//singletonPetternTest()

	//2.工厂模式测试
	//factoryPatternTest()

	//3.建造者模式测试
	//builderTest()

	//4.原型模式测试
	//prototypeTest()

	//------------------------------------------结构型-------------------------------------------------//
	//5.装饰器模式测试
	//decoratorTest()

	//6.适配器模式测试
	//adapterTest()

	//7.代理模式测试
	//proxyTest()

	//8.桥接模式测试
	//bridgeTest()

	//9.外观模式测试
	//facadeTest()

	//10.组合模式测试
	//compositeTest()

	//11.享元模式测试
	//flyweightTest()

	//------------------------------------------行为型-------------------------------------------------//
	//12.策略模式测试
	//strategyTest()

	//13.状态模式测试
	//statusTest()

	//14.命令模式测试
	//commandTest()

	//15.观察者模式测试
	//observerTest()

	//16.访问者模式测试
	//visitorTest()

	//17.中介者模式测试
	//mediatorTest()

	//18.责任链模式测试
	//chainTestV1()

	//1.创建校验节点
	lastNode := chain.NewCheckNode(chain.NewDivisibleBy4NodeV2())
	secondNode := chain.NewCheckNode(chain.NewDivisibleBy10NodeV2())
	firstNode := chain.NewCheckNode(chain.NewGreaterThan0NodeV2())
	firstNode.SetNextNode(secondNode)
	secondNode.SetNextNode(lastNode)

	//2.校验数据
	fmt.Printf("check number: %v result is %v\n", -1, firstNode.CheckNumberV2(-1))
	fmt.Printf("check number: %v result is %v\n", 3, firstNode.CheckNumberV2(3))
	fmt.Printf("check number: %v result is %v\n", 30, firstNode.CheckNumberV2(30))
	fmt.Printf("check number: %v result is %v\n", 40, firstNode.CheckNumberV2(40))
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

// 组合模式测试
func compositeTest() {

	//1.创建1级组件
	level11 := &composite.Level1{Name: "1级组件-1"}
	level12 := &composite.Level1{Name: "1级组件-2"}

	//2.创建2级组件
	components := make([]composite.Component, 0)
	components = append(components, level11)
	components = append(components, level12)

	level21 := &composite.Level2{Name: "2级组件-1", Children: components}

	//3.创建3级组件
	components = make([]composite.Component, 0)
	components = append(components, level21)

	level31 := &composite.Level3{Name: "3级组件-1", Children: components}

	//4.3级组件展示细节
	level31.ShowDetail()
}

// 组合模式测试
func flyweightTest() {

	//1.获取棋子工厂(享元工厂)
	chessPieceFactory := flyweight.NewChessPieceFactory()

	//2.创建一批黑色棋子
	blackChessPiece := chessPieceFactory.GetChessPieceByColor("black")
	blackChessPiece.ShowCoordinates(flyweight.ChessPieceContext{X: 1, Y: 1})
	blackChessPiece.ShowCoordinates(flyweight.ChessPieceContext{X: 2, Y: 2})
	blackChessPiece.ShowCoordinates(flyweight.ChessPieceContext{X: 3, Y: 3})

	//3.创建一批白色棋子
	whiteChessPiece := chessPieceFactory.GetChessPieceByColor("white")
	whiteChessPiece.ShowCoordinates(flyweight.ChessPieceContext{X: 4, Y: 4})
	whiteChessPiece.ShowCoordinates(flyweight.ChessPieceContext{X: 5, Y: 5})
	whiteChessPiece.ShowCoordinates(flyweight.ChessPieceContext{X: 6, Y: 6})
}

// 策略模式测试
func strategyTest() {

	//1.创建3种获取价格接口
	originalPrice := strategy.NewOriginalPrice()
	discountPrice := strategy.NewDiscountPrice(0.8)
	fullReductionPrice := strategy.NewFullReductionStrategy(100, 30)

	//2.创建价格上下文结构体，初始策略为原价
	productPriceContext := strategy.NewProductPriceContext(100, originalPrice)
	fmt.Printf("当前商品价格: %+v\n", productPriceContext.GetProductPrice())

	//3.调整为打折模式
	productPriceContext.SetIProductPrice(discountPrice)
	fmt.Printf("当前商品价格: %+v\n", productPriceContext.GetProductPrice())

	//4.调整为满减模式
	productPriceContext.SetIProductPrice(fullReductionPrice)
	fmt.Printf("当前商品价格: %+v\n", productPriceContext.GetProductPrice())
}

// 状态模式测试
func statusTest() {

	//1.创建初始状态
	orderCreated := status.NewOrderCreated()

	//2.创建状态上下文结构体
	orderStatusContext := status.NewOrderStatusContext(orderCreated)

	//3.获取订单状态
	fmt.Printf("order status: %s\n", orderStatusContext.GetCurrentStatusAndProcessNext(orderStatusContext))
	fmt.Printf("order status: %s\n", orderStatusContext.GetCurrentStatusAndProcessNext(orderStatusContext))
	fmt.Printf("order status: %s\n", orderStatusContext.GetCurrentStatusAndProcessNext(orderStatusContext))
	fmt.Printf("order status: %s\n", orderStatusContext.GetCurrentStatusAndProcessNext(orderStatusContext))
	fmt.Printf("order status: %s\n", orderStatusContext.GetCurrentStatusAndProcessNext(orderStatusContext))
	fmt.Printf("order status: %s\n", orderStatusContext.GetCurrentStatusAndProcessNext(orderStatusContext))
}

// 命令模式测试
func commandTest() {

	//1.创建士兵
	soldier1 := command.NewSoldier("马冬梅")
	soldier2 := command.NewSoldier("夏洛")

	//2.创建命令
	walkCommand1 := command.NewWalkCommand(soldier1)
	runCommand1 := command.NewRunCommand(soldier1)
	runCommand12 := command.NewRunCommand(soldier1)
	attackCommand1 := command.NewAttentionCommand(soldier1)

	walkCommand2 := command.NewWalkCommand(soldier2)
	runCommand2 := command.NewRunCommand(soldier2)
	runCommand22 := command.NewRunCommand(soldier2)
	attackCommand2 := command.NewAttentionCommand(soldier2)

	//3.存入切片
	commands := []command.Command{walkCommand1, runCommand1, runCommand12, attackCommand1, walkCommand2, runCommand2, runCommand22, attackCommand2}

	//4.创建长官
	commander := command.NewCommander(commands)

	//5.执行命令
	commander.ExecuteCommands()
}

// 观察者模式测试
func observerTest() {

	//1.创建发布者
	publisher := observer.NewPublisher()

	//2.创建订阅者
	subscriber1 := observer.NewSubscriber("马冬梅")
	subscriber2 := observer.NewSubscriber("夏洛")
	subscriber3 := observer.NewSubscriber("大傻春")

	//3.订阅
	publisher.AddSubscriber(subscriber1.Name, subscriber1)
	publisher.AddSubscriber(subscriber2.Name, subscriber2)
	publisher.AddSubscriber(subscriber3.Name, subscriber3)

	//4.发布消息
	publisher.Notify("hello world first time")

	//5.大傻春取消订阅
	publisher.RemoveSubscriber(subscriber3.Name)

	//6.再次发布消息
	publisher.Notify("hello world second time")
}

// 访问者模式测试
func visitorTest() {

	//1.创建5个位置（元素）
	pg := visitor.PG{}
	sg := visitor.SG{}
	sf := visitor.SF{}
	pf := visitor.PF{}
	c := visitor.C{}

	//2.创建访问者
	bodyTrain := visitor.BodyTrain{}
	shootTrain := visitor.ShootTrain{}
	skillTrain := visitor.SkillTrain{}

	//3.访问元素，进行功能扩展
	pg.Accept(&bodyTrain)
	pg.Accept(&shootTrain)
	pg.Accept(&skillTrain)
	sg.Accept(&bodyTrain)
	sg.Accept(&shootTrain)
	sg.Accept(&skillTrain)
	sf.Accept(&bodyTrain)
	sf.Accept(&shootTrain)
	sf.Accept(&skillTrain)
	pf.Accept(&bodyTrain)
	pf.Accept(&shootTrain)
	pf.Accept(&skillTrain)
	c.Accept(&bodyTrain)
	c.Accept(&shootTrain)
	c.Accept(&skillTrain)
}

// 中介者模式测试
func mediatorTest() {

	//1.创建媒婆(中介者)
	matchmaker := mediator.NewMatchmaker()

	//2.创建男孩女孩(同事类)
	boy := mediator.NewBoy(matchmaker)
	girl := mediator.NewGirl(matchmaker)

	//3.中介者设置男孩女孩
	matchmaker.SetBoy(boy)
	matchmaker.SetGirl(girl)

	//4.男孩找相亲
	boy.FindGirl()

	//5.女孩找相亲
	girl.FindBoy()
}

// 责任链模式测试-(过滤器链)
func chainTestV1() {

	//1.创建节点
	greaterThan0Node := chain.NewGreaterThan0Node()
	divisibleBy10Node := chain.NewDivisibleBy10Node()
	divisibleBy4Node := chain.NewDivisibleBy4Node()

	//2.创建责任链
	verifyChain := chain.NewVerifyChain(greaterThan0Node, divisibleBy10Node, divisibleBy4Node)

	//3.校验数字
	fmt.Printf("check number: %v result is %v\n", -1, verifyChain.CheckNumber(-1))
	fmt.Printf("check number: %v result is %v\n", 3, verifyChain.CheckNumber(3))
	fmt.Printf("check number: %v result is %v\n", 30, verifyChain.CheckNumber(30))
	fmt.Printf("check number: %v result is %v\n", 40, verifyChain.CheckNumber(40))
}
