package main

import (
	"fmt"
	"testing"
)

func BenchmarkRevenge(b *testing.B) {
	testCases := []string{
		"ten-sequences",
		"hundred-sequences",
		"thousand-sequences",
		"ten-thousand-sequences",
		"hundred-thousand-sequences",
		"million-sequences"}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("Test Case %v: %v", i+1, tc), func(b *testing.B) {
			source := fmt.Sprintf("./sequences/%v.txt", tc)
			data, _ := ReadDataFromFile(source)

			shotSequence := string(data)
			shotSequence = ConvertInputToCorrectFormat(shotSequence)

			b.ResetTimer()
			b.ReportAllocs()

			revenge(shotSequence)
		})
	}
}
