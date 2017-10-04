package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	ics "github.com/PuloV/ics-golang"
	toggl "github.com/jason0x43/go-toggl"
)

// Scheduler enters time entries into toggl based on .ics file
type Scheduler struct {
	token  string // api token
	path   string // path to .ics file
	hour   int    // time to update
	minute int
	second int
}

// begin scheduling. waits until first schedule time, then schedules, then waits
// 24 hours, then schedules, repeat...
func (sch *Scheduler) do() {

	t := time.Now()

	// the time when we want to do the initial scheduling
	firstScheduleTime := time.Date(t.Year(), t.Month(), t.Day(), sch.hour,
		sch.minute, sch.second, 0, t.Location())

	// d is the time until next schedule time
	d := firstScheduleTime.Sub(t)

	// if time until first schedule time is negative: add 24 hours (next day)
	if d < 0 {
		firstScheduleTime = firstScheduleTime.Add(24 * time.Hour)
		d = firstScheduleTime.Sub(t)
	}

	// wait until next time to schedule, then schedule
	for {
		wait(d, "Until next scheduletime")
		d = 24 * time.Hour
		sch.beginScheduling()
	}
}

func (sch *Scheduler) beginScheduling() {

	// get parser for the file we're working with
	parser := prepareParser(sch.path)

	// (try to) find out what lectures I had today
	events, err := lecturesAt(parser, time.Now())

	// if something went wrong, just give up... (continues running though)
	// TODO make it try again later
	if err != nil {
		fmt.Println("Failed to scheduleJobs: ", err.Error())
		return
	}

	// print today's events
	fmt.Println("Events for today:")
	for _, e := range events {
		fmt.Println("\t- ", e.GetDescription())
	}

	// open session and start entering the events
	session := toggl.OpenSession(sch.token)
	enterTimes(session, events)

	fmt.Println("DONE FOR TODAY")
}

// sleep, and enter time entries at the appropriate times
func enterTimes(session toggl.Session, events []*ics.Event) {

	// if there are no events, give up
	if len(events) == 0 {
		fmt.Println("no events")
		return
	}

	var err error

	// time until next toggl action
	diff := events[0].GetStart().Sub(time.Now())

	// for each event
	for i := range events {

		// sleep until start of event
		wait(diff, "Until start of next event")

		// start time entry (log)
		fmt.Println("\nSTART: ", events[i].GetDescription())
		id := getIDFromCode(events[i].GetSummary())
		var te toggl.TimeEntry
		if id != 0 {
			te, err = session.StartTimeEntryForProject(events[i].GetDescription(),
				id, false)
		} else {
			te, err = session.StartTimeEntry(events[i].GetDescription())
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		// sleep until end of event
		diff = events[i].GetEnd().Sub(time.Now())
		wait(diff, "Until end of this event")
		fmt.Println("\nEND: ", events[i].GetDescription())
		session.StopTimeEntry(te)

		// if this isn't the last event, calculate diff until start of event
		if i < len(events)-1 {
			diff = events[i+1].GetStart().Sub(time.Now())
		}
	}
}

// determine toggl id from IMT code found in ics file
func getIDFromCode(s string) int {
	switch s {
	case "IMT1362":
		return 62056917
	case "IMT2021":
		return 62056509
	case "IMT2571":
		return 62056716
	case "IMT2681":
		return 61803225
	default:
		return 0
	}
}

// Get lectures at particular time as ics events
func lecturesAt(parser *ics.Parser, when time.Time) ([]*ics.Event, error) {

	/*
		events := make([]*ics.Event, 3)

		events[0] = &ics.Event{}
		events[1] = &ics.Event{}
		events[2] = &ics.Event{}

		// set start, end (fake, rapid)
		now := time.Now()

		events[0].SetStart(now.Add(time.Second * 1))
		events[0].SetEnd(now.Add(time.Second * 10))
		events[0].SetDescription("Forelesning i ..")
		events[0].SetSummary("IMT1362")

		events[1].SetStart(now.Add(time.Second * 20))
		events[1].SetEnd(now.Add(time.Second * 30))
		events[1].SetDescription("Forelesning i nokke anna")
		events[1].SetSummary("IMT2021")

		events[2].SetStart(now.Add(time.Second * 40))
		events[2].SetEnd(now.Add(time.Second * 50))
		events[2].SetDescription("Forelesning i rare greier")
		events[2].SetSummary("IMT2681")

		return events, nil
	*/

	// get all calendars from parser
	cals, errCals := parser.GetCalendars()

	// if error or no calendars, error
	if errCals != nil {
		return nil, errCals
	} else if len(cals) == 0 {
		return nil, errors.New("No calendars (needed one)")
	}

	// get events for time 'when' (using first calendar)
	eventsForDay, errEvents := cals[0].GetEventsByDate(when)

	if errEvents != nil { // error -> error
		return nil, errEvents
	}

	// Filter out events that don't start with "Forelesning"
	// (To prevent toggl trakcking labs I'm not in) TODO: change?
	events := make([]*ics.Event, 0)
	for _, e := range eventsForDay {
		if strings.HasPrefix(e.GetDescription(), "Forelesning") {
			events = append(events, e)
			fmt.Println(when, " - ", e)
		}
	}
	return events, nil
}

// return parser ready to parse ics file pointed to by path
func prepareParser(path string) *ics.Parser {
	parser := ics.New()
	inputChan := parser.GetInputChan()
	inputChan <- path
	parser.Wait()

	return parser
}

// print formated message -> sleep d
func wait(d time.Duration, msg string) {
	fmt.Println("Sleeping (", d, "): ", msg)
	time.Sleep(d)
}
