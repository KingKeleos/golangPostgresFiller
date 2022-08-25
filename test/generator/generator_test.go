package generator_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/KingKeleos/golangPostgresFiller/src/generator"
	"github.com/stretchr/testify/assert"
)

func TestCreateValue(t *testing.T) {
	testMap := make(map[string]string)

	testMap["text"] = "text"
	testMap["value"] = "integer"

	genString, err := generator.CreateValue(testMap)
	assert.NoError(t, err)

	match, err := regexp.MatchString("(['a-zA-Z0-9' ,;]*[0-9]*)", genString)
	assert.NoError(t, err)

	fmt.Println(genString)
	fmt.Println(match)

	assert.True(t, match)
}

func TestWrongValues(t *testing.T) {
	testMap := make(map[string]string)

	testMap["text"] = "test"
	testMap["value"] = "integer"

	genString, err := generator.CreateValue(testMap)
	assert.Equal(t, "", genString)
	assert.Error(t, err)
	assert.EqualError(t, err, "the column did not match either text, nor integer")
}
