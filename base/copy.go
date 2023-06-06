//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
		深拷贝：拷贝的数据本身 值类型
		浅拷贝：拷贝的数据地址 引用类型
	*/

	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 0)
	fmt.Printf("slice2 len:%d cap:%d\n", len(slice2), cap(slice2))

	for i := 0; i < len(slice1); i++ {
		slice2 = append(slice2, slice1[i])
	}

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)

	fmt.Println("============================")

	//copy 深拷贝切片
	slice3 := []int{5, 6, 7}
	fmt.Println(slice2) //[1 2 3 4]
	fmt.Println(slice3) //[5 6 7]

	//s3 拷贝到 s2 从下标0开始
	copy(slice2, slice3)
	fmt.Println(slice2) //[5 6 7 4]
	fmt.Println(slice3)

	fmt.Println("============================")

	copy(slice2[1:], slice3[2:]) //拷贝一部分，粘贴一部分
	fmt.Println(slice2)
	fmt.Println(slice3)

}
