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
	after := flag.Int("after", 0, "Designate to return all annual reports after the specified year, including the year given.")
	before := flag.Int("before", 0, "Designate to return all annual reports before the specified year, including the year given.")
	recent := flag.Bool("recent", false, "Designate to return the most recent annual report available.")
	flag.Parse()

	utils.ValidateFlags(symbol, year, before, after, currentYear)

	url := utils.ApiUrlBuilder(symbol)

	data, err := utils.GetAnnualReportLink(url)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	reports, err := utils.GetAnnualReports(data, year)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	if *recent {
		reports = reports[:1]
	}

	if *open && !*recent && *year == 0 {
		fmt.Println(utils.ErrorColor + "WARNING: Please specify a year to utilize the '-open' flag." + utils.ResetColor)
		utils.DisplayReports(*symbol, reports, before, after)
		os.Exit(0)
	} else if *open {
		if err := utils.OpenDefaultBrowser(reports[0].Link); err != nil {
			fmt.Println("Error: There was an error automatically opening the browser.", err)
		}
	}

	utils.DisplayReports(*symbol, reports, before, after)
}
