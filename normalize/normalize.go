// For more details about phone number normalizer implementation:
// https://exercism.io/tracks/go/exercises/phone-number/solutions/c777be4b16d44f4da9f5f702836a8470
package normalize

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	regCleaner = regexp.MustCompile(`[+()\s-.]+`)
	regPunc    = regexp.MustCompile(`\W`)
	regLetter  = regexp.MustCompile(`[^\d]`)
)

var (
	errLessThan10       = errors.New("incorrect number of digits")
	err11MustStartWith1 = errors.New("11 digits must start with 1")
	errMoreThan11       = errors.New("more than 11 digits")
	errPuncFound        = errors.New("punctuations not permitted")
	errLetterFound      = errors.New("letters not permitted")
	errAreaCode0        = errors.New("area code cannot start with zero")
	errAreaCode1        = errors.New("area code cannot start with one")
	errExchangeCode0    = errors.New("exchange code cannot start with zero")
	errExchangeCode1    = errors.New("exchange code cannot start with one")
)

// Number cleans and checks whether given number is valid or not,
// returns the normalized version.
func Number(inp string) (string, error) {
	var (
		num = regCleaner.ReplaceAllString(inp, "")
		siz = len(num)
		err error
	)
	if siz == 11 && num[0] == '1' {
		num = num[1:]
		siz--
	}

	switch {
	case siz < 10:
		err = errLessThan10
	case siz == 11 && num[0] != '1':
		err = err11MustStartWith1
	case siz > 11:
		err = errMoreThan11
	case regPunc.MatchString(num):
		err = errPuncFound
	case regLetter.MatchString(num):
		err = errLetterFound
	case num[0] == '0':
		err = errAreaCode0
	case num[0] == '1':
		err = errAreaCode1
	case num[3] == '0':
		err = errExchangeCode0
	case num[3] == '1':
		err = errExchangeCode1
	}
	return num, err
}

// Format will format the given phone number in a generic form,
// it works on the assumption that the input is normalized.
func Format(inp string) string {
	num, err := Number(inp)
	if err != nil {
		return fmt.Sprintf("%-20s, INVALID: %v", inp, err)
	}
	return fmt.Sprintf("(%v) %v-%v", num[:3], num[3:6], num[6:])
}
