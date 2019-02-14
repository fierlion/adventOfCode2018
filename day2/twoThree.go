package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type twoThreeResult struct {
	twoCount   int
	threeCount int
}

func main() {
	twos := 0
	threes := 0
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		countResult := countTwosThrees(curString)
		twos += countResult.twoCount
		threes += countResult.threeCount
	}
	fmt.Printf("totalTwo: %d, totalThree: %d, total: %d\n", twos, threes, twos*threes)
}

func countTwosThrees(stringIn string) twoThreeResult {
	concordance := make(map[rune]int)
	twoRes := 0
	threeRes := 0
	for _, thisChar := range stringIn {
		if charCount, ok := concordance[thisChar]; ok {
			concordance[thisChar] = charCount + 1
		} else {
			concordance[thisChar] = 1
		}
	}
	for _, v := range concordance {
		if v == 2 {
			twoRes = 1
		}
		if v == 3 {
			threeRes = 1
		}
	}
	return twoThreeResult{twoRes, threeRes}
}
