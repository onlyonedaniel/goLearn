package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("%d\n", time.Now().Unix())
	rand.Seed(time.Now().Unix())
	// this is just for delay simulation
	duration := rand.Intn(10) + 1 //nolint
	fmt.Printf("%d\n", duration)
	time.After(time.Duration(duration) * time.Second)
	EndTime := uint64(time.Now().Unix())
	fmt.Printf("%d\n", EndTime)

}
