//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 存储一组相同数据类型的数据结构
	 */

	//定义数组，数组的内存地址就是数组内第一个元素的内存地址
	var arr1 [5]int
	fmt.Printf("arr1内存地址:%p\n", &arr1)
	fmt.Printf("arr1[0]内存地址:%p\n", &arr1[0])
	fmt.Printf("arr1[1]内存地址:%p\n", &arr1[1])

	//数组赋值
	var arr2 = [4]float32{3.14, 5.36, 9.36, 12.5}
	//访问数组
	fmt.Println(arr1[0])
	//赋值
	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1[1])
	fmt.Println(arr2[2])

	//容器中实际的存储的数据量
	fmt.Println("数组【arr1】的长度：", len(arr1))
	//容器中能够存储的最大数据量
	fmt.Println("数组【arr1】的容量：", cap(arr1))
	//因为数组定长，长度和容器相同
	fmt.Print("============================\n")
	fmt.Println("数组【arr2】的长度：", len(arr2))
	fmt.Println("数组【arr2】的容量：", cap(arr2))
	fmt.Println(arr2)

	//[...]自动获取长度
	fmt.Print("============================\n")
	arr3 := [...]int{1, 2, 34, 54, 356, 544}
	fmt.Println("数组【arr3】的长度：", len(arr3))
	fmt.Println(arr3)

	fmt.Print("============================\n")
	arr4 := [...]int{1: 1, 2: 2, 3: 34, 4: 54, 5: 356, 6: 544}
	fmt.Println("数组【arr4】的长度：", len(arr4))
	fmt.Println(arr4)
}
