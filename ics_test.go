package main

import (
	"testing"
)

func TestGetDaysOfMonth(t *testing.T) {
	actualLeapYear := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if actualLeapYear != getDaysOfMonth(2000) {
		t.Errorf("expected return an array of leap year.")
	}

	actualNotLeapYear := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if actualNotLeapYear != getDaysOfMonth(2001) {
		t.Errorf("expected return an array of not leap year.")
	}
}

func TestIsLeapYear(t *testing.T) {
	if !isLeapYear(2000) {
		t.Errorf("2000 was expected to be Leap year.")
	}
	if !isLeapYear(2004) {
		t.Errorf("2000 was expected to be Leap year.")
	}
	if isLeapYear(2001) {
		t.Errorf("2001 was not expected to be Leap year.")
	}
	if isLeapYear(2100) {
		t.Errorf("2001 was not expected to be Leap year.")
	}
}
