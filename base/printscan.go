//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * https://studygolang.com/pkgdoc
	 *
	 * 输入输出 fmt
	 * Print   普通打印
	 * Printf  格式化打印
	 * Println 打印后换行
	 */

	/*
	 * %v 原样输出
	 * %T 数据类型
	 * %t bool类型
	 * %s 字符串
	 * %f 浮点
	 * %d 10进制的整数
	 * %b 2进制的整数
	 * %X 16进制 A-F
	 * %x 16进制 a-f
	 * %c 打印字符
	 * %p 内存地址
	 */
	a := 100
	b := 3.14
	c := "YIN"
	d := `Tian`
	e := 'A'
	f := false
	fmt.Printf("%T, %d\n", a, a)
	fmt.Printf("%T, %f\n", b, b)
	fmt.Printf("%T, %s\n", c, c)
	fmt.Printf("%T, %s\n", d, d)
	fmt.Printf("%T, %d, %c\n", e, e, e)
	fmt.Printf("%T, %t\n", f, f)
	fmt.Print("============================\n")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Print("============================\n")

	/*
	 * 输出 Scanln bufio
	 */
	var x int
	var y float64
	fmt.Println("请输入一个整数，一个小数")
	fmt.Scanln(&x, &y)
	fmt.Printf("x数值是%d,y的数值是%f", x, y)
}
