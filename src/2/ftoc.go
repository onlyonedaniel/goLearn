package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g °F = %g °C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g °F = %g °C\n", boilingF, fToC(boilingF))

}

func fToC(f float64) float64 {
	// 摄氏度转华氏度
	return ((f - 32) * 5) / 9
}
