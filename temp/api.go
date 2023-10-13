package temp

import (
	"os"

	xdl "github.com/mt1976/appFrame/dataloader"
)

// Path: temp/temp.go

// The package name is the last element of the import path: "appFrame/temp" has package name rand.
// The package name is used to qualify identifiers defined in that package in the import namespace:

type TempData struct {
	name string
	path string
	Data *xdl.Payload
}

var PathSeparator string = string(os.PathSeparator)

func init() {
	// ...

}

func Fetch(name string) (TempData, error) {
	return fetch(name)
}

func Store(t TempData) error {
	return store(t)
}
