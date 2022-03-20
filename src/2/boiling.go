package main

import (
	"fmt"
)

const boolingF = 212.0

func main() {
	var f = boolingF
	var c = (f - 32) * 5 / 9
	fmt.Printf(" boinng point = %g fu °F or %g °C\n", f, c)
}
