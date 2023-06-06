//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 通过已有数组创建切片
	 * s := arr[startindex:endindex] 不包括endindex
	 * arr[:end] 从头到endindex
	 * arr[start:] 从startindex到尾
	 * len 从start到end切割的数据量 end-strart
	 * cap 从start到数组末尾 len(arr)-start
	 */
	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	fmt.Printf("arr:%p\n", &arr)
	fmt.Println("============================")

	slice := arr[1:4]
	fmt.Println(slice)
	fmt.Printf("len:%d,cap:%d\n", len(slice), cap(slice))
	fmt.Printf("slice:%p\n", slice)
	fmt.Println("============================")

	slice1 := arr[:3]
	slice2 := arr[2:]
	slice3 := arr[:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Printf("slice1:%p\n", slice1)
	fmt.Printf("slice2:%p\n", slice2)
	fmt.Printf("slice3:%p\n", slice3)

	fmt.Println("============================")

	//引用传递，数组的元素更改了，切片的元素都要跟着被修改，只存储了地址
	arr[4] = 50
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	fmt.Println("============================")
	//更改切片元素，数组和别的切片都被修改。都是操作的底层数组
	slice2[1] = 40
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	fmt.Println("============================")
	//切片追加元素，操作的也是底层数组
	slice1 = append(slice1, 400, 500, 600)
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	fmt.Println("============================")
	//切片追加元素，导致扩容，只影响slice1，扩容，导致底层数组扩容，地址变成新的数组地址。老数组没有变化，slice2和3还是老数组的地址
	fmt.Printf("len:%d,cap:%d\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 1001, 1002, 1003, 1004, 1005)
	fmt.Printf("len:%d,cap:%d\n", len(slice1), cap(slice1))
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
