package utils

import (
	"fmt"
	"strings"
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

	for _, line := range output {
		fmt.Printf("|  %-*s  |\n", longestLine, line)
	}

	fmt.Println("|" + strings.Repeat(" ", boxWidth) + "|")
	fmt.Println("+" + strings.Repeat("-", boxWidth) + "+")
	fmt.Print("\n")
}
