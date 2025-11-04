package utils

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLength(t *testing.T) {

	testCases := []struct {
		testCaseName string
		length       int
		err          error
	}{
		{"case error length is 0", 0, errors.New("length must be between 1 and 1,000,000")},
		{"case success length is 1", 1, nil},
		{"case success length is 1000000", 1000000, nil},
		{"case error length is 1000001", 1000001, errors.New("length must be between 1 and 1,000,000")},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v : %v", i+1, tc.testCaseName)
		t.Run(testName, func(t *testing.T) {
			err := ValidateLength(tc.length)

			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.err, nil)
			}
		})
	}
}

func TestConvertInputToCorrectFormat(t *testing.T) {
	testCases := []struct {
		testCaseName   string
		sequence       string
		expectedOutput string
	}{
		{"case sequence with white spaces", " SSSRRR ", "SSSRRR"},
		{"case sequencec with lower case", "sssrrr", "SSSRRR"},
		{"case sequence with lower case and white spaces", " sssrrr ", "SSSRRR"},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v", i+1)
		t.Run(testName, func(t *testing.T) {
			actualOutput := ConvertInputToCorrectFormat(tc.sequence)

			assert.Equal(t, tc.expectedOutput, actualOutput)
		})
	}
}
