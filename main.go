package main

import (
	"image/png"
	"os"
	"github.com/theY4Kman/gomandelbrot"
)


func main() {
	m := gomandelbrot.Mandelbrot(640, 480, 23, 2.3, 12345)
	f, err := os.Create("myshit.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, m)
}
