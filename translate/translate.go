package translate

import (
	"fmt"
	"strings"

	xio "github.com/mt1976/appFrame/fileio"
	xlogger "github.com/mt1976/appFrame/logs"
	xstrings "github.com/mt1976/appFrame/strings"
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
		out = outFormat(translationData[search])

	} else {

		search2 := strings.ToLower(in)
		search2 = strings.ReplaceAll(search2, " ", ".")
		if translationData[search2] != "" {
			out = outFormat(translationData[search2])
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

func outFormat(in string) string {

	fmt.Printf("in: %v\n", in)
	out := in

	out = strings.ReplaceAll(out, "\\n", "\n")
	out = strings.ReplaceAll(out, "\\t", "\t")
	out = strings.ReplaceAll(out, "\\r", "\r")
	out = strings.ReplaceAll(out, "null", "")
	out = xstrings.ReplaceWildcard(out, "null", "\n")
	out = xstrings.ReplaceWildcard(out, "eq", "=")
	out = xstrings.ReplaceWildcard(out, "gt", ">")
	out = xstrings.ReplaceWildcard(out, "lt", "<")
	out = xstrings.ReplaceWildcard(out, "amp", "&")
	out = xstrings.ReplaceWildcard(out, "apos", "'")
	out = xstrings.ReplaceWildcard(out, "quot", "\"")
	out = xstrings.ReplaceWildcard(out, "nbsp", " ")
	out = xstrings.ReplaceWildcard(out, "space", " ")
	//out = xstrings.RemoveSpecialChars(out)
	fmt.Printf("out: %v\n", out)
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
	xlogs.WithFields(xlogger.Fields{"test": "test"}).Warning("test")
	xlogs.Println("test")
}
