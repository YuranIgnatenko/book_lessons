package main

import (
	"fmt"
	"os"
	"strconv"
)

func season(mounth string) string {
	if len(mounth) > 2 {
		panic("Need mounth from is one/two numbers (1,2,3...11,12)")
	}
	m_num, er := strconv.Atoi(mounth)
	if m_num < 1 || m_num > 12 {
		panic("Need mounth of range numbers (1,2,3...11,12)")
	}
	if er != nil {
		panic(er)
	}
	switch m_num {
	case 12, 1, 2:
		return "зима"

	case 3, 4, 5:
		return "весна"

	case 6, 7, 8:
		return "лето"

	case 9, 10, 11:
		return "осень"
	}
	return "не определенно"
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println(season(os.Args[1]))
	}
}
