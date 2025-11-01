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

func TestConvertInputToCorrectFormat(t *testing.T) {
	t.Run("case white spaces", func(t *testing.T) {
		shotSequence := " SSSRRR "

		result := ConvertInputToCorrectFormat(shotSequence)

		assert.Equal(t, "SSSRRR", result)
	})

	t.Run("case lower case", func(t *testing.T) {
		shotSequence := "sssrrr"

		result := ConvertInputToCorrectFormat(shotSequence)

		assert.Equal(t, "SSSRRR", result)
	})

	t.Run("case lower case with white spaces", func(t *testing.T) {
		shotSequence := " sssrrr "

		result := ConvertInputToCorrectFormat(shotSequence)

		assert.Equal(t, "SSSRRR", result)
	})
}

func TestIsRevengeFirstWithNoShot(t *testing.T) {
	t.Run("case shot first", func(t *testing.T) {
		shotSequence := "SRRSR"

		result := isRevengeFirstWithNoShot(shotSequence)

		assert.Equal(t, false, result)
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

		count, err := extractShotSequenceAndCount(shotSequence)

		assert.NoError(t, err)
		assert.Equal(t, 0, count)
	})

	t.Run("success: some shots are not revenged", func(t *testing.T) {
		shotSequence := "SRSSSRSSRRR"

		count, err := extractShotSequenceAndCount(shotSequence)

		assert.NoError(t, err)
		assert.Equal(t, 1, count)
	})

	t.Run("success: white space", func(t *testing.T) {
		shotSequence := "SRSSS RRR"

		count, err := extractShotSequenceAndCount(shotSequence)

		assert.Error(t, err)
		assert.EqualError(t, err, "incorrect character in the sequence")
		assert.Equal(t, 0, count)
	})

	t.Run("success: special character", func(t *testing.T) {
		shotSequence := "SRSSS%RRRS"

		count, err := extractShotSequenceAndCount(shotSequence)

		assert.Error(t, err)
		assert.EqualError(t, err, "incorrect character in the sequence")
		assert.Equal(t, 0, count)
	})

	t.Run("success: character neither S nor R", func(t *testing.T) {
		shotSequence := "SRSSSARRRS"

		count, err := extractShotSequenceAndCount(shotSequence)

		assert.Error(t, err)
		assert.EqualError(t, err, "incorrect character in the sequence")
		assert.Equal(t, 0, count)
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
			actualResult, err := revenge(tc.shotSequence)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, actualResult)
		})
	}
}
