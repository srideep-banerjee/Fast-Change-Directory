package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTagValid(t *testing.T) {
	tests := [] struct {
		input string
		expected bool
		expectedChar rune
	} {
		{input: "", expected: true, expectedChar: '0'},
		{input: " ", expected: false, expectedChar: ' '},
		{input: "fff$", expected: false, expectedChar: '$'},
		{input: "feerer", expected: true, expectedChar: '0'},
	}
	
	for _, test := range tests {
		actualBool, actualChar := IsTagValid(test.input)
		assert.Equalf(t, test.expected, actualBool, "IsTagValid(\"%s\") Expected %v Returned %v", test.input, test.expected, actualBool)
		assert.Equalf(t, test.expectedChar, actualChar, "IsTagValid(\"%s\") Expected '%c' Returned '%c'", test.input, test.expectedChar, actualChar)
	}
}

func TestIsCharAllowed(t *testing.T) {
	tests := [] struct {
		input rune
		expected bool
	} {
		{' ', false},
		{'a', true},
		{'z', true},
		{'m', true},
		{'A', true},
		{'Z', true},
		{'M', true},
		{'0', true},
		{'9', true},
		{'5', true},
		{'$', false}, 
	}
	
	for _, test := range tests {
		assert.Equalf(t, test.expected, isCharAllowed(test.input), "Input = %c", test.input)
	}
}