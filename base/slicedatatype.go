//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 切片数据类型 引用类型
	 */

	//数组
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1 //值传递，修改arr2元素，arr1没有变化
	fmt.Println(arr1, arr2)
	arr2[1] = 20
	fmt.Println(arr1, arr2)

	fmt.Println("============================")

	//切片 引用类型
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1
	fmt.Println(slice1, slice2)
	slice1[0] = 100 //修改了底层数组元素0的值，切片2是引用的也是底层数组
	fmt.Println(slice1, slice2)

	fmt.Println("============================")

	//切片指向数组，数组的内存的地址
	fmt.Printf("slice1:%p\n", slice1)
	fmt.Printf("slice2:%p\n", slice2)

	fmt.Println("============================")

	//切片的内存地址
	fmt.Printf("&slice1:%p\n", &slice1)
	fmt.Printf("&slice2:%p\n", &slice2)
}
