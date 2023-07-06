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
	fmt.Println("fun1()调用前a值：", a)
	fun1(a)
	fmt.Println("fun1()调用后a值：", a)

	fmt.Println("========================================")

	fun2(&a)
	fmt.Println("fun2()调用后a值：", a)

	fmt.Println("========================================")

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("fun3()调用前arr1值：", arr1)
	fun3(arr1)
	fmt.Println("fun3()调用后arr1值：", arr1)

	fmt.Println("========================================")

	fun4(&arr1)
	fmt.Println("fun4()调用后arr1值：", arr1)
}

// 参数是int
func fun1(num int) { /* num=a=10 */
	fmt.Println("fun1()中num值：", num)
	num = 100
	fmt.Println("fun1()中num值修改后：", num) /* num被销毁 */
}

// 参数是int的指针
func fun2(p1 *int) { /* p1=&a */
	fmt.Println("fun2()中，p1：", *p1)
	*p1 = 100 /* 操作p1对应元素的值 */
	fmt.Println("fun2()中，修改p1后：", *p1)
}

// 参数是数组
func fun3(arr [4]int) {
	fmt.Println("fun3()中，arr：", arr)
	arr[0] = 1000
	fmt.Println("fun3()中，修改arr后：", arr)
}

// 参数是数组指针
func fun4(p2 *[4]int) {
	fmt.Println("fun4()中，p2：", p2)
	p2[1] = 2000
	fmt.Println("fun4()中，修改p2后：", p2)
}
