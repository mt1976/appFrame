package fileio

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jimlawless/cfg"
	xlogs "github.com/mt1976/appFrame/logs"
)

func Copy(fileName string, fromPath string, toPath string) bool {

	xlogs.Warning("Copying " + fileName + " from " + fromPath + " to " + toPath)

	content, err := Read(fileName, fromPath)
	if err != nil {
		xlogs.Fatal("File Read Error", err)
	}

	ok, err2 := Write(fileName, toPath, content)
	if err2 != nil {
		xlogs.Fatal("File Write Error", err2)
	}

	if !ok {
		xlogs.Fatal("Unable to Copy "+fileName+" from "+fromPath+" to "+toPath, nil)
	}

	return true
}

func Read(fileName string, path string) (string, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}

	// Check it exists - If not create it
	if !(Touch(filePath)) {
		WriteData(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		xlogs.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return string(content), err
}

func Write(fileName string, path string, content string) (bool, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		xlogs.Fatal("Write Error", err)
		return false, err
	}
	return false, nil
}

func WriteData(fileName string, path string, content string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		xlogs.Fatal(err)
		return -1
	}

	//	log.Println("File Write : " + fileName + " in " + path + "[" + filePath + "]")
	xlogs.Info(fileName, filePath)
	return 1
}

// Touch returns true if the specified file existing on the filesystem
func Touch(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Empty clears the contents of a specified directory
func Empty(dir string) error {
	xlogs.Println("TRASH", dir)
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		xlogs.Println(err)
		return err
	}
	//	fmt.Println("do Clear", files)
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			xlogs.Println(err)
			return err
		}
	}
	return err
}

// Load a Properties File
func GetPropertiesFile(fileName string) (map[string]string, error) {
	theseProperties := make(map[string]string)
	//machineName, _ := os.Hostname()
	// For docker - if can't find properties file (create one from the template properties file)
	propertiesFileName := "config/" + fileName

	xlogs.WithFields(xlogs.Fields{"File": fileName, "Path": propertiesFileName}).Info("Properties")

	err := cfg.Load(propertiesFileName, theseProperties)
	if err != nil {
		xlogs.Fatal("cannot access properties file "+propertiesFileName, err)
	}
	return theseProperties, nil
}