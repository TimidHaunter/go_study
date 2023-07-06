//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/**
	 * 二维数组，元素是一维数组
	 * go 数组不够灵活呀
	 */
	//三个元素，每个元素都是一个一维数组，每个里面又有四个元素
	var arr = [3][4]int{
		{0, 1, 2, 3},
		{3, 2, 1, 2},
		{3, 2, 3, 2},
	}
	fmt.Println(arr)
	fmt.Println(arr[2][1]) //2

	//&arr = &arr[0] = &arr[0][0]
	//&arr[1] = $arr[1][0]
	//&arr[2] = $arr[2][0]
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p - %p - %p\n", &arr[0], &arr[1], &arr[2])
	fmt.Printf("%p - %p - %p\n", &arr[0][0], &arr[1][0], &arr[2][0])
	fmt.Printf("%p - %p - %p\n", &arr[0][1], &arr[1][1], &arr[2][1])

	//遍历二维数组
	for _, one := range arr {
		for _, two := range one {
			fmt.Println(two)
		}
	}
}
