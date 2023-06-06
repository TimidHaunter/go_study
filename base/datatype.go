//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 数据类型：基本类型，复合类型
	 *
	 * 布尔，数值，字符串
	 * 数值byte=uint8 rune=int32
	 *
	 * array，slice，map，function，pointer，struct，interface，channel
	 */

	//1.bool true false
	var b1 bool = true
	var b2 bool = false
	fmt.Printf("%T, %t\n", b1, b1)
	fmt.Printf("%T, %t\n", b2, b2)

	//2.整数，小数，复数
	var i1 int8 = 100
	fmt.Println(i1)
	//i1 = 200 //cannot use 200 (untyped int constant) as int8 value in assignment，超过取值范围

	var i2 uint8 = 200
	fmt.Println(i2)

	//一般选择int，看操作系统位数，32或者64位
	var i3 int = 1000
	fmt.Println(i3)

	// var i4 int64
	// i4 = i3 //cannot use i3 (variable of type int) as int64 value in assignment
	// fmt.Println(i4)

	//rune=int32
	var i5 rune = 2000
	var i6 int32 = 2000
	i5 = i6
	fmt.Println(i5, i6)

	//数值推断为int
	var i7 = 100
	fmt.Printf("%T, %d\n", i7, i7)

	//浮点，float32，float64
	var f1 float32 = 3.14
	var f2 float64 = 4.67
	fmt.Printf("%T, %.2f\n", f1, f1)
	fmt.Printf("%T, %.3f\n", f2, f2)
	fmt.Println(f1)
	//推断
	var f3 = 3.14
	fmt.Printf("%T, %f\n", f3, f3)

	//3.string
}
