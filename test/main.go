package main

import (
	"image/png"
	"os"
	"runtime"

	"github.com/they4kman/gomandelbrot"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	image := gomandelbrot.Mandelbrot(1024, 768, 68, 0.9, 0)
	file, _ := os.Create("mandelbrot.png")
	defer file.Close()
	png.Encode(file, image)
}
