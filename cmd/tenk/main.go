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
	symbol := flag.String("ticker", "", "Designate the desired US stock ticker symbol.")
	year := flag.Int("year", 0, "Designate the desired year for the annual report.")
	open := flag.Bool("open", false, "Designate to automatically open the report in the default browser. By default the requested report is returned to the shell.")
	flag.Parse()

	currentYear := time.Now().Year()
	if *symbol == "" || *year <= 1995 || *year > currentYear {
		log.Fatalln("Please provide a valid ticker symbol and year")
	}

	url := utils.ApiUrlBuilder(symbol)
	data, err := utils.GetAnnualReportLink(url)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	report, err := utils.ParseYearAnnualReport(data, year)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	if *open {
		err = utils.OpenDefaultBrowser(report.Link)
		if err != nil {
			fmt.Println("Error: There was an error automatically opening the browser.", err)
		}
		os.Exit(0)
	}

	utils.FormatTerminalOutput(report.Link, int(*year))
}
