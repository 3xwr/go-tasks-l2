package main

import "strings"

var monthList = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

func getMonthNumber(month string) int {
	month = strings.ToLower(month)
	for i := range monthList {
		if month == strings.ToLower(monthList[i]) {
			return i + 1
		}
	}
	return -1
}
