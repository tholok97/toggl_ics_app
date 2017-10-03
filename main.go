package main

import (
	"fmt"
	"os"
	"strconv"
)

// what time to schedule the day's toggle jobs (defaults, overridden by envs)
var (
	scheduleHour   = 0
	scheduleMinute = 10
	scheduleSecond = 00
	path           = "tholok_schedule.ics"
	token          = "no token"
)

func main() {

	var err1, err2, err3 error
	scheduleHour, err1 = strconv.Atoi(os.Getenv("SCHEDULE_HOUR"))
	scheduleMinute, err2 = strconv.Atoi(os.Getenv("SCHEDULE_MINUTE"))
	scheduleSecond, err3 = strconv.Atoi(os.Getenv("SCHEDULE_SECOND"))
	token = os.Getenv("API_TOKEN")

	if err1 != nil || err2 != nil || err3 != nil || token == "" {
		panic("Error: schedule env variables wrong format")
	}

	fmt.Println("Schdule time: ", scheduleHour, ":", scheduleMinute, ":", scheduleSecond)
	fmt.Println("Token: ", token)

	sch := Scheduler{
		token:  token,
		path:   "tholok_schedule.ics",
		hour:   scheduleHour,
		minute: scheduleMinute,
		second: scheduleSecond,
	}

	fmt.Println("Running schduler....")
	sch.do()
}
