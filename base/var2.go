//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 1.变量必须定义，才能被使用
	 * 2.变量的类型和赋值必须一致
	 * 3.同一个作用域内，变量名不能冲突
	 * 4.简短定义，左边变量至少有一个新变量
	 * 5.简短定义不能定义全局变量，【:=】必须在函数内使用
	 * 6.变量零值，就是默认值
	 * 7.变量定义了，就要使用，不然不能通过编译
	 */

	//定义变量，定义的过程就是开辟内存的过程
	var num int
	num = 100
	//&取地址符
	fmt.Printf("num的数值是：%d，地址是：%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是：%d，地址是：%p\n", num, &num)
	/*
		PS G:\data\Go\src\basic> go run .\var2.go
		num的数值是：100，地址是：0xc000012098
		num的数值是：200，地址是：0xc000012098
		num变量的地址没有变化
	*/

	//2
	var name string
	// name = 100 //cannot use 100 (untyped int constant) as string value in assignment
	// fmt.Println(name)
	name = "赵云"
	fmt.Println(name)

	//3
	// var name string = "张飞"
	// fmt.Println(name)

	//4.gender 新变量
	num, name, gender := 1000, "刘备", "男"
	fmt.Println(num, name, gender)

	//6.默认值
	var m int
	fmt.Printf("int的默认值是：%d\n", m)
	var f float64
	fmt.Printf("float64的默认值是：%f\n", f)
	var s string
	fmt.Printf("string的默认值是：%s\n", s)
	var sl []int
	fmt.Printf("切片的默认值是：%s\n", sl)
}
