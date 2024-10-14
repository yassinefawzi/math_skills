package main

import (
	"fmt"
	"os"
	"mathskills/mathskills"
	"math"
)

func main() {
	Holder, Err := os.ReadFile(os.Args[1])
	if Err != nil {
		fmt.Println(Err)
		return 
	}
	data := mathskills.Return_numbers(Holder)
	if len(data) == 0 {
		fmt.Println("Data Error.")
		return
	}
	fmt.Println(data)
	Average := mathskills.Average(data)
	Median := mathskills.Median(data)
	Variance := mathskills.Variance(data)
	Deviation := (math.Sqrt(Variance))
	fmt.Printf("Average: %.0f\n", Average)
	fmt.Printf("Median: %.0f\n", Median)
	fmt.Printf("Variance: %.0f\n", Variance)
	fmt.Printf("Standard Deviation: %.0f\n", Deviation)
}