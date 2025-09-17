// @Author:冯铁城 [17615007230@163.com] 2025-09-17 10:32:44
package main

import (
	"degisn-pettern/create/builder"
	"degisn-pettern/create/factory"
	"degisn-pettern/create/prototype"
	"degisn-pettern/create/singleton"
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
