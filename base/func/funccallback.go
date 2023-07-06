//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 高阶函数，可以将一个函数做为另一个函数的参数
	 * fun1() fun2()
	 * 将fun1()函数做为fun2()的参数
	 * fun2() 高阶函数
	 * fun1() 回调函数 参数
	 */
	fmt.Printf("%T \n", add)  //func(int, int) int
	fmt.Printf("%T \n", oper) //func(int, int, func(int, int) int) int

	fmt.Println("=========================================")

	res1 := add(1, 2)
	fmt.Println(res1)

	fmt.Println("=========================================")

	res2 := oper(10, 20, add)
	fmt.Println(res2)

	fmt.Println("=========================================")

	res3 := oper(100, 10, minus)
	fmt.Println(res3)

	fmt.Println("=========================================")

	//匿名函数
	fun1 := func(a, b int) int {
		return a * b
	}
	res4 := oper(3, 7, fun1)
	fmt.Println(res4)

	fmt.Println("=========================================")

	res5 := oper(100, 10, func(a, b int) int {
		return a / b
	})
	fmt.Println(res5)
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

//fun
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}
