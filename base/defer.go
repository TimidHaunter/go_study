//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 使用defer关键字来延迟一个函数和方法的执行
	 * 外围函数
	 */

	fun1("Yin")
	fmt.Println("=====")
	defer fun1("tian") //defer 在外围函数执行完毕后，再执行defer函数
	fmt.Println("=====")
}

func fun1(s string) {
	fmt.Println(s)
}
