package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var strings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		strings = append(strings, curString)
	}

	// n x m runtime
	for i := 0; i < len(strings[0]); i++ {
		result := dupeStringsWithoutIndex(i, strings)
		fmt.Printf("%s\n", result)
	}
}

func dupeStringsWithoutIndex(cutIdx int, strings []string) string {
	resultSet := make(map[string]bool)
	for _, curString := range strings {
		// create new string minus index
		start := curString[:cutIdx]
		end := curString[cutIdx+1:]
		testString := start + end
		if _, ok := resultSet[testString]; ok {
			return testString
		} else {
			resultSet[testString] = true
		}
	}
	return ""
}
