package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// strings:=""
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
