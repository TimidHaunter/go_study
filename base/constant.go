//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 常量定义：const identifier [type] = value
	 * 1.显式定义 const ONE_DAY int = 60*60*24
	 * 2.隐式定义 const KEY = "HDAURVAHD&*DADA"
	 *
	 * 定义后可以不使用，不能被修改
	 */
	const BAIDU string = "http://www.baidu.com"
	const PAI = "3.14"
	fmt.Println(BAIDU)
	fmt.Println(PAI)

	//cannot assign to PAI (untyped string constant "3.14")
	// PAI = "3.1415926"

	//定义一组常量
	const LENGTH, WIDTH, HIGHT = "10", "8", "6"
	fmt.Println(LENGTH, WIDTH, HIGHT)

	//常量零值，和上一个保持一致
	const (
		A int = 1
		B
	)
	fmt.Println(A, B)

	// const (
	// 	//第一个没有默认值
	// 	C
	// 	D int = 1
	// 	E
	// )
	// fmt.Println(C, D, E)

	//枚举
	const (
		MALE   = 1
		FEMALE = 2
		UNISEX = 3
	)
	fmt.Println(MALE, FEMALE, UNISEX)
}
