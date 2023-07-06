//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
		多返回值
	*/
	perimeter, area := rectangle(2.3, 4.3)
	fmt.Println("周长：", perimeter)
	fmt.Println("面积：", area)

	fmt.Println("=======================================")

	perimeter2, area2 := rectangle2(4, 6)
	fmt.Println("周长：", perimeter2)
	fmt.Println("面积：", area2)
}

//求矩形周边和面积
func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

func rectangle2(len, wid float64) (perimeter float64, area float64) {
	perimeter = (len + wid) * 2
	area = len * wid
	return perimeter, area
}
