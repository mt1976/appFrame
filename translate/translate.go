package translate

import (
	"strings"

	fio "github.com/mt1976/appFrame/fileio"
)

// Contains Nothing
var TRANSLATIONS map[string]string

func init() {
	TRANSLATIONS, _ = fio.GetProperties("translate.dat")
}

func Get(in string) string {
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
