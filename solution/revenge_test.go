package solution

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
			err := validateLength(tc.length)

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

func TestIsRevengeFirstWithNoShot(t *testing.T) {
	testCases := []struct {
		testCaseName   string
		shotSequence   string
		expectedOutput bool
	}{
		{"case with 'S' at the beginning", "SRRSR", false},
		{"case with 'R' at the beginning", "RSSSS", true},
	}
	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v : %v", i+1, tc.testCaseName)
		t.Run(testName, func(t *testing.T) {
			actualOutout := isRevengeFirstWithNoShot(tc.shotSequence)

			assert.Equal(t, tc.expectedOutput, actualOutout)
		})
	}
}

func TestIsLastShotWithNoRevenge(t *testing.T) {
	testCases := []struct {
		testCaseName   string
		shotSequence   string
		length         int
		expectedOutput bool
	}{
		{"case with 'R' at the end of the sequence", "SRRSR", 5, false},
		{"case with 'S' at the end of the sequence", "SRRSS", 5, true},
	}
	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v : %v", i+1, tc.testCaseName)
		t.Run(testName, func(t *testing.T) {
			actualOutout := isLastShotWithNoRevenge(tc.shotSequence, tc.length)

			assert.Equal(t, tc.expectedOutput, actualOutout)
		})
	}
}

func TestExtractShotSequenceAndCount(t *testing.T) {
	testCases := []struct {
		testCaseName   string
		shotSequence   string
		expectedOutput int
		err            error
	}{
		{"case success all shots are revenged", "SRRSR", 0, nil},
		{"case success some shots are not revenged", "SRSSSRSSRRR", 1, nil},
		{"case error the shot sequence has white space", "SRSSS RRR", 0, errors.New("incorrect character in the sequence")},
		{"case error the shot sequence has special character", "SRSSS%RRRS", 0, errors.New("incorrect character in the sequence")},
		{"case error the shot sequence haracter neither S nor R", "SRSSSARRRS", 0, errors.New("incorrect character in the sequence")},
		{"case error the shot sequence integer", "SRSSS1RRRS", 0, errors.New("incorrect character in the sequence")},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v : %v", i+1, tc.testCaseName)
		t.Run(testName, func(t *testing.T) {
			actualOutout, err := extractShotSequenceAndCount(tc.shotSequence)

			assert.Equal(t, tc.expectedOutput, actualOutout)
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
			actualResult, err := Revenge(tc.shotSequence)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, actualResult)
		})
	}
}
