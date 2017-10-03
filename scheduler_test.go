package main

import (
	"testing"
	"time"
)

const (
	testParser = "test.ics"
)

func TestLecturesAt(t *testing.T) {

	// declare test variables
	when := time.Date(2017, 10, 18, 12, 0, 0, 0, time.Local)
	parser := prepareParser(testParser)

	// get events
	events, err := lecturesAt(parser, when)

	// check for error
	if err != nil {
		t.Error("Failed to get events: ", err.Error())
		return
	}

	// check empty events
	if len(events) == 0 {
		t.Error("No events returned")
		return
	}

	// validate events
	if events[0].GetSummary() != "IMT2021" ||
		events[0].GetSummary() != "IMT2021" {
		t.Error("Wrong event returned. (returned ", events[0].GetSummary(), ")")
	}
}
