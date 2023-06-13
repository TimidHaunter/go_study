//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%T \n", a)

	//函数数据类型 func()
	fmt.Printf("%T \n", fun1)

	//func(int) int
	fmt.Printf("%T \n", fun2)

	//func(float64, int, int) (int, int)
	fmt.Printf("%T \n", fun3)
	fmt.Println(fun3) //fun3函数体内存地址0xf6fa80，通过地址访问函数

	fmt.Println("=============================")

	var c func(int) int
	fmt.Printf("%T \n", c) //func(int)
	fmt.Println(c)         //nil

	//给c赋值
	c = fun2
	fmt.Println(c) //0x5ffb80

	fmt.Println("=============================")

	res1 := fun1
	res2 := fun1(1, 2)

	fmt.Println(res1) //0x5ffb80
	fmt.Println(res2)

	fmt.Println(res1(2, 2))
}

func fun1(a, b int) int {
	return a + b
}

func fun2(a int) int {
	return a
}

func fun3(a float64, b, c int) (int, int) {
	return b, c
}
