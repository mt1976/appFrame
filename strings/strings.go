package strings

import (
	"encoding/base64"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	logs "github.com/mt1976/AppFrame/logs"
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
		logs.Fatal(err)
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

/*
* leftPad and rightPad just repoeat the padStr the indicated
* number of times
*
 */
func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}
func rightPad(s string, padStr string, pLen int) string {
	return s + strings.Repeat(padStr, pLen)
}

/* the Pad2Len functions are generally assumed to be padded with short sequences of strings
* in many cases with a single character sequence
*
* so we assume we can build the string out as if the char seq is 1 char and then
* just substr the string if it is longer than needed
*
* this means we are wasting some cpu and memory work
* but this always get us to want we want it to be
*
* in short not optimized to for massive string work
*
* If the overallLen is shorter than the original string length
* the string will be shortened to this length (substr)
*
 */
func rightPad2Len(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
func leftPad2Len(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
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
		logs.Panic("malformed input" + encodedStr)
	}

	return string(decodedStr)
}
