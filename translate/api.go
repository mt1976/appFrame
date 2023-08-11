package translate

// Get returns a translated string or the original string if no translation is found
// The Get function returns the result of the get function when given a string input.
func Get(property string) string {
	return get(property, "")
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

// GetValue returns translated string for a given property and name
// The GetValue function returns the value associated with a given property and name.
func GetValue(property string, category string) string {
	return get(property, category)
}

// The GetString function returns a string value for a given property.
func GetString(property string) (string, error) {
	return get(property, ""), nil
}

// The GetStringofCategory function returns a string value based on the given property and category.
func GetStringofCategory(property string, category string) (string, error) {
	return get(property, category), nil
}

// The GetInt function returns an integer value for a given property.
func GetInt(property string) (int, error) {
	return getInt(property, "")
}

// The function GetIntofCategory retrieves an integer value based on a given property and category.
func GetIntofCategory(property string, category string) (int, error) {
	return getInt(property, category)
}

// The GetBool function returns a boolean value based on a given property.
func GetBool(property string) (bool, error) {
	return getBool(property, "")
}

// The function GetBoolofCategory returns a boolean value based on the given property and category.
func GetBoolofCategory(property string, category string) (bool, error) {
	return getBool(property, category)
}

// The GetFloat function retrieves a float value for a given property.
func GetFloat(property string) (float64, error) {
	return getFloat(property, "")
}

// The GetFloatofCategory function retrieves a float value based on a given property and category.
func GetFloatofCategory(property string, category string) (float64, error) {
	return getFloat(property, category)
}

// The GetList function returns a list of strings based on a given property.
func GetList(property string) ([]string, error) {
	return getList(property, "")
}

// The function GetListofCategory returns a list of strings based on the given property and category.
func GetListofCategory(property string, category string) ([]string, error) {
	return getList(property, category)
}

// The function GetListString returns a list of strings based on a given property.
func GetListString(property string) ([]string, error) {
	return getList(property, "")
}

// The function GetListStringofCategory retrieves a list of strings based on a given property and
// category.
func GetListStringofCategory(property string, category string) ([]string, error) {
	return getList(property, category)
}

// The function GetListInt returns a list of integers and an error based on a given property.
func GetListInt(property string) ([]int, error) {
	return getListInt(property, "")
}

// The function GetListIntofCategory returns a list of integers based on the given property and
// category.
func GetListIntofCategory(property string, category string) ([]int, error) {
	return getListInt(property, category)
}

// The function GetListBool returns a list of boolean values and an error.
func GetListBool(property string) ([]bool, error) {
	return getListBool(property, "")
}

// The function GetListBoolofCategory returns a list of boolean values based on the given property and
// category.
func GetListBoolofCategory(property string, category string) ([]bool, error) {
	return getListBool(property, category)
}

// The function GetListFloat returns a list of float64 values and an error, based on a given property.
func GetListFloat(property string) ([]float64, error) {
	return getListFloat(property, "")
}

// The function GetListFloatofCategory returns a list of float64 values based on the given property and
// category.
func GetListFloatofCategory(property string, category string) ([]float64, error) {
	return getListFloat(property, category)
}

// The GetMap function returns a map of string key-value pairs and an error, based on the provided
// property.
func GetMap(property string) (map[string]string, error) {
	return getMap(property, "")
}

// The function GetMapofCategory returns a map of strings and an error based on the given property and
// category.
func GetMapofCategory(property string, category string) (map[string]string, error) {
	return getMap(property, category)
}
