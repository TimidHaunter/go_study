//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	 * 定义数据类型[struct/interface]
	 * type 类型名 Type
	 */
	var i1 yint
	var i2 int = 1000

	i1 = 10
	fmt.Printf("i1 %d %T \n", i1, i1)
	fmt.Printf("i2 %d %T \n", i2, i2)

	//i1 = i2 cannot use i2 (variable of type int) as yint value in assignment 不能将i2(int)赋值给i1 yint 其实是两种数据类型

	fmt.Println("=================================")

	var s1 ystring
	s1 = "霸天虎"

	var s2 string = "威震天"
	fmt.Printf("s1 %s %T \n", s1, s1)
	fmt.Printf("s2 %s %T \n", s2, s2)

	fmt.Println("=================================")

	// 2.定义函数类型
	res := fun1()       /* fun1 返回值是一个函数 */
	res1 := res(10, 20) /* 调用返回的函数 */
	fmt.Println(res1)

	fmt.Println("=================================")

	// 3.别名
	var i3 yint1
	i3 = 100
	fmt.Printf("i3 %d %T \n", i3, i3) /* i3 100 int i3也是int*/

	i3 = i2 /* 都是int，赋值就不会报错 */
	fmt.Println(i3)
}

// 1.自定义类型
type yint int       /* yint 类型就是 int */
type ystring string /* ystring 类型就是 string */

// 2.定义函数类型，支持函数式编程
type yfun func(int, int) string

// 定义函数fun1 返回值是一个函数
func fun1() yfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

// 3.类型别名
type yint1 = int /* yint2和int可以混用 */
