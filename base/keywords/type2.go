//go:build ignore
// +build ignore

package main

import "time"

func main() {

}

// 重命名time.Duration
// type yDuration = time.Duration /* 别名 */
type yDuration time.Duration /* 重新定义 */

// 给time.Duration加功能
func (yd yDuration) SimpleSet() { //cannot define new methods on non-local type time.Duration 在main包不能修改time包

}
