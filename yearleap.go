package main

import (
	"fmt"
	"os"
	"time"
)

func isleap(year string) bool {
	const shortForm = "2006-Jan-02"
	_, err := time.Parse(shortForm, year+"-Feb-29")
	if err != nil {
		return false
	}
	return true
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println(isleap(os.Args[1]))
	}
}
