//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 1.创建map存储个人信息
	 * 2.每个map存储一个人的信息
	 * 3.将map存储到slice
	 * 4.遍历
	 */
	map1 := make(map[string]string)
	map1["name"] = "Yin"
	map1["age"] = "32"
	map1["gender"] = "male"
	map1["address"] = "Hanzhong,Hantai,Xinghan"
	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["name"] = "Zeng"
	map2["age"] = "29"
	map2["gender"] = "female"
	map2["address"] = "Xian,Yanta"
	fmt.Println(map2)

	map3 := make(map[string]string)
	map3["name"] = "Liuxi"
	map3["age"] = "29"
	map3["gender"] = "female"
	map3["address"] = "America"
	fmt.Println(map3)

	fmt.Println("============================")

	//将map存储到slice
	slice := make([]map[string]string, 0, 3)
	slice = append(slice, map1)
	slice = append(slice, map2)
	slice = append(slice, map3)
	fmt.Println(slice)

	//遍历
	for index, value := range slice {
		fmt.Println("============================")
		fmt.Printf("第%d个人信息是：\n", index)
		fmt.Printf("姓名：%s \n", value["name"])
		fmt.Printf("年龄：%s \n", value["age"])
		fmt.Printf("性别：%s \n", value["gender"])
		fmt.Printf("地址：%s \n", value["address"])
	}
}
