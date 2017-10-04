package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// (try to) return value of environment variable as string
func envAsString(name string) (string, error) {
	if val, exists := os.LookupEnv(name); exists {
		return val, nil
	} else {
		return "", errors.New("Env variable " + name + " does not exist")
	}
}

// (try to) return value of environment variable as int
func envAsInt(name string) (int, error) {
	if str, err1 := envAsString(name); err1 == nil {
		if val, err2 := strconv.Atoi(str); err2 == nil {
			return val, nil
		} else {
			return 923, err2
		}
	} else {
		return 923, err1
	}
}

func main() {

	// get environment variables and start the scheduler

	var scheduleHour, scheduleMinute, scheduleSecond int
	var err1, err2, err3, err4 error
	var token string

	if scheduleHour, err1 = envAsInt("SCHEDULE_HOUR"); err1 != nil {
		fmt.Println("Env error: (", err1.Error(), ")")
	}

	if scheduleMinute, err2 = envAsInt("SCHEDULE_MINUTE"); err2 != nil {
		fmt.Println("Env error: (", err2.Error(), ")")
	}

	if scheduleSecond, err3 = envAsInt("SCHEDULE_SECOND"); err3 != nil {
		fmt.Println("Env error: (", err3.Error(), ")")
	}

	if token, err4 = envAsString("API_TOKEN"); err4 != nil {
		fmt.Println("Env error: (", err4.Error(), ")")
	}

	// print so it shows up in the logs
	fmt.Println("Schdule time: ", scheduleHour, ":", scheduleMinute, ":",
		scheduleSecond)
	fmt.Println("Token: ", token)

	// create scheduler with given properties
	sch := Scheduler{
		token:  token,
		path:   "tholok_schedule.ics",
		hour:   scheduleHour,
		minute: scheduleMinute,
		second: scheduleSecond,
	}

	// let the scheduler do it's thing
	fmt.Println("Running schduler....")
	sch.do()
}
