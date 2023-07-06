//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	sum := getSum(5)
	fmt.Println(sum)
}

//递归
func getSum(n int) int {
	//函数出口
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}
