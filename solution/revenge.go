package solution

import (
	"errors"
)

const (
	goodBoyOutput = "Good Boy"
	badBoyOutput  = "Bad boy"

	shotChar    = "S"
	revengeChar = "R"
)

func Revenge(shotSequence string) (string, error) {
	length := len(shotSequence)
	if isRevengeFirstWithNoShot(shotSequence) || isLastShotWithNoRevenge(shotSequence, length) {
		return badBoyOutput, nil
	}

	unrevengedCount, err := extractShotSequenceAndCount(shotSequence)
	if err != nil {
		return "", err
	}

	if unrevengedCount > 0 {
		return badBoyOutput, nil
	}

	return goodBoyOutput, nil
}

func isRevengeFirstWithNoShot(shotSequence string) bool {
	return shotSequence[0:1] == revengeChar
}

func isLastShotWithNoRevenge(shotSequence string, length int) bool {
	return shotSequence[length-1:] == shotChar
}

func extractShotSequenceAndCount(shotSequence string) (int, error) {
	unrevengedCount := 0
	for i := range shotSequence {
		char := shotSequence[i : i+1]
		if char == shotChar {
			unrevengedCount++
		} else if char == revengeChar && unrevengedCount != 0 {
			unrevengedCount--
		} else if char != revengeChar {
			return 0, errors.New("incorrect character in the sequence")
		}
	}

	return unrevengedCount, nil
}
