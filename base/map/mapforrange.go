//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	 * map遍历 for range
	 */
	map1 := make(map[int]string)
	map1[1] = "菩提祖师"
	map1[2] = "唐僧"
	map1[3] = "猪八戒"
	map1[4] = "沙僧"
	map1[5] = "白龙马"
	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Printf("%d -> %s \n", k, v)
	}

	fmt.Println("============================")
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	//给keys排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

}
