package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
)

func main() {
	tp := flag.String("type", "radians", "Specifies whether to use radians or degrees")
	val := flag.Float64("value", 0, "The value to calculate")
	flag.Parse()

	switch *tp {
	case "radians":
		fmt.Println("\033[1m\033[33mTHIS IS CALCULATING IN \"RADIANS\"\033[0m")
		calculateRad(*val)
	case "degrees":
		fmt.Println("\033[1m\033[33mTHIS IS CALCULATING IN \"DEGREES\"\033[0m")
		calculateDeg(*val)
	default:
		fmt.Println("\033[31mPlease specify a viable type\033[0m")
	}
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

func calculateDeg(val float64) {
	pi := 3.14159265359
	deg := val * (pi / 180)
	calculateRad(deg)
}
