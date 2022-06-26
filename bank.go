package main

import (
	"fmt"
	"os"
	"strconv"
)

const percent = 10

func bank(deposit, year string) float64 {
	total := 0.0
	y, err := strconv.Atoi(year)
	if err != nil {
		panic(err)
	}
	d, err := strconv.ParseFloat(deposit, 64)
	if err != nil {
		panic(err)
	}
	total = d
	for i := 0; i < y; i++ {
		total += total / percent
	}
	return total
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println(bank(os.Args[1], os.Args[2]))
	}
}
