package generator

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringGenerator(length int) string {
	return stringWithCharset(length, charset)
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func CreateValue(column_type map[string]string) (string, error) {
	var valueString string

	for _, a := range column_type {
		switch a {
		case "integer":
			valueString = valueString + strconv.Itoa(rand.Intn(9999-0)) + ","
		case "text":
			valueString = valueString + "'" + stringGenerator(15) + "'" + ","
		default:
			return "", errors.New("the column did not match either text, nor integer")
		}
	}

	valueString = valueString[:len(valueString)-1]

	return valueString, nil
}
