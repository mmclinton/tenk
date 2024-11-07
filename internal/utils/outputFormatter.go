package utils

import (
	"fmt"
	"strings"
)

const (
	ResetColor  = "\033[0m"
	cyanColor   = "\033[36m"
	ErrorColor  = "\033[31m"
	defaultLine = "Sorry, that range did not return any results!"
)

// why not add some abstraction while we're at it
func DisplayReports(symbol string, reports []SecFilings, before *int, after *int) {
	fmt.Printf("\nHere are the available 10-k reports for %v.", symbol)
	FormatTerminalOutput(reports, before, after)
}

func FormatTerminalOutput(reports []SecFilings, before *int, after *int) {

	output, longestLine := getReportRange(reports, before, after, 0)

	boxWidth := longestLine + 4
	fmt.Println("\n+" + strings.Repeat("-", boxWidth) + "+")
	fmt.Println("|" + strings.Repeat(" ", boxWidth) + "|")

	for i, line := range output {
		if len(output) > 1 && i%2 == 1 {
			fmt.Printf("|  %s%-*s%s  |\n", cyanColor, longestLine, line, ResetColor)
		} else {
			fmt.Printf("|  %-*s  |\n", longestLine, line)
		}
	}

	fmt.Println("|" + strings.Repeat(" ", boxWidth) + "|")
	fmt.Println("+" + strings.Repeat("-", boxWidth) + "+")
	fmt.Print("\n")
}

func getReportRange(reports []SecFilings, before *int, after *int, longestLine int) ([]string, int) {
	var output []string

	processReport := func(report SecFilings) bool {
		dateYear, err := GetYearFromDateString(report.Date)
		if err != nil {
			fmt.Println("Error parsing report date:", err)
			return false
		}

		if (*before == 0 || dateYear <= *before) && (*after == 0 || dateYear >= *after) {
			line := fmt.Sprintf("%v: %s", dateYear, report.Link)
			output = append(output, line)
			if len(line) > longestLine {
				longestLine = len(line)
			}
			return true
		}
		return false
	}

	for _, report := range reports {
		processReport(report)
	}

	if len(output) == 0 {
		longestLine = 45
		output = append(output, ErrorColor+defaultLine+ResetColor)
	}

	return output, longestLine
}
