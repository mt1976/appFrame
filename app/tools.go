package app

import (
	"strconv"
	"time"

	xlogger "github.com/mt1976/appFrame/logs"
)

// SnoozeFor snoozes the application for a given amount of time
// The function snooze takes in a polling interval as a string, converts it to an integer, and then
// sleeps for that amount of time.
func snooze(inPollingInterval string) {
	pollingInterval, _ := strconv.Atoi(inPollingInterval)
	xlogs := xlogger.New()
	xlogs.Printf("Snoooze... Zzzzzz.... %d seconds...", pollingInterval)
	time.Sleep(time.Duration(pollingInterval) * time.Second)
}
