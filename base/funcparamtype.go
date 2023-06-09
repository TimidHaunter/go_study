//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 函数参数传递
	 *
	 * -值传递
	 * int float string bool array
	 * -引用传递
	 * slice map chan
	 */
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前，数组的数据：", arr1) //[1 2 3 4]
	fun2(arr1)
	fmt.Println("函数调用后，数组的数据：", arr1) //[1 2 3 4]
	fmt.Printf("arr1：%p \n", &arr1)

	fmt.Println("=======================================")

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("函数调用前，切片的数据：", slice1) //[1 2 3 4]
	fun3(slice1)
	fmt.Println("函数调用后，切片的数据：", slice1) //[1000 2 3 4]
	fmt.Printf("slice1：%p \n", slice1)
}

//arr1 实参传递给 arr2 形参 arr2=arr1 数据赋值给arr2
func fun2(arr2 [4]int) {
	fmt.Println("函数调用中，数组的数据：", arr2) //[1 2 3 4]
	//在函数中修改数组的第一个元素
	arr2[0] = 100
	fmt.Println("函数调用中，修改数组元素后，数组的数据：", arr2) //[100 2 3 4]
	fmt.Printf("arr2：%p \n", &arr2)
}

func fun3(slice2 []int) {
	fmt.Println("函数调用中，切片的数据：", slice2) //[1 2 3 4]
	//在函数中修改切片的第一个元素
	slice2[0] = 1000
	fmt.Println("函数调用中，修改数组元素后，切片的数据：", slice2) //[1000 2 3 4]
	fmt.Printf("slice2：%p \n", slice2)
}
