//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 函数指针：指针，一个指向函数的指针
	 *   在go中，func默认看做指针，没有*
	 *   slice、map、function
	 *
	 * 指针函数：函数，一个返回指针的函数
	 */
	var a func() //定义一个函数
	a = fun1     //将fun1赋值给a，就是将fun1地址赋值给a
	a()

	arr2 := fun2()
	fmt.Printf("arr2的类型：%T，arr2的地址：%p，arr2的数值：%v \n", arr2, &arr2, arr2)

	arr3 := fun3()
	fmt.Printf("arr3的类型：%T，arr3的地址：%p，arr3的数值：%v \n", arr3, &arr3, arr3)
	fmt.Printf("arr3中存储的数组地址：%p \n", arr3)

}

func fun1() {
	fmt.Println("fun1()")
}

//==========================================================================

//普通函数，返回数组长度4，int
func fun2() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}

//指针函数，返回值是一个指针
func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("fun3中arr地址：%p \n", &arr)
	return &arr
}
