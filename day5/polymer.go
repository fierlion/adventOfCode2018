package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputString := ""
	for scanner.Scan() {
		inputString = scanner.Text()
	}

	// initialize minLen to maxInt
	minLen := int(^uint(0) >> 1)
	for i := 65; i < 91; i++ {
		capital := string(i)
		lower := string(i + 32)
		noCap := strings.Replace(inputString, capital, "", -1)
		thisString := strings.Replace(noCap, lower, "", -1)
		result := reducePolymer(thisString)
		if len(result) < minLen {
			minLen = len(result)
		}

	}
	fmt.Printf("%d\n", minLen)

}

func reducePolymer(strIn string) string {
	for {
		reacted := false
		for i := 1; i < len(strIn); i++ {
			if isReaction(strIn[i-1], strIn[i]) {
				strIn = strings.Replace(strIn, strIn[i-1:i+1], "", 1)
				reacted = true
				i += 1
			}
		}
		if !reacted {
			return strIn
		}
	}

}

func isReaction(first, second byte) bool {
	return math.Abs(float64(first)-float64(second)) == float64(32)
}
