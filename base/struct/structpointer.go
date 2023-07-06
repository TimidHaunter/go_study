//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 值传递：int、float、bool、string、array、struct
	 * 引用传递：slice、map、function、pointer
	 *
	 * 结构体本来是值传递，但是可以通过结构体指针实现引用传递
	 */

	// 结果体式值传递
	p1 := Person{"Yt", 31, "male", "中国"}
	fmt.Println(p1)
	fmt.Printf("%p, %T \n", &p1, p1) /* main.Person */

	p2 := p1
	fmt.Println(p2) /* {Yt 31 male 中国} */
	fmt.Printf("%p, %T \n", &p2, p2)
	// 操作p2 name字段
	p2.name = "Zx"
	fmt.Println(p2) /* {Zx 31 male 中国} */
	fmt.Printf("%p, %T \n", &p2, p2)

	fmt.Println("===============================")

	// 如果想实现结构体浅拷贝[引用传递]
	// 定义结构体指针
	var pp1 *Person /* 定义结构体指针 */
	pp1 = &p1       /* 赋值 */
	fmt.Println(pp1)
	fmt.Printf("%p, %T \n", &pp1, pp1) /* *main.Person */
	fmt.Println(*pp1)                  /* 指针存储地址对应的数据{Yt 31 male 中国} */

	fmt.Println("============ 浅拷贝 ============")

	(*pp1).name = "SongJ"
	fmt.Println(p1)
	pp1.name = "LuJY"
	fmt.Println(p1)

	fmt.Println("============ new() ============")

	// new() 专门创建某种类型指针
	pp2 := new(Person)
	fmt.Println(pp2)
	fmt.Printf("%p, %T \n", &pp2, pp2)

	(*pp2).name = "WuY"
	pp2.age = 3
	pp2.gender = "male"
	pp2.address = "水泊梁山"
	fmt.Println(pp2)

	fmt.Println("===============================")

	// new() 不是nil，空指针。是指向新分配的类型的内存空间，里面存储的零值
	pp3 := new(int)
	fmt.Printf("%T \n", pp3)
	pp4 := new(string)
	fmt.Printf("%T \n", pp4)
	pp5 := new([4]int)
	fmt.Printf("%T \n", pp5)
}

type Person struct {
	name    string
	age     int
	gender  string
	address string
}
