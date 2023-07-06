//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 结构体：由一系列相同或者不相同类型的数据构成的数据集合
	 * 结构体成员由一系列的成员变量构成，这些成员变量也成为"字段"
	 * 类似PHP中的class
	 */

	// 初始化
	// 方法一
	var p1 Person
	fmt.Println(p1)
	// 结构体赋值
	p1.name = "Yint"
	p1.age = 31
	p1.gender = "male"
	p1.address = "中国北京崇文"
	fmt.Printf("姓名：%s，年龄：%d，性别：%s，地址：%s \n", p1.name, p1.age, p1.gender, p1.address)

	// 方法二
	p2 := Person{}
	p2.name = "Haunter"
	p2.age = 93
	p2.gender = "female"
	p2.address = "ghost"
	fmt.Printf("姓名：%s，年龄：%d，性别：%s，地址：%s \n", p2.name, p2.age, p2.gender, p2.address)

	// 方法三
	p3 := Person{name: "Psyduck", age: 54, gender: "male", address: "duck"}
	fmt.Printf("姓名：%s，年龄：%d，性别：%s，地址：%s \n", p3.name, p3.age, p3.gender, p3.address)

	// 方法四
	p4 := Person{"Raichu", 26, "male", "雷"}
	fmt.Printf("姓名：%s，年龄：%d，性别：%s，地址：%s \n", p4.name, p4.age, p4.gender, p4.address)
}

// 定义结构体
type Person struct {
	// 首字母大写，跨包可以访问；小写，本包访问
	name    string
	age     int
	gender  string
	address string
}
