package main

import (
	"testing"
	"time"

	ics "github.com/PuloV/ics-golang"
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

func TestSortEvents(t *testing.T) {

	// prepare test events
	events := make([]*ics.Event, 3)

	events[0] = ics.NewEvent()
	events[1] = ics.NewEvent()
	events[2] = ics.NewEvent()

	now := time.Now()
	events[0].SetStart(now)
	events[1].SetStart(now.Add(time.Hour * 2))
	events[2].SetStart(now.Add(time.Hour * 1))

	// try to sort
	sorted := sortEvents(events)
	if !(sorted[0].GetStart().Unix() < sorted[1].GetStart().Unix() && sorted[1].GetStart().Unix() < sorted[2].GetStart().Unix()) {
		t.Error("Wrong result returned from sort")
	}
}
