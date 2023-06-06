//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 数组的数据类型 长度+类型 [Size]Type
	 *
	 * 值传递，将数据复制其他变量
	 * int float string bool array
	 *
	 * 引用传递，将数据存储地址传递给其他变量
	 * slice map
	 */
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [3]float64{3.14, 2.5, 6.8}
	arr3 := [3]string{"Yin", "tian"}
	fmt.Printf("%T\n", arr1) //[4]int
	fmt.Printf("%T\n", arr2) //[3]float64
	fmt.Printf("%T\n", arr3) //[3]string

	//赋值
	num1 := 10
	num2 := num1
	fmt.Printf("%d - %d\n", num1, num2) //10 - 10

	//改变num2值，num1没影响，值传递，两个变量的内存地址是独立
	num2 = 20
	fmt.Printf("%d - %d\n", num1, num2)   //10 - 20
	fmt.Printf("%p - %p\n", &num1, &num2) //0xc0000a60a0 - 0xc0000a60a8
	fmt.Println("=========================")

	arr4 := arr1
	fmt.Println(arr1)
	fmt.Println(arr4)
	//数组也是值传递
	arr4[0] = 100
	fmt.Println(arr1)
	fmt.Println(arr4)
	fmt.Printf("%p - %p\n", &arr1, &arr4)
}
