// @Author:冯铁城 [17615007230@163.com] 2025-09-17 10:32:44
package main

import (
	"degisn-pettern/singleton"
	"fmt"
)

func main() {

	//1.单例模式测试

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
