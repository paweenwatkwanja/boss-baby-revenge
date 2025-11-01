package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var shotSequence string

	source := "sequence.txt"
	data, err := ReadDataFromFile(source)
	if err != nil {
		panic(err)
	}
	shotSequence = string(data)

	if shotSequence == "" {
		shotSequence, err = getShotSequenceFromTerminal()
		if err != nil {
			panic(err)
		}
	}

	start := time.Now()

	shotSequence = ConvertInputToCorrectFormat(shotSequence)
	outcome, err := revenge(shotSequence)
	if err != nil {
		panic(err)
	}

	fmt.Println(outcome)

	elapsed := time.Since(start)
	fmt.Printf("Duration: %v\n", elapsed)
}

const (
	goodBoyOutput = "Good Boy"
	badBoyOutput  = "Bad boy"

	shotChar    = "S"
	revengeChar = "R"
)

func revenge(shotSequence string) (string, error) {
	length := len(shotSequence)
	if err := validateLength(length); err != nil {
		return "", err
	}

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

func ConvertInputToCorrectFormat(input string) string {
	return strings.ToUpper(strings.TrimSpace(input))
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

func getShotSequenceFromTerminal() (string, error) {
	result := ""

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter sequence of shots: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return result, err
	}
	result = string(input)

	return result, nil
}

func ReadDataFromFile(source string) ([]byte, error) {
	data, err := os.ReadFile(source)
	if err != nil {
		return nil, err
	}
	result := bytes.TrimSpace(data)

	return result, nil
}
