package mathskills

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// \
func Return_numbers(str []byte) []int {
	var_hold := 0
	holder := strings.Split(string(str), "\n")
	data := []int{}
	var Err error
	for i := 0; i < len(holder); i++ {
		if holder[i] != "" {
			var_hold, Err = strconv.Atoi(holder[i])
			if Err != nil {
				fmt.Println(Err)
				return nil
			}
			data = append(data, var_hold)
		}
	}
	return data
}

func Average(data []int) float64 {
	var average float64
	for i := 0; i < len(data); i++ {
		average += float64(data[i])
	}
	return float64(average) / float64(len(data))
}

func Median(data []int) float64 {
	sort.Ints(data)
	var median float64
	i := (len(data) / 2) - 1
	if len(data)%2 == 0 {
		median = float64(data[i]+data[i+1]) / 2
	} else {
		if len(data) < 2 {
			i = 0
		}
		median = float64(data[i])
	}
	return (median)
}

func Variance(data []int) float64 {
	difference := []float64{}
	var variance float64
	average := Average(data)
	for i := 0; i < len(data); i++ {
		difference = append(difference, (float64(data[i]) - average) * (float64(data[i]) - average))
	}
	for i := 0; i < len(difference); i++ {
		variance += difference[i]
	}
	holder := float64(variance) / float64(len(data))
	return holder
}
