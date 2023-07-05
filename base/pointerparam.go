//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 指针作为参数
	 *
	 * 参数的传递
	 *   值传递   副本
	 *   引用传递 传递的地址，指向同一个内存
	 */
	a := 10
	fmt.Println("fun1()调用前num值：", a)
	fun1(a)
	fmt.Println("fun1()调用后num值：", a)
}

// 参数是int
func fun1(num int) {
	fmt.Println("fun1()中num值：", num)
	num = 100
	fmt.Println("fun1()中num值修改后：", num)
}

// 参数是int的指针
func fun2(p1 *int) {
	fmt.Println("fun2()中，p1：", *p1)
}
