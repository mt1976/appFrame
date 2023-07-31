package translate

// Get returns a translated string or the original string if no translation is found
// The Get function returns the result of the get function when given a string input.
func Get(in string) string {
	return get(in)
}

// Verbose turns on verbose logging
// The Verbose function sets the verbose mode to true.
func Verbose() {
	setVerbose(true)
}

// Normal turns off verbose logging
// The function "Normal" sets the verbose flag to false.
func Normal() {
	setVerbose(false)
}

// GetValue returns translated string for a given key and name
// The GetValue function returns the value associated with a given key and name.
func GetValue(key string, name string) string {
	return get(key + "$" + name)
}
