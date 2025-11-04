package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func GetShotSequenceFromTerminal() (string, error) {
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

func ConvertInputToCorrectFormat(input string) string {
	return strings.ToUpper(strings.TrimSpace(input))
}
