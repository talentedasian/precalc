package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
)

func main() {
	val := flag.Float64("value", 0, "The value to calculate")
	flag.Parse()

	calculateRad(*val)
}

func calculateRad(val float64) {
	sin := math.Sin(val)
	cos := math.Cos(val)
	tan := math.Tan(val)
	cot := 1.0 / tan
	sec := 1.0 / cos
	csc := 1.0 / sin

	data := []float64{sin, cos, tan, cot, sec, csc}
	graph := asciigraph.Plot(data, asciigraph.Height(20))

	fmt.Print("\033[32m")
	fmt.Println(graph)

	fmt.Printf("Sine: %f\nCosine: %f\nTangent: %f\nCotangent: %f\nSecant: %f\nCosecant: %f\n\033[0m",
		sin, cos, tan, cot, sec, csc)
}
