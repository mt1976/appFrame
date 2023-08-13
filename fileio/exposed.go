package fileio

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jimlawless/cfg"
	xlogger "github.com/mt1976/appFrame/logs"
)

// The Copy function copies a file from one path to another and returns true if the copy was
// successful.
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

// The Read function reads the content of a file given its name and path, and returns the content as a
// string.
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

// The Write function writes content to a file specified by fileName and path, and returns a boolean
// indicating success and an error if any.
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

// The function `WriteData` writes the given content to a file with the specified name and path.
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
// The Touch function takes a filename as input and returns a boolean value indicating whether the file
// was successfully touched.
func Touch(filename string) bool {
	return touch(filename)
}

// Empty clears the contents of a specified directory
// The function "Empty" deletes all files in a given directory.
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
// The function `GetPropertiesFile` loads a properties file and returns its contents as a map of
// key-value pairs.
func GetPropertiesFile(fileName string) (map[string]string, error) {
	theseProperties := make(map[string]string)
	//machineName, _ := os.Hostname()
	//TODO For docker - if can't find properties file (create one from the template properties file)
	//propertiesFileName := "config/" + fileName

	xlogs.WithFields(xlogger.Fields{"File": fileName}).Info("Properties")

	err := cfg.Load(fileName, theseProperties)
	if err != nil {
		//xlogs.Fatal("cannot access properties file ["+propertiesFileName+"]", err)
		xlogs.WithFields(xlogger.Fields{"File": fileName, "Error": err.Error()}).Fatal("cannot access properties file")
	}
	return theseProperties, nil
}

func GetPropertiesPayload(fileName, extension, path string) (map[string]string, error) {

	xlogs.WithFields(xlogger.Fields{"File": fileName, "Path": path, "ext": extension}).Info("Properties X")

	if extension == "" {
		extension = "properties"
	}

	if fileName == "" {
		fileName = "system"
	}

	fileName = fileName + "." + extension

	propertiesFileName := "config" + string(os.PathSeparator) + fileName

	if len(path) != 0 {
		propertiesFileName = path + string(os.PathSeparator) + "config" + string(os.PathSeparator) + fileName
	}

	return GetPropertiesFile(propertiesFileName)
}
