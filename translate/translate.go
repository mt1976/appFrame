package translate

import (
	"strings"

	xio "github.com/mt1976/appFrame/fileio"
	xlogs "github.com/mt1976/appFrame/logs"
)

var verbose bool

func init() {
	TRANSLATIONS, _ = xio.GetPropertiesFile("translate.dat")
	setVerbose(false)
}

func get(in string) string {
	//log.Info("TextGet: ", in)
	//log.Info("TextGet: ", lowerFirst(in)+"TXT")
	search := strings.ToLower(in)
	search = strings.ReplaceAll(search, " ", "")
	//out := str.LowerFirst(in) + "TODO"
	out := in
	if TRANSLATIONS[search] != "" {
		out = TRANSLATIONS[search]
	} else {

		search2 := strings.ToLower(in)
		search2 = strings.ReplaceAll(search2, " ", ".")
		if TRANSLATIONS[search2] != "" {
			out = TRANSLATIONS[search2]
		} else {

			if verboseLogging() {
				xlogs.WithFields(xlogs.Fields{"in": in, "search": search, "alternate": search2}).Warn("No Transalation Found")
				//log.Println("TextGet: No Translation for ", in)
			}

		}
	}
	//log.Info("TextGet: In :", in)
	//log.Info("TextGet: Out :", out)
	return out
}

func setVerbose(v bool) {
	verbose = v
}

func verboseLogging() bool {
	return verbose
}
