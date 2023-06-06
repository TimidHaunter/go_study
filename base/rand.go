//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// num1 := rand.Int()
	// fmt.Println(num1)

	//Intn(n) [0,n) 不包括n
	for i := 0; i < 10; i++ {
		// num2 := rand.Intn(10)
		//Intn(n) 想从1开始
		// num3 := rand.Intn(10) + 1
		// fmt.Printf("%d %d\n", i, num2)
		// fmt.Printf("%d %d\n", i, num3)
	}

	var nowtime int64 = time.Now().UnixNano()
	fmt.Println("当前时间戳:", nowtime)

	//1.20 弃用rand.Seed()
	// rand.Seed(nowtime)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("->", rand.Intn(100))
	// }
	// fmt.Print("============================\n")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("->", rand.Intn(100))
	// }

	//获取指定区间随机数，[15,25] [0,10]+15
	//[a,b] Intn(b-a)+a
	for i := 0; i < 10; i++ {
		fmt.Println("->", rand.Intn(11)+15)
	}

}
