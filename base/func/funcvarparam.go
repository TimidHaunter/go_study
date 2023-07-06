//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 可变参数，参数类型确定，数量不确定
	 * 参数名 ...参数类型
	 */

	// 求整数和，整数个数不确定
	// 函数中，最多只能用一个可变参数
	getSum(1, 2, 3, 4, 5, 6, 7)

	s1 := []int{1, 2, 3, 4, 5}
	getSum(s1...)

	fun1("Yin", "tian", 1.20, 123, 15.2, 83.2)
}

func getSum(nums ...int) {
	fmt.Printf("%T \n", nums) //[]int 切片
	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Printf("总和为：%d \n", sum)
}

//append 也是使用可变参数

func fun1(s1, s2 string, nums ...float64) {
	fmt.Println(s1, s2)
	fmt.Println(nums)
}
