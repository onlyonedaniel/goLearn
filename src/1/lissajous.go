package main

import(
	"fmt"
	"os"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UTC().UnixNano())

}

func lissajous(out io.Writer)  {
	const {
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
	}
}