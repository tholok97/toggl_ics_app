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
	var err error
	var token, path string

	if scheduleHour, err = envAsInt("SCHEDULE_HOUR"); err != nil {
		fmt.Println("Env error: (", err.Error(), ")")
	}

	if scheduleMinute, err = envAsInt("SCHEDULE_MINUTE"); err != nil {
		fmt.Println("Env error: (", err.Error(), ")")
	}

	if scheduleSecond, err = envAsInt("SCHEDULE_SECOND"); err != nil {
		fmt.Println("Env error: (", err.Error(), ")")
	}

	if token, err = envAsString("TOKEN"); err != nil {
		fmt.Println("Env error: (", err.Error(), ")")
	}

	if path, err = envAsString("TOKEN"); err != nil {
		fmt.Println("Env error: (", err.Error(), ")")
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
