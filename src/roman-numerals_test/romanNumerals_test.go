package romannumeralstest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Table driven test
func TestDecimalToRoman (t *testing.T) {
	/*
		Anonymous struct

		1st curly brackets ---> definition of struct
		2nd curly brackets ---> definition of slice
		3rd curly brackets ---> initialisation of struct
	*/
	testSuite := []struct{
		name string
		input int
		res string
	}{
		{
			"4--->IV",
			4,
			"IV",
		},
		{
			"8--->VIII",
			8,
			"VIII",
		},
		{
			"4000--->\"\"",
			4000,
			"",
		},
		{
			"8000--->\"\"",
			8000,
			"",
		},
	}

	for _, value := range testSuite {
		t.Run(value.name, func(t *testing.T) {
			res, err := DecimalToRoman(value.input)
			assert.Equal(t, value.res, res)

			if value.res == "" {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

/* func TestDecimalToRoman (t *testing.T) {
	res, err := DecimalToRoman(4)
	assert.Equal(t, "IV", res)
	assert.Nil(t, err)
}

func TestDecimalToRoman_err (t *testing.T) {
	res, err := DecimalToRoman(4000)
	assert.Equal(t, "", res)
	assert.NotNil(t, err)
} */