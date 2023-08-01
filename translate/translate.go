package translate

import (
	"strings"

	xio "github.com/mt1976/appFrame/fileio"
	xlogger "github.com/mt1976/appFrame/logs"
)

// verbose is a flag to indicate verbose logging
var verbose bool

var xlogs xlogger.XLogger

// translationData is a map of translations, shared by all instances of the translate package
var translationData map[string]string

// init loads the translate.dat file and sets verbose to false
func init() {
	xlogs = xlogger.New()
	xlogs.ToFileAndConsole("translate")
	translationData, _ = xio.GetPropertiesFile("translate.dat")
	setVerbose(false)
}

// Get returns a translated string or the original string if no translation is found
func get(in string) string {
	//log.Info("TextGet: ", in)
	//log.Info("TextGet: ", lowerFirst(in)+"TXT")
	search := strings.ToLower(in)
	search = strings.ReplaceAll(search, " ", "")
	//out := str.LowerFirst(in) + "TODO"
	out := in
	if translationData[search] != "" {
		out = translationData[search]
	} else {

		search2 := strings.ToLower(in)
		search2 = strings.ReplaceAll(search2, " ", ".")
		if translationData[search2] != "" {
			out = translationData[search2]
		} else {

			if verboseLogging() {
				xlogs.WithFields(xlogger.Fields{"using": in, "search": search, "alternate": search2}).Warn("No Translation")
				//log.Println("TextGet: No Translation for ", in)
			}

		}
	}
	//log.Info("TextGet: In :", in)
	//log.Info("TextGet: Out :", out)
	return out
}

// setVerbose sets the verbose flag
func setVerbose(v bool) {
	verbose = v
}

// verboseLogging returns the verbose flag
func verboseLogging() bool {
	return verbose
}

func test() {
	xlogs.Info("test")
	xlogs.Info("fruit")
	xlogs.Warning("test")
	xlogs.Println("test")
}
