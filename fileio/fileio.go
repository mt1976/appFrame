package fileio

import (
	"os"

	xlogger "github.com/mt1976/appFrame/logs"
)

var xlogs xlogger.XLogger

func init() {
	xlogs = xlogger.New()
}

// FileExists returns true if the specified file existing on the filesystem
func fileExists(filename string) bool {
	return touch(filename)
}

// Touch returns true if the specified file existing on the filesystem
func touch(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
