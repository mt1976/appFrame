package translate

import (
	fio "github.com/mt1976/AppFrame/fileio"
	str "github.com/mt1976/AppFrame/strings"
)

// Contains Nothing
var TRANSLATIONS map[string]string

func init() {
	TRANSLATIONS, _ = fio.GetProperties("translations")
}

func Get(in string) string {
	//log.Info("TextGet: ", in)
	//log.Info("TextGet: ", lowerFirst(in)+"TXT")
	out := str.LowerFirst(in) + "TODO"
	//log.Info("TextGet: In :", in)
	//log.Info("TextGet: Out :", out)
	return out
}
