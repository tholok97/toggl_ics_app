package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// get environment variables and start the scheduler

	var scheduleHour, scheduleMinute, scheduleSecond int
	var err1, err2, err3 error

	scheduleHour, err1 = strconv.Atoi(os.Getenv("SCHEDULE_HOUR"))
	if err1 != nil {
		fmt.Println("Env error: (", err1.Error(), ")")
	}

	scheduleMinute, err2 = strconv.Atoi(os.Getenv("SCHEDULE_MINUTE"))
	if err2 != nil {
		fmt.Println("Env error: (", err2.Error(), ")")
	}

	scheduleSecond, err3 = strconv.Atoi(os.Getenv("SCHEDULE_SECOND"))
	if err3 != nil {
		fmt.Println("Env error: (", err3.Error(), ")")
	}

	token := os.Getenv("API_TOKEN")
	if token == "" {
		fmt.Println("Env error: empty token")
	}

	path := os.Getenv("PATH_TO_ICS")
	if path == "" {
		fmt.Println("Env error: empty ics-path")
	}

	// print so it shows up in the logs
	fmt.Println("Schdule time: ", scheduleHour, ":", scheduleMinute, ":",
		scheduleSecond)
	fmt.Println("Token: ", token)
	fmt.Println("Path: ", path)

	// create scheduler with given properties
	sch := Scheduler{
		token:  token,
		path:   path,
		hour:   scheduleHour,
		minute: scheduleMinute,
		second: scheduleSecond,
	}

	// let the scheduler do it's thing
	fmt.Println("Running schduler....")
	sch.do()
}
