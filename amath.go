package main

import (
	"CtrlCmd"
	"fmt"
	"strings"
)

func f(x, y, op interface{}) string {
	n1 := fmt.Sprintf("%v", x)
	n2 := fmt.Sprintf("%v", y)
	oper := fmt.Sprintf("%v", op)
	res := CtrlCmd.LineGet(fmt.Sprintf(
		`echo 'print(eval("%v%v%v"))' > outfileoutgooutctrl.py && python3 outfileoutgooutctrl.py > out.f && cat out.f `, n1, oper, n2))
	CtrlCmd.LineGet("rm outfileoutgooutctrl.py && rm out.f")
	return fmt.Sprintf("%v", strings.TrimSpace(res))
}

func main() {
	fmt.Println(f(12, 13, "+"))
	fmt.Println(f(12, 2, "*"))
	fmt.Println(f(3, 2, "**"))
}
