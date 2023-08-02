package banking

import (
	"fmt"
	"math/big"
	"strings"

	xmock "github.com/mt1976/appFrame/mock"
)

// isValidIBAN checks if the given IBAN is valid.
func isValidIBAN(iban string) bool {
	// Remove spaces and convert to uppercase
	iban = strings.ToUpper(strings.ReplaceAll(iban, " ", ""))

	// Check if the IBAN length is valid for the country code
	countryCode := iban[:2]
	countryInfo, err := xmock.GetCountryInfo(countryCode)
	if err != nil || len(iban) != countryInfo.IBANLength {
		return false
	}

	// Move the first 4 characters to the end
	iban = iban[4:] + iban[:4]

	// Convert characters to numbers (A = 10, B = 11, ..., Z = 35)
	var numericIBAN string
	for _, char := range iban {
		if '0' <= char && char <= '9' {
			numericIBAN += string(char)
		} else {
			numericIBAN += fmt.Sprintf("%d", int(char-'A'+10))
		}
	}

	// Convert numeric IBAN to a big.Int for modulo calculation
	bigIntIBAN, _ := new(big.Int).SetString(numericIBAN, 10)

	// Check if the modulo of the numeric IBAN with 97 is equal to 1
	return new(big.Int).Mod(bigIntIBAN, big.NewInt(97)).Int64() == 1
}

// isValidLEI checks if the given LEI is valid.
func isValidLEI(lei string) bool {
	// Remove spaces and convert to uppercase
	lei = strings.ToUpper(strings.ReplaceAll(lei, " ", ""))

	// Check if the LEI length is valid (it should be 20 characters)
	if len(lei) != 20 {
		return false
	}

	fmt.Printf("LEI: %s\n", leiToPrintable(lei))

	// Verify the LEI format using a regular expression (optional)
	// Implement a regular expression here to match the LEI format if needed.

	// Calculate the LEI checksum
	checksum := calculateLEIChecksum(lei[:18])

	// Compare the calculated checksum with the checksum provided in the LEI
	return checksum == lei[18:]

}

func leiToPrintable(lei string) string {
	lou := lei[:4]
	reserved := lei[4:6]
	entity := lei[6:18]
	checksum := lei[18:]
	return fmt.Sprintf("%s %s %s %s", lou, reserved, entity, checksum)
}

// calculateLEIChecksum calculates the LEI checksum using the ISO 7064 Mod 97-10 algorithm.
func calculateLEIChecksum(leiBase string) string {
	// Convert characters to numbers (A = 10, B = 11, ..., Z = 35)
	var numericLEI string
	for _, char := range leiBase {
		if '0' <= char && char <= '9' {
			numericLEI += string(char)
		} else {
			numericLEI += fmt.Sprintf("%d", int(char-'A'+10))
		}
	}

	// Calculate the checksum using modulo 97-10 algorithm
	remainder := 0
	for _, numStr := range strings.Split(numericLEI, "") {
		num := new(big.Int)
		num, _ = num.SetString(numStr, 10)
		remainder = (remainder*10 + int(num.Int64())) % 97
	}

	// Calculate the checksum value (98 - remainder) with leading zeros if necessary
	checksumValue := 98 - remainder
	checksum := fmt.Sprintf("%02d", checksumValue)

	return checksum

}
