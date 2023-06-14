//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 函数本质，复合类型，一种特殊的变量
	 * fun1   指向函数体内存地址
	 * fun1() 调用
	 */
	fmt.Printf("%T \n", fun1)
	fmt.Println(fun1)
	// fun1()调用 fun1数据类型

	fmt.Println("============================")

	//函数可以赋值给一个变量
	var c func(int, int)
	//nil 空
	fmt.Println(c)

	//将fun1赋值给c，函数体地址赋值给c
	c = fun1
	fmt.Println(c)

	fmt.Println("============================")

	//两种方式可以调用
	fun1(1, 2)
	c(100, 200)

	fmt.Println("============================")

	//将fun2值【函数体地址】赋值给res1
	res1 := fun2
	//调用fun2()，执行函数，将函数执行结果返回给res2
	res2 := fun2(1, 2)
	fmt.Println(res1)
	fmt.Println(res2)

	//res1()可以调用，相当于fun2()
	fmt.Println(res1(100, 200))
}

func fun1(a, b int) {
	fmt.Printf("a:%d, b:%d \n", a, b)
}

func fun2(a, b int) int {
	return a + b
}
