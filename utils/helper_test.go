package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
