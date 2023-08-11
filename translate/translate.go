package translate

import (
	"errors"
	"strconv"
	"strings"

	xio "github.com/mt1976/appFrame/fileio"
	xlogger "github.com/mt1976/appFrame/logs"
	xstrings "github.com/mt1976/appFrame/strings"
)

// verbose is a flag to indicate verbose logging
var verbose bool

var xlogs xlogger.XLogger

// translationData is a map of translations, shared by all instances of the translate package
var translationData map[string]string

// init loads the translate.dat file and sets verbose to false
func init() {
	xlogs = xlogger.New()
	xlogs.ToFileAndConsole("translate")
	translationData, _ = xio.GetPropertiesFile("translate.dat")
	setVerbose(false)
}

// Get returns a translated string or the original string if no translation is found
func get(in string, kind string) string {
	//log.Info("TextGet: ", in)
	//log.Info("TextGet: ", lowerFirst(in)+"TXT")
	if in == "" {
		if verboseLogging() {
			xlogs.WithFields(xlogger.Fields{"using": in, "kind": kind}).Warn("Nothing to find")
			//log.Println("TextGet: No Translation for ", in)
		}

		return ""
	}

	search := strings.ToLower(setSearch(in, kind))
	search = strings.ReplaceAll(search, " ", "")
	//out := str.LowerFirst(in) + "TODO"
	out := in
	if translationData[search] != "" {
		out = outFormat(translationData[search])

	} else {

		search2 := strings.ToLower(in)
		search2 = strings.ReplaceAll(search2, " ", ".")
		if translationData[search2] != "" {
			out = outFormat(translationData[search2])
		} else {

			if verboseLogging() {
				xlogs.WithFields(xlogger.Fields{"using": in, "search": search, "alternate": search2}).Warn("No Translation")
				//log.Println("TextGet: No Translation for ", in)
			}

		}
	}
	//log.Info("TextGet: In :", in)
	//log.Info("TextGet: Out :", out)
	return out
}

func outFormat(in string) string {

	//fmt.Printf("in: %v\n", in)
	out := in

	out = strings.ReplaceAll(out, "\\n", "\n")
	out = strings.ReplaceAll(out, "\\t", "\t")
	out = strings.ReplaceAll(out, "\\r", "\r")
	out = xstrings.ReplaceWildcard(out, "null", "")
	out = strings.ReplaceAll(out, "null", "")
	out = xstrings.ReplaceWildcard(out, "eq", "=")
	out = xstrings.ReplaceWildcard(out, "gt", ">")
	out = xstrings.ReplaceWildcard(out, "lt", "<")
	out = xstrings.ReplaceWildcard(out, "amp", "&")
	out = xstrings.ReplaceWildcard(out, "apos", "'")
	out = xstrings.ReplaceWildcard(out, "quot", "\"")
	out = xstrings.ReplaceWildcard(out, "nbsp", " ")
	out = xstrings.ReplaceWildcard(out, "space", " ")
	//out = xstrings.RemoveSpecialChars(out)
	//fmt.Printf("out: %v\n", out)
	return out
}

// setVerbose sets the verbose flag
func setVerbose(v bool) {
	verbose = v
}

// verboseLogging returns the verbose flag
func verboseLogging() bool {
	return verbose
}

func test() {
	xlogs.Info("test")
	xlogs.Info("fruit")
	xlogs.Warning("test")
	xlogs.WithFields(xlogger.Fields{"test": "test"}).Warning("test")
	xlogs.Println("test")
}

// The setSearch function concatenates the property and category strings with a "$" separator if the
// category is not empty.
func setSearch(property string, category string) string {
	if category == "" {
		return property
	}
	if property == "" {
		return category
	}
	return property + "$" + category
}

// The function `getInt` converts a string value retrieved from a specific property and category into
// an integer, returning the integer value and any error encountered during the conversion.
func getInt(property string, category string) (int, error) {
	rtn, err := strconv.Atoi(get(property, category))
	if err != nil {
		return 0, err
	}
	return rtn, nil
}

// The function "getBool" parses a string value retrieved from a property and category, and returns a
// boolean value along with any error encountered during parsing.
func getBool(property string, category string) (bool, error) {
	rtn, err := strconv.ParseBool(get(property, category))
	if err != nil {
		return false, err
	}
	return rtn, nil
}

// The function "getFloat" parses a string value retrieved from a specific property and category into a
// float64 value.
func getFloat(property string, category string) (float64, error) {
	rtn, err := strconv.ParseFloat(get(property, category), 64)
	if err != nil {
		return 0, err
	}
	return rtn, nil
}

// The getList function takes in a property and category as parameters, retrieves a base string using
// the get function, splits the base string by commas, and returns the resulting list of strings.
func getList(property string, category string) ([]string, error) {
	if property == "" && category == "" {
		return nil, errors.New("property and category cannot both be empty")
	}
	if property == "" {
		return nil, errors.New("property cannot be empty")
	}
	base := get(property, category)
	if base == property {
		return nil, errors.New("property not found")
	}
	rtn := strings.Split(base, ",")
	return rtn, nil
}

// The function `getListInt` takes a property and category as input, retrieves a string value based on
// the property and category, splits the string into a slice of strings, converts each string element
// to an integer, and returns a slice of integers along with any error that occurred during the
// conversion.
func getListInt(property string, category string) ([]int, error) {
	base := get(property, category)
	if base == property {
		return nil, errors.New("property not found")
	}
	rtn := strings.Split(base, ",")
	var rtn2 []int
	for _, v := range rtn {
		rtn3, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		rtn2 = append(rtn2, rtn3)
	}
	return rtn2, nil
}

// The getListBool function takes a property and category as input, retrieves a comma-separated string
// value from a source, converts each value to a boolean, and returns a slice of booleans along with
// any error encountered.
func getListBool(property string, category string) ([]bool, error) {
	base := get(property, category)
	rtn := strings.Split(base, ",")
	var rtn2 []bool
	for _, v := range rtn {
		rtn3, err := strconv.ParseBool(v)
		if err != nil {
			return nil, err
		}
		rtn2 = append(rtn2, rtn3)
	}
	return rtn2, nil
}

// The function `getListFloat` takes a property and category as input, retrieves a comma-separated
// string value from a source, converts each value to a float64, and returns a slice of float64 values
// along with any error encountered during the conversion process.
func getListFloat(property string, category string) ([]float64, error) {
	base := get(property, category)
	if base == property {
		return nil, errors.New("property not found")
	}
	rtn := strings.Split(base, ",")
	var rtn2 []float64
	for _, v := range rtn {
		rtn3, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		rtn2 = append(rtn2, rtn3)
	}
	return rtn2, nil
}

// The function "getMap" takes a property and category as input, retrieves a base string using the
// "get" function, and returns a map of key-value pairs parsed from the base string.
func getMap(property string, category string) (map[string]string, error) {
	base := get(property, category)
	if base == property {
		return nil, errors.New("property not found")
	}
	rtn := make(map[string]string)
	for _, v := range strings.Split(base, ",") {
		rtn[strings.Split(v, ":")[0]] = strings.Split(v, ":")[1]
	}
	return rtn, nil
}
