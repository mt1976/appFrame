package common

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
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

// GetCurrentFuncName will return the current function's name.
// It can be used for a better log debug system.(I'm NOT sure.)
func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
