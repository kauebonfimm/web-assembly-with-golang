package passwordengine

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GeneratePassword generates a password of length with the specified parameters
func GeneratePassword(length uint16, hasLetters, hasDigits, hasSymbos bool, removing string) (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	if !hasLetters && !hasDigits && !hasSymbos {
		return "test", fmt.Errorf("password must have at least one of the following: letters, digits, symbols")
	}

	var randParameters []byte

	if hasLetters {
		randParameters = append(randParameters, []byte(LETTERS)...)
	}

	if hasDigits {
		randParameters = append(randParameters, []byte(DIGITS)...)
	}

	if hasSymbos {
		randParameters = append(randParameters, []byte(SYMBOLS)...)
	}

	if removing != "" {
		for _, r := range removing {
			password := strings.ReplaceAll(string(randParameters), string(r), "")
			randParameters = []byte(password)
		}
	}

	var res string

	for {
		password := build(randParameters, length)

		if hasLetters {
			if !strings.ContainsAny(password, LETTERS) {
				continue
			}
		}

		if hasDigits {
			if !strings.ContainsAny(password, DIGITS) {
				continue
			}
		}

		if hasSymbos {
			if !strings.ContainsAny(password, SYMBOLS) {
				continue
			}
		}

		res = password
		break
	}

	return res, nil
}

func build(randParameters []byte, length uint16) string {
	password := make([]byte, length)

	for i := range password {
		password[i] = randParameters[rand.Intn(len(randParameters))]
	}

	return string(password)
}
