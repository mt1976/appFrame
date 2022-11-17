package fileio

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jimlawless/cfg"
	logs "github.com/mt1976/AppFrame/logs"
)

// FileExists returns true if the specified file existing on the filesystem
func fileExists(filename string) bool {
	return FileTouch(filename)
}

func FileCopy(fileName string, fromPath string, toPath string) bool {

	logs.Warning("Copying " + fileName + " from " + fromPath + " to " + toPath)

	content, err := FileRead(fileName, fromPath)
	if err != nil {
		logs.Fatal("File Read Error", err)
	}

	ok, err2 := FileWrite(fileName, toPath, content)
	if err2 != nil {
		logs.Fatal("File Write Error", err2)
	}

	if !ok {
		logs.Fatal("Unable to Copy "+fileName+" from "+fromPath+" to "+toPath, nil)
	}

	return true
}

func FileRead(fileName string, path string) (string, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}

	// Check it exists - If not create it
	if !(FileTouch(filePath)) {
		FileSystem_WriteData(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		logs.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return string(content), err
}

func FileWrite(fileName string, path string, content string) (bool, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		logs.Fatal("Write Error", err)
		return false, err
	}
	return false, nil
}
func FileSystem_WriteData(fileName string, path string, content string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		logs.Fatal(err)
		return -1
	}

	//	log.Println("File Write : " + fileName + " in " + path + "[" + filePath + "]")
	logs.Info(fileName, filePath)
	return 1
}

// FileTouch returns true if the specified file existing on the filesystem
func FileTouch(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// RemoveContents clears the contents of a specified directory
func RemoveContents(dir string) error {
	logs.Println("TRASH", dir)
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		logs.Println(err)
		return err
	}
	//	fmt.Println("do Clear", files)
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			logs.Println(err)
			return err
		}
	}
	return err
}

// Load a Properties File
func GetProperties(fileName string) (map[string]string, error) {
	theseProperties := make(map[string]string)
	//machineName, _ := os.Hostname()
	// For docker - if can't find properties file (create one from the template properties file)
	propertiesFileName := "config/" + fileName

	logs.WithFields(logs.Fields{"File": fileName, "Path": propertiesFileName}).Info("Properties")

	err := cfg.Load(propertiesFileName, theseProperties)
	if err != nil {
		logs.Fatal("cannot access properties file "+propertiesFileName, err)
	}
	return theseProperties, nil
}
