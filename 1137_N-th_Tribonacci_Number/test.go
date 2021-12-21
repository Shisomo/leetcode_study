package main

import (
	"fmt"
	"strconv"
)

func main() {
	date := "dddd-01-04"
	result := dayOfYear(date)
	fmt.Printf("result is : %d", result)
}

func dayOfYear(date string) int {
	cemetery := map[string]int{"01": 0, "02": 31, "03": 59, "04": 90, "05": 120, "06": 151, "07": 181, "08": 212, "09": 243, "10": 273, "11": 304, "12": 334}
	year, _ := strconv.Atoi(date[0:4])
	day, _ := strconv.Atoi(date[8:10])
	result := cemetery[date[5:7]]
	result = result + day
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		if date[5:7] != "01" && date[5:7] != "02" {
			result++
		}
	}
	return result
}
