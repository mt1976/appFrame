package translate

// Get returns a translated string or the original string if no translation is found
func Get(in string) string {
	return get(in)
}

// Verbose turns on verbose logging
func Verbose() {
	setVerbose(true)
}

// Normal turns off verbose logging
func Normal() {
	setVerbose(false)
}

// GetValue returns translated string for a given key and name
func GetValue(key string, name string) string {
	return get(key + "$" + name)
}
