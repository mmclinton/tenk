package main

import (
	"flag"
	"fmt"
	"log"
	"mmclinton/tenk/internal/utils"
	"os"
	"time"
)

func main() {
	currentYear := time.Now().Year()
	symbol := flag.String("ticker", "", "Designate the desired US stock ticker symbol.")
	year := flag.Int("year", 0, "Designate the desired year for the annual report. If no year is given, all annual reports are returned.")
	open := flag.Bool("open", false, "Designate to automatically open the report in the default browser. By default the requested report is returned to the shell.")
	flag.Parse()

	if *symbol == "" || (*year != 0 && (*year <= 1995 || *year > currentYear)) {
		log.Fatalln("Please provide a valid ticker symbol and year")
	}

	url := utils.ApiUrlBuilder(symbol)
	data, err := utils.GetAnnualReportLink(url)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	reports, err := utils.GetAnnualReports(data, year)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	if *open {
		err = utils.OpenDefaultBrowser(reports[0].Link)
		if err != nil {
			fmt.Println("Error: There was an error automatically opening the browser.", err)
		}
		utils.FormatTerminalOutput(reports)
		os.Exit(0)
	}

	if *year == 0 {
		fmt.Printf("\nHere are the available 10-k reports for %v.", *symbol)
		utils.FormatTerminalOutput(reports)
	} else {
		fmt.Printf("\nHere is the requested 10-k for the year %v.", *year)
		utils.FormatTerminalOutput(reports)
	}
}
