package strings

import (
	"strings"
)

// Lowers the first character of a string
func LowerFirst(s string) string {
	return lowerFirst(s)
}

// ArrToString converts an array of strings to a printable string
func ArrToString(strArray []string) string {
	return strings.Join(strArray, "\n")
}

// StrArrayToString converst a string array into a string
func StrArrayToString(inArray []string) string {
	return strArrayToStringWithSep(inArray, "\n")
}

// StrArrayToStringWithSep converts a string array to a string using a given separator
func StrArrayToStringWithSep(inArray []string, inSep string) string {
	return strArrayToStringWithSep(inArray, inSep)
}

// RemoveSpecialChars removes special characters from a string and replaces them with a dash
func RemoveSpecialChars(in string) string {
	return removeSpecialChars(in)
}

// ReplaceWildcard replaces a wildcard {{wildcard}} with a given string
func ReplaceWildcard(orig string, replaceThis string, withThis string) string {
	return replaceWildcard(orig, replaceThis, withThis)
}

// PadRight pads a string on the right
func PadRight(s string, p string, l int) string {
	return padRight(s, p, l)
}

// PadLeft pads a string on the left
func PadLeft(s string, p string, l int) string {
	return padLeft(s, p, l)
}

// EncodeString encodes a string to base64 - same as Encode
// Deprecated: Use Encode instead
func EncodeString(rawStr string) string {
	return Encode(rawStr)
}

// Encode encodes a string to base64
func Encode(rawStr string) string {
	return encode(rawStr)
}

// DecodeString decodes a base64 encoded string - same as Decode
// Deprecated: Use Decode instead
func DecodeString(encodedStr string) string {
	return Decode(encodedStr)
}

// Decode decodes a base64 encoded string
func Decode(encodedStr string) string {
	return decode(encodedStr)
}

// DQuote wraps the string in double quotes - ""
func DQuote(s string) string {
	return dQuote(s)
}

// SQuote wraps the string in single quotes - ”
func SQuote(s string) string {
	return sQuote(s)
}

// DBracket wraps the string in double brackets - [[]]
func DBracket(s string) string {
	return dBracket(s)
}

// SBracket wraps the string in square brackets - []
func SBracket(s string) string {
	return sBracket(s)
}

// DChevrons wraps the string in double chevrons - <<>>
func DChevrons(s string) string {
	return dChevrons(s)
}

// SChevrons wraps the string in single chevrons - <>
func SChevrons(s string) string {
	return sChevrons(s)
}

// DCurlies wraps the string in double curlies - {{}}
func DCurlies(s string) string {
	return dCurlies(s)
}

// SCurlies wraps the string in single curlies - {}
func SCurlies(s string) string {
	return sCurlies(s)
}

// DParentheses wraps the string in double Parentheses - (())
func DParentheses(s string) string {
	return dParentheses(s)
}

// SParentheses wraps the string in single Parentheses - ()
func SParentheses(s string) string {
	return sParentheses(s)
}
