package app

import (
	"strconv"
	"time"

	xlogs "github.com/mt1976/appFrame/logs"
)

// SnoozeFor snoozes the application for a given amount of time
func snooze(inPollingInterval string) {
	pollingInterval, _ := strconv.Atoi(inPollingInterval)
	xlogs.Printf("Snoooze... Zzzzzz.... %d seconds...", pollingInterval)
	time.Sleep(time.Duration(pollingInterval) * time.Second)
}
