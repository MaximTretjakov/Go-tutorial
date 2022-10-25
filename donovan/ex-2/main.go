package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]

	arg, err := strconv.ParseFloat(args, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ParseFloat: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Unit %.2f to Celsius: %.2f\n", arg, arg)
	fmt.Printf("Unit %.2f to Farenheit: %.2f\n", arg, func(arg float64) float64 { return 9 / 5 * (arg + 32) }(arg))
	fmt.Printf("Unit %.2f to Foot: %.2f\n", arg, func(arg float64) float64 { return arg * 0.3048 }(arg))
}
