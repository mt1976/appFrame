package strings

import (
	"encoding/base64"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	xlogs "github.com/mt1976/appFrame/logs"
)

const (
	// DEFAULTDATEFORMAT is the date format used in Siena
	DEFAULTDATEFORMAT = "2006-01-02"
	// SPECIALCHARS is a list of special characters that are not allowed
	SPECIALCHARS = "[^A-Za-z0-9]+"
	// WILDCARDOPEN is the open wildcard
	WILDCARDOPEN = "{{"
	// WILDCARDCLOSE is the close wildcard
	WILDCARDCLOSE = "}}"
)

// Lowers the first character of a string
func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// ArrToString converts an array of strings to a printable string
func ArrToString(strArray []string) string {
	return strings.Join(strArray, "\n")
}

// StrArrayToString converst a string array into a string
func StrArrayToString(inArray []string) string {
	return StrArrayToStringWithSep(inArray, "\n")
}

// StrArrayToStringWithSep converts a string array to a string using a given separator
func StrArrayToStringWithSep(inArray []string, inSep string) string {

	outString := ""
	noRows := len(inArray)
	for ii := 0; ii < noRows; ii++ {
		outString += inArray[ii] + inSep
	}
	return outString
}

func RemoveSpecialChars(in string) string {
	reg, err := regexp.Compile(SPECIALCHARS)
	if err != nil {
		xlogs.Fatal(err)
	}
	newStr := reg.ReplaceAllString(in, "-")
	return newStr
}

// Replaces a Wildcard with a value
func ReplaceWildcard(orig string, replaceThis string, withThis string) string {
	wrkThis := WILDCARDOPEN + replaceThis + WILDCARDCLOSE
	//log.Printf("Replace %s with %q", wrkThis, withThis)
	return strings.ReplaceAll(orig, wrkThis, withThis)
}

func PadRight(s string, p string, l int) string {
	return rightPad2Len(s, p, l)
}
func PadLeft(s string, p string, l int) string {
	return leftPad2Len(s, p, l)
}

func EncodeString(rawStr string) string {

	// base64.StdEncoding: Standard encoding with padding
	// It requires a byte slice so we cast the string to []byte
	encodedStr := base64.URLEncoding.EncodeToString([]byte(rawStr))

	return encodedStr
}

func DecodeString(encodedStr string) string {

	decodedStr, err := base64.URLEncoding.DecodeString(encodedStr)
	if err != nil {
		xlogs.Panic("malformed input" + encodedStr)
	}

	return string(decodedStr)
}
