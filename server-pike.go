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
		w, err := strconv.Atoi(r.FormValue("w"))
		if err != nil {
			fmt.Fprint(wr, "w must be 32-bit integer")
			return
		}

		h, err := strconv.Atoi(r.FormValue("h"))
		if err != nil {
			fmt.Fprint(wr, "h must be 32-bit integer")
			return
		}

		i, err := strconv.Atoi(r.FormValue("i"))
		if err != nil {
			fmt.Fprint(wr, "i must be 32-bit integer")
			return
		}

		z, err := strconv.ParseFloat(r.FormValue("z"), 32)
		if err != nil {
			fmt.Fprint(wr, "z must be float")
			return
		}

		seed, err := strconv.ParseInt(r.FormValue("seed"), 10, 64)
		if err != nil {
			fmt.Fprint(wr, "seed must be 64-bit integer")
			return
		}

		wr.Header().Set("Content-Type", "image/png")

		m := gomandelbrot.Mandelbrot(w, h, i, float32(z), seed)
		png.Encode(wr, m)
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:9017", nil))
}
