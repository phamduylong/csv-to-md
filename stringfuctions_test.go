package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/* STRING FUNCTION */
const STRINGS_SHOULD_BE_THE_SAME = "The two strings should be the same"

func TestDurationToString(t *testing.T) {
	d, _ := time.ParseDuration("3.0337h")
	expected := "3 hours 2 minutes 1 second 320 ms"
	res := durationToReadableString(d)

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestPadStart(t *testing.T) {
	originalString := "start"
	expected := "     start"
	res, err := padStart(originalString, 10, ' ')

	assert.Nil(t, err, "padStart should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestPadEnd(t *testing.T) {
	originalString := "end"
	expected := "end       "
	res, err := padEnd(originalString, 10, ' ')

	assert.Nil(t, err, "padEnd should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestPadCenterEven(t *testing.T) {
	originalString := "eleven"
	expected := "  eleven  "
	res, err := padCenter(originalString, 10, ' ')

	assert.Nil(t, err, "padCenter should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestPadCenterOdd(t *testing.T) {
	originalString := "eight"
	expected := "  eight   "
	res, err := padCenter(originalString, 10, ' ')

	assert.Nil(t, err, "padCenter should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

/* Conversion */
func TestConvertGeneric(t *testing.T) {
	testCSV := `First name,Last name,Email,Phone
Jane,Smith,jane.smith@email.com,555-555-1212
John,Doe,john.doe@email.com,555-555-3434
Alice,Wonder,alice@wonderland.com,555-555-5656`

	expected := `| First name | Last name |        Email         |    Phone     |
| :--------: | :-------: | :------------------: | :----------: | 
|    Jane    |   Smith   | jane.smith@email.com | 555-555-1212 |
|    John    |    Doe    |  john.doe@email.com  | 555-555-3434 |
|   Alice    |  Wonder   | alice@wonderland.com | 555-555-5656 |`

	res, err := Convert(testCSV, createGenericConfig())

	assert.Nil(t, err, "Convert should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)

}