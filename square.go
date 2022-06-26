package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func square(length string) {
	l, er := strconv.ParseFloat(length, 64)
	if er != nil {
		panic(er)
	}
	fmt.Println("Периметр : ", l*4.0)
	fmt.Println("Площадь : ", l*l)
	fmt.Println("Диагональ : ", math.Sqrt(2.0)*l)
}

func main() {
	if len(os.Args) > 1 {
		square(os.Args[1])
	}
}
