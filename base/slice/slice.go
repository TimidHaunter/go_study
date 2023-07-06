//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 切片 同数组类似，也叫变长数组或者动态数组
	 * 一个引用类容器，指向一个底层数组
	 * var var_name []type
	 */

	//定长数组
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Printf("%T\n", arr) //[4]int

	//定义
	var slice1 []int
	fmt.Println(slice1)
	fmt.Printf("%T\n", slice1) //[]int

	//定义
	slice2 := []int{}
	fmt.Println(slice2)

	//定义并赋值
	slice3 := []int{1, 2, 3, 4}
	fmt.Println(slice3)
	fmt.Printf("%T\n", slice3)
	fmt.Print("============================\n")

	/*
	 * make()也可以创建切片，make创建引用类型数据结构 slice map chan
	 * make(Type, Size, IntegerType)
	 * Type 类型 slice map chan
	 * Size 长度 len 实际存储元素数量
	 * IntegerType 容量 cap 最多存储元素数量
	 */

	slice4 := make([]int, 3, 8)
	fmt.Println(slice4)
	fmt.Printf("slice长度:%d,容量:%d\n", len(slice4), cap(slice4))

	//操作切片
	slice4[0] = 1
	slice4[1] = 2
	slice4[2] = 3
	fmt.Println(slice4)
	fmt.Print("============================\n")

	//append() 向切片末尾追加元素
	slice5 := make([]int, 0, 5)
	fmt.Println(slice5)
	fmt.Printf("slice长度:%d,容量:%d\n", len(slice5), cap(slice5))
	slice5 = append(slice5, 1, 2)
	fmt.Println(slice5)
	fmt.Printf("slice长度:%d,容量:%d\n", len(slice5), cap(slice5))
	slice5 = append(slice5, 3, 4, 5, 6) //追加的元素数量超过已有的容量5，自动扩容到10
	fmt.Println(slice5)
	fmt.Printf("slice长度:%d,容量:%d\n", len(slice5), cap(slice5))

	//追加slice4切片里面的元素
	slice5 = append(slice5, slice4...)
	fmt.Println(slice5)
	fmt.Printf("slice长度:%d,容量:%d\n", len(slice5), cap(slice5))
	fmt.Print("============================\n")

	//遍历切片
	for index, value := range slice5 {
		fmt.Printf("%d -> %d\n", index, value)
	}
}
