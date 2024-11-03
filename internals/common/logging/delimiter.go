package logging

import "fmt"

func Delimiter() {
	fmt.Println(Bold + DarkGray + "___________________________________________________________________________________________________________" + Reset)
	Empty()
}

func Empty() {
	fmt.Println()
}
