package main_test

import (
	"fmt"
	"testing"

	"github.com/paweenwatkwanja/boss-baby-revenge/solution"
	"github.com/paweenwatkwanja/boss-baby-revenge/utils"
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
			data, _ := utils.ReadDataFromFile(source)

			shotSequence := string(data)
			shotSequence = utils.ConvertInputToCorrectFormat(shotSequence)

			b.ResetTimer()
			b.ReportAllocs()

			solution.Revenge(shotSequence)
		})
	}
}
