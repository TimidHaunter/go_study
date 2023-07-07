//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 结构体模拟继承
	 */
	// 创建父类对象
	p1 := Parents{name: "关羽", age: 56}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	fmt.Println("==============================")

	s1 := Son{Parents{"关平", 22}, "蜀中高级中学"}
	fmt.Println(s1)

	s2 := Son{Parents: Parents{name: "关银屏", age: 16}, school: "蜀中初级中学"}
	fmt.Println(s2)

	var s3 Son
	s3.Parents.name = "关兴"
	s3.Parents.age = 20
	s3.school = "蜀中高级中学"
	fmt.Println(s3)

	// 提升字段，结构体中有一个字段是匿名字段，刚好这个匿名字段是一个匿名结构体，这个匿名结构体中的字段就是提升字段。可以直接访问。
	// s4.Parents.name => s4.name
	var s4 Son
	s4.name = "关索" /* 提升字段 */
	s4.age = 18    /* 直接访问 */
	s4.school = "蜀中高级中学"
	fmt.Println(s4)

	fmt.Println("==============================")
	fmt.Println(s1.name, s1.age, s1.school)
	fmt.Println(s2.name, s2.age, s2.school)
	fmt.Println(s3.name, s3.age, s3.school)
	fmt.Println(s4.name, s4.age, s4.school)
}

// 父类
type Parents struct {
	name string
	age  int
}

// 子类
type Son struct {
	Parents        /* 匿名字段，模拟继承 */
	school  string /* 子类新增属性 */
}
