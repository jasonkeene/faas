package main

import (
	"image/png"
	"net/http"
	"github.com/theY4Kman/gomandelbrot"
	"fmt"
	"log"
	"strconv"
)


func main() {
	http.HandleFunc("/mandelbrot", func(wr http.ResponseWriter, r *http.Request) {
		width, err := strconv.Atoi(r.FormValue("width"))
		if err != nil {
			fmt.Fprint(wr, "width must be 32-bit integer")
			return
		}

		height, err := strconv.Atoi(r.FormValue("height"))
		if err != nil {
			fmt.Fprint(wr, "height must be 32-bit integer")
			return
		}

		num_colors, err := strconv.Atoi(r.FormValue("colors"))
		if err != nil {
			fmt.Fprint(wr, "colors must be 32-bit integer")
			return
		}

		zoom, err := strconv.ParseFloat(r.FormValue("zoom"), 32)
		if err != nil {
			fmt.Fprint(wr, "zoom must be float")
			return
		}

		r_seed := r.FormValue("seed")
		var seed int64 = 0
		if r_seed != "" {
			seed, err = strconv.ParseInt(r_seed, 10, 64)
			if err != nil {
				fmt.Fprint(wr, "seed must be 64-bit integer")
				return
			}
		}

		wr.Header().Set("Content-Type", "image/png")

		m := gomandelbrot.Mandelbrot(width, height, num_colors, float32(zoom), seed)
		png.Encode(wr, m)
	})

	bind := "127.0.0.1:9017"
	log.Printf("Serving mandelbroten on %s ...", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}
