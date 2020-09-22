package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetDaysOfThisMonth() []string {
	thisYear := time.Now().Year()
	thisMonth := int(time.Now().Month())
	daysOfMonth := GetDaysOfMonth(thisYear)
	ret := []string{}
	for i := 1; i <= daysOfMonth[thisMonth-1]; i++ {
		ss := []string{strconv.Itoa(thisYear), fmt.Sprintf("%02d", thisMonth), fmt.Sprintf("%02d", i)}
		str := strings.Join(ss, "-")
		ret = append(ret, str)
	}

	nm := GetNextMonth(thisYear, thisMonth)
	//daysOfMonth := getDaysOfMonth(nm[0])
	for i := 1; i <= daysOfMonth[nm[1]-1]; i++ {
		ss := []string{strconv.Itoa(nm[0]), fmt.Sprintf("%02d", nm[1]), fmt.Sprintf("%02d", i)}
		str := strings.Join(ss, "-")
		ret = append(ret, str)
	}

	/*	for _, s := range ret {
		fmt.Println(s)
	}*/
	return ret
}

func GetNextMonth(year int, month int) [2]int {
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	t := date.AddDate(0, 1, 0)
	return [2]int{t.Year(), int(t.Month())}
}

func GetDaysOfMonth(year int) [12]int {
	daysOfLeapYear := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	daysOfNotLeapYear := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeapYear(year) {
		return daysOfLeapYear
	}
	return daysOfNotLeapYear
}

// isLeapYear is check leap year.
func isLeapYear(year int) bool {
	if year%100 == 0 && year%400 != 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}
