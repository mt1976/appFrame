package translate

import (
	fio "github.com/mt1976/appFrame/fileio"
	str "github.com/mt1976/appFrame/strings"
)

// Contains Nothing
var TRANSLATIONS map[string]string

func init() {
	TRANSLATIONS, _ = fio.GetProperties("translate.dat")
}

func Get(in string) string {
	//log.Info("TextGet: ", in)
	//log.Info("TextGet: ", lowerFirst(in)+"TXT")
	out := str.LowerFirst(in) + "TODO"
	//log.Info("TextGet: In :", in)
	//log.Info("TextGet: Out :", out)
	return out
}
