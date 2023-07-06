//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 作用域，变量可以使用的范围
	 *
	 * 局部变量 函数内部定义的变量 变量在哪里定义，就在那个范围使用，超出这个范围，变量被销毁
	 * 全局变量 函数外部定义的变量 所有函数都可以使用，共享这个变量数据
	 */

	//n 只能在主函数使用
	n := 10
	fmt.Println(n)

	//a 在if语句定义，只能在if语句内使用
	if a := 1; a <= 10 {
		n := 20        //作用域不同，在if内定义n
		fmt.Println(a) //1
		fmt.Println(n) //20
	}

	//fmt.Println(a) //undefined: a
	fmt.Println(n)
}

func fun1() {
	//fmt.Println(n) //undefined: n
}
