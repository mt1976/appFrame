package dataloader

// The Payload type represents a set of data used for translation, including flags, a map of
// translations, and file-related information.
// @property {bool} verbose - A boolean flag indicating whether verbose logging should be enabled or
// not.
// @property data - The `data` property is a map of translations, where the keys are strings
// representing the original text and the values are strings representing the translated text. This map
// is shared by all instances of the translate package.
// @property {string} filename - The filename property is a string that represents the name of the
// file.
// @property {string} extension - The extension property represents the file extension of the payload
// file.
// @property {string} path - The "path" property is a string that represents the directory path where
// the file is located.
// @property {string} absolutepath - The `absolutepath` property represents the absolute path of a
// file.
type Payload struct {
	verbose      bool              // verbose is a flag to indicate verbose logging
	data         map[string]string // data is a map of translations, shared by all instances of the translate package
	filename     string
	extension    string
	path         string
	absolutepath string
}

// The function `New` returns a new instance of the data payload package
func New(filename, extension, path string) *Payload {
	return new(filename, extension, path)
}

// Get returns a translated string or the original string if no translation is found
// The Get function returns the result of the get function when given a string inpuP.
func (P *Payload) Get(property string) string {
	return P.getRawContent(property, "")
}

// Verbose turns on verbose logging
// The Verbose function sets the verbose mode to true.
func (P *Payload) Verbose() {
	P.setVerbose(true)
}

// Normal turns off verbose logging
// The function "Normal" sets the verbose flag to false.
func (P *Payload) Normal() {
	P.setVerbose(false)
}

// GetValue returns translated string for a given property and name
// The GetValue function returns the value associated with a given property and name.
func (P *Payload) GetValue(property string, category string) string {
	return P.getRawContent(property, category)
}

// The GetString function returns a string value for a given property.
func (P *Payload) GetString(property string) (string, error) {
	return P.getRawContent(property, ""), nil
}

// The GetStringofCategory function returns a string value based on the given property and category.
func (P *Payload) GetStringofCategory(property string, category string) (string, error) {
	return P.getRawContent(property, category), nil
}

// The GetInt function returns an integer value for a given property.
func (P *Payload) GetInt(property string) (int, error) {
	return P.getInt(property, "")
}

// The function GetIntofCategory retrieves an integer value based on a given property and category.
func (P *Payload) GetIntofCategory(property string, category string) (int, error) {
	return P.getInt(property, category)
}

// The GetBool function returns a boolean value based on a given property.
func (P *Payload) GetBool(property string) (bool, error) {
	return P.getBool(property, "")
}

// The function GetBoolofCategory returns a boolean value based on the given property and category.
func (P *Payload) GetBoolofCategory(property string, category string) (bool, error) {
	return P.getBool(property, category)
}

// The GetFloat function retrieves a float value for a given property.
func (P *Payload) GetFloat(property string) (float64, error) {
	return P.getFloat(property, "")
}

// The GetFloatofCategory function retrieves a float value based on a given property and category.
func (P *Payload) GetFloatofCategory(property string, category string) (float64, error) {
	return P.getFloat(property, category)
}

// The GetList function returns a list of strings based on a given property.
func (P *Payload) GetList(property string) ([]string, error) {
	return P.getList(property, "")
}

// The function GetListofCategory returns a list of strings based on the given property and category.
func (P *Payload) GetListofCategory(property string, category string) ([]string, error) {
	return P.getList(property, category)
}

// The function GetListString returns a list of strings based on a given property.
func (P *Payload) GetListString(property string) ([]string, error) {
	return P.getList(property, "")
}

// The function GetListStringofCategory retrieves a list of strings based on a given property and
// category.
func (P *Payload) GetListStringofCategory(property string, category string) ([]string, error) {
	return P.getList(property, category)
}

// The function GetListInt returns a list of integers and an error based on a given property.
func (P *Payload) GetListInt(property string) ([]int, error) {
	return P.getListInt(property, "")
}

// The function GetListIntofCategory returns a list of integers based on the given property and
// category.
func (P *Payload) GetListIntofCategory(property string, category string) ([]int, error) {
	return P.getListInt(property, category)
}

// The function GetListBool returns a list of boolean values and an error.
func (P *Payload) GetListBool(property string) ([]bool, error) {
	return P.getListBool(property, "")
}

// The function GetListBoolofCategory returns a list of boolean values based on the given property and
// category.
func (P *Payload) GetListBoolofCategory(property string, category string) ([]bool, error) {
	return P.getListBool(property, category)
}

// The function GetListFloat returns a list of float64 values and an error, based on a given property.
func (P *Payload) GetListFloat(property string) ([]float64, error) {
	return P.getListFloat(property, "")
}

// The function GetListFloatofCategory returns a list of float64 values based on the given property and
// category.
func (P *Payload) GetListFloatofCategory(property string, category string) ([]float64, error) {
	return P.getListFloat(property, category)
}

// The GetMap function returns a map of string key-value pairs and an error, based on the provided
// property.
func (P *Payload) GetMap(property string) (map[string]string, error) {
	return P.getMap(property, "")
}

// The function GetMapofCategory returns a map of strings and an error based on the given property and
// category.
func (P *Payload) GetMapofCategory(property string, category string) (map[string]string, error) {
	return P.getMap(property, category)
}

// The function Debug provide debug output for the current payload.
func (P *Payload) Debug() {
	P.debug()
}

func (P *Payload) GetLocalised(property, locale string) string {
	return P.getRawContent(property, locale)
}
