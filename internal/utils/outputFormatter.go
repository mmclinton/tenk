package utils

import (
	"fmt"
	"strings"
)

func FormatTerminalOutput(url string, year int) {
	boxWidth := len(url) + 4 // 2 left 2 right

	fmt.Print("\n")
	fmt.Printf("Here is the requested 10-k for the year %v\n", year)
	fmt.Println("+" + strings.Repeat("-", boxWidth) + "+")
	fmt.Println("|" + strings.Repeat(" ", boxWidth) + "|")
	// fmt.Println("|" + strings.Repeat(" ", (boxWidth/2)-2) + fmt.Sprint(year) + strings.Repeat(" ", (boxWidth/2)-1) + "|")
	fmt.Printf("|  %s  |\n", url)
	fmt.Println("|" + strings.Repeat(" ", boxWidth) + "|")
	fmt.Println("+" + strings.Repeat("-", boxWidth) + "+")
	fmt.Print("\n")
}
