//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/**
	 * 空接口，(interface{})
	 * 不包含任何方法，所有的类型都实现了空接口，空接口可以存储任意类型的数值
	 *
	 * fmt包下Println现在是any任意
	 *
	 * 空接口当容器，可以存任意类型数据
	 */

	// 可以赋值任意类型给空接口
	var a1 A
	fmt.Println(a1) /* <nil> */
	a1 = Cat{"白色"}

	fmt.Println("==========================")

	var a2 A = Person{"曹操", 56}
	var a3 A = 3.14
	var a4 A = "曹昂"
	fmt.Println(a1) /* {白色} */
	fmt.Println(a2) /* {曹操 56} */
	fmt.Println(a3) /* 3.14 */
	fmt.Println(a4) /* 曹昂 */
	/* a1,a2,a3,a4都是空接口，可以赋值任意类型 */

	fmt.Println("==========================")

	// 参数是结构体
	test1(Cat{"白色"})
	// 参数是int
	test1(10)

	test2(a2)
	test2(a3)

	fmt.Println("==========================")

	// map容器 key string => value 任意
	map1 := make(map[string]interface{})
	map2 := make(map[string]any)

	map1["name"] = "曹丕"
	map1["age"] = 24
	map1["friend"] = Person{"司马懿", 30}

	map2["name"] = "曹植"
	map2["age"] = 22
	map2["friend"] = Person{"杨修", 26}

	fmt.Println(map1)
	fmt.Println(map2)

	fmt.Println("==========================")

	// slice容器
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, 1000, 53.6, "七步诗")

	slice2 := make([]any, 0, 10)
	slice2 = append(slice2, a1, a2, a3, a4, 200, 3.1415, "诗酒")
	fmt.Println(slice1)
	fmt.Println(slice2)
}

// 定义空接口，不包含任何方法
type A interface {
}

// 定义结构体
type Cat struct {
	color string
}

type Person struct {
	name string
	age  int
}

// 空接口当参数，可以接收任意参数
// a形参，A类型
func test1(a A) {
	fmt.Println("test1 =>", a)
}

// 简写
func test2(a interface{}) {
	fmt.Println("test2 =>", a)
}
