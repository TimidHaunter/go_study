//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/**
	 * 高阶函数，可以将一个函数做为另一个函数的返回值
	 *
	 * 将fun()函数做为increment()的返回值
	 * increment() 外层函数
	 * func() 内层函数 返回值
	 *
	 * 一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外部函数中的参数、外部函数中直接定义变量），并且该外层函数的返回值就是这个内层函数。
	 *
	 * 这个内层函数和外层函数中的局部变量，统称未闭包函数
	 * 局部变量生命周期会发生变化，正常的局部变量随着函数调用而创建，随着函数的结束而销毁
	 * 但是在闭包函数中，局部变量不会随着外部函数的结束而销毁，因为内层函数还要继续使用
	 */

	res1 := increment()
	fmt.Printf("%T \n", res1)
	fmt.Println(res1)
	fmt.Println(res1()) //调用increment()中的匿名函数
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	fmt.Println("===================================")

	res2 := increment()
	fmt.Println(res2())
}

func increment() func() int { //外层函数

	//定义局部变量
	i := 0
	//定义匿名函数，给变量自增，返回
	fun := func() int { //内层函数
		i++
		return i
	}

	return fun //返回变量
}
