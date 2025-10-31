package main

import (
	"errors"
	"fmt"
)

func main() {
	var shotSequence string
	fmt.Print("Enter sequence of shots: ")
	fmt.Scanln(&shotSequence)
	fmt.Println(revenge(shotSequence))
}

const (
	goodBoyOutput = "Good Boy"
	badBoyOutput  = "Bad boy"

	shotChar    = "S"
	revengeChar = "R"
)

func revenge(shotSequence string) string {
	length := len(shotSequence)
	if err := validateLength(length); err != nil {
		return err.Error()
	}

	if isRevengeFirstWithNoShot(shotSequence) || isLastShotWithNoRevenge(shotSequence, length) {
		return badBoyOutput
	}

	count := extractShotSequenceAndCount(shotSequence)
	if count > 0 {
		return badBoyOutput
	}

	return goodBoyOutput
}

func validateLength(length int) error {
	if length < 1 || length > 1000000 {
		return errors.New("unable to process")
	}
	return nil
}

func isRevengeFirstWithNoShot(shotSequence string) bool {
	return shotSequence[0:1] == revengeChar
}

func isLastShotWithNoRevenge(shotSequence string, length int) bool {
	return string(shotSequence[length-1]) == shotChar
}

func extractShotSequenceAndCount(shotSequence string) int {
	count := 0
	for i := range shotSequence {
		char := shotSequence[i : i+1]
		if char == shotChar {
			count++
		} else if char == revengeChar && count != 0 {
			count--
		}
	}
	return count
}
