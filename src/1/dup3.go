package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		// 返回的是byte slice ,需要转换为string
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1{
				fmt.Printf("%s ", filename)
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s\n", n, line)
		}
	}
}
