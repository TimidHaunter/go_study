//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 函数返回值
	 */
	sum := getSum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sum)
}

//func 函数名(params) 类型
func getSum(nums ...int) int {
	fmt.Printf("%T \n", nums) //[]int 切片
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

//(sum int) 定义sum类型，sum=0就不需要定义了
func getSum2(nums ...int) (sum int) {
	fmt.Printf("%T \n", nums) //[]int 切片
	sum = 0
	for _, value := range nums {
		sum += value
	}
	return sum
}
