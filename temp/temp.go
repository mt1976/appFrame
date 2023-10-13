package temp

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
	xdl "github.com/mt1976/appFrame/dataloader"
)

func fetch(name string) (TempData, error) {

	if name == "" {
		return TempData{}, errors.New("no temp filename specified")
	}

	ReturnData := TempData{}
	ReturnData.name = "test"
	ReturnData.path = PathSeparator + "temp" + PathSeparator + "test"
	ReturnData.Data = xdl.Payload{}

	spew.Dump(ReturnData)

	return ReturnData, nil
}

func store(t TempData) error {
	spew.Dump(t)
	return nil
}
