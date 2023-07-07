//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/**
	 * 结构体嵌套
	 * 结构体的字段是另一个结构体
	 *
	 * 建议结构体指针嵌套，不会发生值传递，节省内存
	 */

	b1 := Book{}
	b1.name = "深入理解计算机系统"
	b1.price = 239.00

	s1 := Student{}
	s1.num = 7
	s1.name = "QinM"
	s1.book = b1

	fmt.Println(s1.num)
	fmt.Println(s1.name)
	fmt.Println(s1.book.name)
	fmt.Println(s1.book.price)
	fmt.Println(s1) /* {7 QinM {深入理解计算机系统 239}} */

	fmt.Println("====================================")

	s2 := Student{num: 8, name: "HuYZ", book: Book{name: "UNIX环境高级编程", price: 128.00}}
	fmt.Println(s2)

	fmt.Println("====================================")

	// 结果体是值传递 s1.book = b1 b1复制给s1.book
	s1.book.name = "JavaScript权威指南"
	fmt.Println(s1) /* {7 QinM {JavaScript权威指南 239}} */
	fmt.Println(b1) /* {深入理解计算机系统 239} */

	fmt.Println("====================================")

	//可不可以将b指针赋值给s
	b2 := Book{name: "C程序设计语言", price: 35.00}

	s3 := Student2{}
	s3.num = 7
	s3.name = "QinM"
	s3.book = &b2
	fmt.Println(b2)
	fmt.Println(s3.book)
	fmt.Println(s3)

	s3.book.name = "C程序设计语言 改"
	fmt.Println(b2)
	fmt.Println(s3)
}

type Book struct {
	name  string
	price float64
}

// book字段是结构体Book
type Student struct {
	num  int
	name string
	book Book
}

// book字段是结构体book指针
type Student2 struct {
	num  int
	name string
	book *Book
}
