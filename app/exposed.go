package app

import (
	"math/rand"
	"strconv"
	"time"
)

// SnoozeFor snoozes the application for a given amount of time
// The function SnoozeFor takes in a polling interval and calls the snooze function with that interval.
func SnoozeFor(inPollingInterval string) {
	snooze(inPollingInterval)
}

// Snooze snoozes for a random period
// The Snooze function generates a random number between 0 and 10 and then calls the snooze function
// with that number as a string argument.
func Snooze() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) // n will be between 0 and 10
	snooze(strconv.Itoa(n))
}
