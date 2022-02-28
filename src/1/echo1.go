package main

import (
	"fmt"
	"os"
)


func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// ++ 只能放到变量后面
	fmt.Println(s)
}
