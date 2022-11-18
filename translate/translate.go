package translate

import (
	"strings"

	xio "github.com/mt1976/appFrame/fileio"
)

func init() {
	TRANSLATIONS, _ = xio.GetPropertiesFile("translate.dat")
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
	}
	//log.Info("TextGet: In :", in)
	//log.Info("TextGet: Out :", out)
	return out
}
