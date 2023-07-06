//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	map2 := make(map[string]float64)
	fmt.Printf("%T \n", map1)
	fmt.Printf("%T \n", map2)

	fmt.Println("============================")

	//map结构k是string v还是map
	map3 := make(map[string]map[string]string)
	m1 := make(map[string]string)
	m1["name"] = "Liuxi"
	m1["age"] = "29"
	m1["gender"] = "female"
	m1["address"] = "America"
	map3["userinfo:FIU679GIN"] = m1
	fmt.Println(map3)
	fmt.Printf("%T \n", map3)

	fmt.Println("============================")

	//map传递的是地址
	map5 := make(map[string]string)
	map5["name"] = "Liuxi"
	map5["age"] = "29"
	map5["gender"] = "female"
	map5["address"] = "America"

	map6 := map5
	fmt.Println(map5)
	fmt.Println(map6)

	fmt.Printf("map5地址:%p \n", map5)
	fmt.Printf("map6地址:%p \n", map6)
}
