//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 匿名结构体：没有名字的结构体，在创建匿名结构体时，同时创建对象
	 * 匿名字段：一个结构体的字段没有字段名
	 *
	 *
	 */

	//创建结构体对象
	s1 := Student{4, "GongSS"}
	fmt.Println(s1)

	fmt.Println("==================================")

	// 通过匿名函数引出匿名结构体
	// 定义匿名函数
	fun1 := func() {
		fmt.Println("我是一个匿名函数~~~")
	}
	// 调用匿名函数
	fun1()

	fmt.Println("==================================")

	// 定义并且调用
	func() {
		fmt.Println("我是另一个匿名函数~~~")
	}()

	fmt.Println("==================================")

	// 匿名结构体
	s2 := struct {
		num  int
		name string
	}{ /* 定义完立马就要用，一次性使用，意义不大 */
		num:  5,
		name: "GuanS",
	}
	fmt.Println(s2)

	fmt.Println("==================================")

	s3 := Worker{6, "LinC"}
	fmt.Println(s3)
	// 匿名字段访问
	fmt.Println(s3.int)
	fmt.Println(s3.string)

}

// 定义普通结构体
type Student struct {
	num  int
	name string
}

// 匿名字段
type Worker struct {
	int /* 匿名字段，默认使用数据类型为名字，字段数据类型不能重复*/
	string
}
