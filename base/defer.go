//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 使用defer关键字来延迟一个函数和方法的执行，参数还是原来的。不是延迟函数的调用
	 * 外围函数
	 */

	// defer fun1("Yin") //多个defer遵循栈，先进后出
	// fmt.Println("=====")
	// defer fun1("tian") //defer 在外围函数【main】执行完毕后，再执行defer函数
	// fmt.Println("=====")

	a := 1
	fmt.Println(a) //1
	defer fun2(a)  //1，只是延迟函数的执行
	a++
	fmt.Printf("main中a:%d \n", a) //2

	fmt.Println(fun3())
}

func fun1(s string) {
	fmt.Println(s)
}

func fun2(a int) {
	fmt.Printf("fun2中a:%d \n", a)
}

//当执行外围函数中的return语句时，只要其中所有的延迟函数都执行完毕，外围函数才会真正的return
//func3函数调用
//func1           defer
//0               return
func fun3() int {
	fmt.Println("func3函数调用")
	defer fun1("func1")
	return 0
}
