package translate

// Contains Nothing
var TRANSLATIONS map[string]string

func Get(in string) string {
	return get(in)
}

func Verbose() {
	setVerbose(true)
}

func Normal() {
	setVerbose(false)
}

func GetValue(key string, name string) string {
	return get(key + "$" + name)
}
