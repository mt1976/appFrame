package app

import (
	"math/rand"
	"strconv"
	"time"
)

// SnoozeFor snoozes the application for a given amount of time
func SnoozeFor(inPollingInterval string) {
	snooze(inPollingInterval)
}

// Snooze snoozes for a random period
func Snooze() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) // n will be between 0 and 10
	snooze(strconv.Itoa(n))
}
