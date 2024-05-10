package utils

import (
    "fmt"
    "strconv"
)

func PadStrNewLine(pad int, str string) string {
	middle := "%" + strconv.Itoa(pad) + "s"
	return fmt.Sprintf(middle, str) + "\n"
}

func PadStr(pad int, str string) string {
	middle := "%" + strconv.Itoa(pad) + "s"
	return fmt.Sprintf(middle, str)
}
