//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("[")
	// 关注特定的错误
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println("files:", files)
}
