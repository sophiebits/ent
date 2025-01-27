package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var alphanumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz"
var symbols = "!#$%&'()*+,-./:;<=>?@[]\\\"^_`{|}~"
var numbers = "0123456789"
var lowerCase = "abcdefghijklmnopqrstuvwxyz"
var upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Generate a random integer code of length n
func GenerateRandCode(n int) string {
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(strconv.Itoa(rand.Intn(9)))
	}
	return sb.String()
}

// Generate a random alphanumeric string
func GenerateRandAlphaNumericKey(n int) string {
	var sb strings.Builder
	Die(generateRandKey(n, alphanumeric, &sb))
	return sb.String()
}

func GenerateRandDBName() string {
	var sb strings.Builder
	Die(generateRandKey(10, alphanumeric, &sb))
	str := strings.ToLower(sb.String())
	if unicode.IsNumber(rune(str[0])) {
		// add leading digit
		return string(lowerCase[rand.Intn(len(lowerCase))]) + str
	}
	return str
}

func generateRandKey(n int, charset string, sb *strings.Builder) error {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		if _, err := sb.WriteString(string(charset[rand.Intn(len(charset))])); err != nil {
			return err
		}
	}
	return nil
}

func GenerateRandEmail() string {
	return fmt.Sprintf("test-%s@email.com", GenerateRandCode(9))
}

func GenerateRandPassword() string {
	// generate a random password. used for tests
	// not going to win any awards
	var sb strings.Builder
	Die(generateRandKey(3, symbols, &sb))
	Die(generateRandKey(3, numbers, &sb))
	Die(generateRandKey(14, lowerCase, &sb))
	Die(generateRandKey(10, upperCase, &sb))

	result := []byte(sb.String())

	// "shuffle" everything so that order isn't as well defined
	for i := range result {
		idx := rand.Intn(len(result))
		// swap := result[idx]
		// result[idx] = result[i]
		// result[i] = swap
		result[i], result[idx] = result[idx], result[i]
	}
	return string(result)
}

func GenerateRandPhoneNumber() string {
	return GenerateRandCode(9)
}
