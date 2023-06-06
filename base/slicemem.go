//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 切片引用类型数据，存储的是内存地址
	 * 创建切片的时候，先创建一个数组[1,2,3]，将数组地址赋值给切片
	 */
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	fmt.Printf("len:%d,cap:%d\n", len(slice1), cap(slice1))
	fmt.Printf("slice1:%p\n", slice1) //切片本身就是一个内存地址

	arr := [3]int{1, 2, 3}
	fmt.Printf("arr:%p\n", &arr) //数组是一个值传递类型，所以要用取址符
	fmt.Print("============================\n")

	//slice1 扩容
	slice1 = append(slice1, 4, 5) //将新的切片地址赋值给slice1
	fmt.Println(slice1)
	fmt.Printf("len:%d,cap:%d\n", len(slice1), cap(slice1))
	fmt.Printf("slice1:%p\n", slice1) //扩容后（翻倍），内存地址发生变化

	slice1 = append(slice1, 6, 7, 8)
	fmt.Println(slice1)
	fmt.Printf("len:%d,cap:%d\n", len(slice1), cap(slice1))
	fmt.Printf("slice1:%p\n", slice1) //扩容后，内存地址发生变化

	slice1 = append(slice1, 9, 10)
	fmt.Println(slice1)
	fmt.Printf("len:%d,cap:%d\n", len(slice1), cap(slice1))
	fmt.Printf("slice1:%p\n", slice1) //没扩容，内存地址【不变】
	fmt.Print("============================\n")

}
