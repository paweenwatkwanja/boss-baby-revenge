package main

import (
	"fmt"
	"time"

	"github.com/paweenwatkwanja/boss-baby-revenge/solution"
	"github.com/paweenwatkwanja/boss-baby-revenge/utils"
)

func main() {
	var shotSequence string

	source := "sequence.txt"
	data, err := utils.ReadDataFromFile(source)
	if err != nil {
		panic(err)
	}
	shotSequence = string(data)

	if shotSequence == "" {
		shotSequence, err = utils.GetShotSequenceFromTerminal()
		if err != nil {
			panic(err)
		}
	}

	start := time.Now()

	shotSequence = utils.ConvertInputToCorrectFormat(shotSequence)
	outcome, err := solution.Revenge(shotSequence)
	if err != nil {
		panic(err)
	}

	fmt.Println(outcome)

	elapsed := time.Since(start)
	fmt.Printf("Duration: %v\n", elapsed)
}
