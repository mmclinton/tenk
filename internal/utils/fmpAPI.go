package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SecFilings struct {
	Date string `json:"fillingDate"`
	Type string `json:"type"`
	Link string `json:"finalLink"`
}

func GetAnnualReportLink(url string) ([]SecFilings, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var allAnnualReports []SecFilings
	if err := json.Unmarshal(body, &allAnnualReports); err != nil {
		return nil, err
	}

	if len(allAnnualReports) == 0 {
		return nil, fmt.Errorf("no data found for the provided ticker")
	}

	return allAnnualReports, nil
}

func ParseYearAnnualReport(data []SecFilings, reportYear *int) (SecFilings, error) {
	for _, report := range data {
		reportFilingDate, err := time.Parse("2006-01-02 15:04:05", report.Date)
		if err != nil {
			fmt.Println("Could not parse data:", err)
		}
		year := reportFilingDate.Year()
		if year == *reportYear {
			return report, nil
		}
	}
	err := fmt.Errorf("this company did not file a 10-k for the year %v", *reportYear)
	return SecFilings{}, err
}
