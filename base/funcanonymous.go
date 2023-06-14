//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * Go支持函数式编程
	 * 1.将匿名函数做为另一个函数的参数，回调函数
	 * 2.将匿名函数做为另一个函数的返回值，闭包函数
	 */
	fun1()

	fun2 := fun1
	fun2()

	//匿名函数
	func() {
		fmt.Println("我是匿名函数")
	}() //匿名函数调用，通常只能使用一次

	fun3 := func() {
		fmt.Println("我也是匿名函数")
	}
	fun3()

	//匿名函数参数
	func(a, b int) { //形参
		fmt.Printf("a:%d, b:%d \n", a, b)
	}(1, 2) //实参

	res1 := func(a, b int) int { //形参
		return a + b
	}(10, 20) //实参 匿名函数调用，将执行结果赋值给res1
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	} //将匿名函数的值，赋值给res2
	fmt.Println(res2)

	fmt.Println(res2(100, 200))
}

func fun1() {
	fmt.Println("我是fun1()函数")
}
