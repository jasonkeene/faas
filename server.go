package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"runtime"
)

func Index(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprintf(res, "Hello!")
}

func Fractal(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	query := req.URL.Query()
	width := query.Get("width")
	height := query.Get("height")
	i := query.Get("i")
	z := query.Get("z")
	content := fmt.Sprintf(`width: %d 
				 height: %d
				 i: %d
				 z: %v
				 seed: %v`, width, height, i, z)
	fmt.Fprintf(res, content)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	r := httprouter.New()
	r.GET("/", Index)
	r.GET("/fractal", Fractal)

	log.Fatal(http.ListenAndServe(":8080", r))
}
