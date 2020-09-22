package main

import (
	"testing"
)

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
