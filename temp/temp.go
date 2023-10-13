package temp

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/davecgh/go-spew/spew"
	xdl "github.com/mt1976/appFrame/dataloader"
	xio "github.com/mt1976/appFrame/fileio"
	xlg "github.com/mt1976/appFrame/logs"
)

var filename = "message"
var filetype = "dat"
var L xlg.XLogger

func init() {
	L = xlg.New()
	L.ToFileAndConsole("tempdata")
	L.Start()
}

// The fetch function takes a filename as input, checks if a "temp" folder exists, creates one if it
// doesn't, and returns the data from the specified file.
func fetch(name string) (TempData, error) {

	if name == "" {
		return TempData{}, errors.New("no temp filename specified")
	}

	filename = name

	storePath, err := os.Getwd()
	storePath = storePath + PathSeparator + "data"
	if err != nil {
		return TempData{}, errors.New("unable to determine current working directory")
	}
	//check the 'temp' folder exists, if not raise an error
	if !folderExists(storePath) {
		e := createFolder(storePath)
		if e != nil {
			return TempData{}, errors.New("no /data folder exists, attempted to auto-create, please create one manually - [" + e.Error() + "]")
		}
		if !folderExists(storePath) {
			return TempData{}, errors.New("no /data folder exists, please create one")
		}
	}

	// Assume all is ok, read the file and return the data

	L.Info("Reading " + filename + "." + filetype + " from " + storePath)

	ReturnData := TempData{}
	ReturnData.name = filename
	ReturnData.path = storePath
	ReturnData.folder = DataPath
	ReturnData.Data = xdl.New(filename, filetype, storePath)

	spew.Dump(ReturnData)

	return ReturnData, nil
}

func store(t TempData) error {

	currentUser, err := user.Current()
	if err != nil {
		return errors.New("unable to determine current user")
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error:", err)
		return errors.New("unable to determine current hostname")
	}

	updateTime := time.Now().Format("20060102@150405")

	t.Data.Update("updTimestamp", updateTime)
	t.Data.Update("updUser", currentUser.Username)
	t.Data.Update("updHost", hostname)

	spew.Dump(t)

	//fullFullPath := t.path + PathSeparator + t.name + "." + filetype

	fmt.Println(t.Data.ToString())

	//store to t.name + time.Now().String() + "." + t.filetype
	filename := t.name + "." + filetype

	L.Info("Storing to " + filename + " in " + t.folder)
	//xio.Touch(t.path + PathSeparator + filename)
	xio.Write(filename, t.folder, t.Data.ToString())

	return nil
}

// The function "folderExists" checks if a folder exists at the given path and returns a boolean value
// indicating its existence.
func folderExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// The function creates a folder at the specified path with permissions set to 0755.
func createFolder(folderPath string) error {
	// Create the folder with 0755 permissions (rwxr-xr-x)
	err := os.MkdirAll(folderPath, 0755)
	return err

}
