package utils

import (
	"log"
	"strconv"
)

func ValidateFlags(symbol *string, year *int, before *int, after *int, currentYear int) {
	if *symbol == "" || (*year != 0 && (*year < 1995 || *year > currentYear)) {
		log.Fatalln("Please provide a valid ticker symbol and year")
	}
	if (*after > currentYear) || (*before != 0 && *before < 1995) {
		log.Fatalln("Please provide a valid range between the years 1995 and " + strconv.Itoa(currentYear))
	}
}
