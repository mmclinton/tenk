package utils

import (
	"fmt"
	"strings"
)

const (
	resetColor = "\033[0m"
	cyanColor  = "\033[36m"
)

func FormatTerminalOutput(reports []SecFilings) {
	var output []string
	longestLine := 0

	for _, report := range reports {
		dateYear, err := GetYearFromDateString(report.Date)
		if err != nil {
			fmt.Println("Error parsing report date:", err)
			continue
		}
		line := fmt.Sprintf("%v: %s", dateYear, report.Link)
		output = append(output, line)

		if len(line) > longestLine {
			longestLine = len(line)
		}
	}

	boxWidth := longestLine + 4
	fmt.Println("\n+" + strings.Repeat("-", boxWidth) + "+")
	fmt.Println("|" + strings.Repeat(" ", boxWidth) + "|")

	for i, line := range output {
		if len(output) > 1 && i%2 == 1 {
			fmt.Printf("|  %s%-*s%s  |\n", cyanColor, longestLine, line, resetColor)
		} else {
			fmt.Printf("|  %-*s  |\n", longestLine, line)
		}
	}

	fmt.Println("|" + strings.Repeat(" ", boxWidth) + "|")
	fmt.Println("+" + strings.Repeat("-", boxWidth) + "+")
	fmt.Print("\n")
}
