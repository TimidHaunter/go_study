//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 类型转换 Type(Value)
	 *
	 * 常量在有需要的时候，自动转换；变量自动转换
	 */
	var a int8 = 8
	var b int16 = int16(a)
	fmt.Printf("%T, %d\n", a, a)
	fmt.Printf("%T, %d\n", b, b)

	//不是所有类型都可以互相转化，布尔类就不能转换成int
	// b1 := true
	// b2 := int(b1) //cannot convert b1 (variable of type bool) to type int
	// fmt.Printf("%T, %d", b2, b2)

	var v1 float64 = 8.32
	sum := v1 + 10 //10，常量自动转换成浮点型
	fmt.Printf("%T, %f\n", sum, sum)
}
