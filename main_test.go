package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLength(t *testing.T) {
	t.Run("success: correct length", func(t *testing.T) {
		length := 1

		err := validateLength(length)

		assert.NoError(t, err)
	})

	t.Run("error: length less than one", func(t *testing.T) {
		length := 0

		err := validateLength(length)

		assert.Error(t, err)
		assert.EqualError(t, err, "unable to process")
	})

	t.Run("error: length greater than a million", func(t *testing.T) {
		length := 1000001

		err := validateLength(length)

		assert.Error(t, err)
		assert.EqualError(t, err, "unable to process")
	})
}

func TestIsRevengeFirstWithNoShot(t *testing.T) {
	t.Run("case shot first", func(t *testing.T) {
		shotSequence := "SRRSR"

		result := isRevengeFirstWithNoShot(shotSequence)

		assert.Equal(t, true, result)
	})

	t.Run("case revenge first with no shot", func(t *testing.T) {
		shotSequence := "RSSSS"

		result := isRevengeFirstWithNoShot(shotSequence)

		assert.Equal(t, true, result)
	})
}

func TestIsLastShotWithNoRevenge(t *testing.T) {
	t.Run("case last shot is revenged", func(t *testing.T) {
		shotSequence := "SRRSR"
		length := len(shotSequence)

		result := isLastShotWithNoRevenge(shotSequence, length)

		assert.Equal(t, false, result)
	})

	t.Run("case last shot is not revenged", func(t *testing.T) {
		shotSequence := "SRRSS"
		length := len(shotSequence)

		result := isLastShotWithNoRevenge(shotSequence, length)

		assert.Equal(t, true, result)
	})
}

func TestExtractShotSequenceAndCount(t *testing.T) {
	t.Run("success: all shots are revenged", func(t *testing.T) {
		shotSequence := "SRRSR"

		count := extractShotSequenceAndCount(shotSequence)

		assert.Equal(t, 0, count)
	})

	t.Run("success: some shots are not revenged", func(t *testing.T) {
		shotSequence := "SRSSSRSSRRR"

		count := extractShotSequenceAndCount(shotSequence)

		assert.Equal(t, 1, count)
	})
}

func TestRevengeWithProvidedTestCases(t *testing.T) {
	testCases := []struct {
		shotSequence   string
		expectedOutput string
	}{
		{"SRSSRRR", goodBoyOutput},
		{"RSSRR", badBoyOutput},
		{"SSSRRRRS", badBoyOutput},
		{"SRRSSR", badBoyOutput},
		{"SSRSRR", goodBoyOutput},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v: %v", i+1, tc.shotSequence)
		t.Run(testName, func(t *testing.T) {
			actualResult := revenge(tc.shotSequence)
			assert.Equal(t, tc.expectedOutput, actualResult)
		})
	}
}
