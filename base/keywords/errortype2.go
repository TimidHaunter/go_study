//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"net"
)

func main() {
	addr1, err1 := net.LookupHost("www.baidu.com")
	fmt.Println(addr1) /* [14.119.104.254 14.119.104.189] */
	fmt.Println(err1)  /* nil没错 */

	fmt.Println("===================================")

	addr2, err2 := net.LookupHost("www.baidusbhahah.com")
	fmt.Println(addr2) /* [] */
	fmt.Println(err2)  /* lookup www.baidusbhahah.com: no such host */

	if ins, ok := err2.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("操作超时")
		} else if ins.Temporary() {
			fmt.Println("临时性错误")
		} else {
			fmt.Println("默认错误")
		}
	}
} 
