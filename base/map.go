//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/**
	 * Map实现和定义
	 * 键值对，键不能重复，有点类似Redis的hash
	 * 引用类型
	 */

	//定义Map
	var map1 map[string]int //map1 只声明，没有初始化
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go": 98, "C++": 100, "PHP": 99, "Ruby": 97}
	fmt.Printf("map1:%T \n", map1)
	fmt.Printf("map2:%T \n", map2)
	fmt.Printf("map3:%T \n", map3)

	fmt.Println("============================")

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map3["PHP"])

	//引用类型的初始值是nil array和slice都是nil 类似PHP NULL

	//map1 只声明，没有初始化
	// map1["age"] = 1 //panic: assignment to entry in nil map

	fmt.Println("============================")

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	if map1 == nil {
		map1 = make(map[string]int)
		fmt.Println(map1 == nil)
	}

	fmt.Println("============================")

	map1["Yin"] = 1
	map1["Zeng"] = 2
	map1["Liu"] = 3
	map1["Yin"] = 4

	fmt.Println(map1)

	//判断map值是否存在
	value, ok := map1["Hu"]
	fmt.Println(value)
	fmt.Println(ok)
	if ok == true {
		fmt.Printf("对应的数值是：%d \n", value)
	} else {
		fmt.Printf("对应的数值是零值 \n")
	}

	fmt.Println("============================")

	//删除key
	fmt.Println(map1)
	delete(map1, "Yin")
	fmt.Println(map1)

	//批量添加
}
