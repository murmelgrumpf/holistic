package logging

import (
	"fmt"
	"strings"
)

func Sysout(params ...any) {
	Delimiter()
	if len(params) == 0 {
		fmt.Println("THIS IS A SYSOUT WAAAAA")
	}
	if len(params) == 1 {
		fmt.Printf("|%v|", params[0])
	}
	if len(params)%2 != 0 {
		fmt.Println(Red + "Sysout must have 0, 1 or an even number of params to display correctly. Always (string, any, string, any, ...)" + Reset)
	}
	_, err := fmt.Printf(strings.Repeat(" %s: |%v|\n", len(params)/2), params...)
	if err != nil {
		fmt.Println(Red+"Wrong Structure for sysout - "+Reset, err)
	}
	Delimiter()
}
