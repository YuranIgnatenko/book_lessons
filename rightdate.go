package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func is_right(d, m, y string) bool {
	m = strings.ToLower(m)
	const form1 = "02-jan-2006"
	const form2 = "2-jan-2006"
	const form3 = "02-янв-2006"
	const form4 = "2-янв-2006"
	const form5 = "02-январь-2006"
	const form6 = "2-январь-2006"
	const form7 = "02-01-2006"
	const form8 = "2-01-2006"
	const form9 = "02-1-2006"
	const form10 = "2-1-2006"
	const form11 = "02-january-2006"
	const form12 = "2-january-2006"
	list_forms := []string{
		form1, form2, form3, form4, form5,
		form6, form7, form8, form9, form10,
	}
	for _, format := range list_forms {
		_, err := time.Parse(format, d+"-"+m+"-"+y)
		if err != nil {
			continue
		} else {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) >= 4 {
		fmt.Println(is_right(os.Args[1], os.Args[2], os.Args[3]))
	}
}
