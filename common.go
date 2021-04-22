package common

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

// SnoozeFor snoozes the application for a given amount of time
func SnoozeFor(inPollingInterval string) {
	pollingInterval, _ := strconv.Atoi(inPollingInterval)
	log.Printf("Snoooze... Zzzzzz.... %d seconds...", pollingInterval)
	time.Sleep(time.Duration(pollingInterval) * time.Second)
}

// Snooze snoozes for a random period
func Snooze() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) // n will be between 0 and 10
	SnoozeFor(strconv.Itoa(n))
}
