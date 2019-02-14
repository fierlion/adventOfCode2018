package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	total := 0
	freqs := map[int]bool{}
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// get input array
	var intArray []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		curInt, _ := strconv.Atoi(curString)
		intArray = append(intArray, curInt)
	}

	// hack--append it onto itself for repetition
	for i := 0; i < 10; i++ {
		intArray = append(intArray, intArray...)
	}

	// hope that the first duplicate frequency is within the 10 repetitions
	for _, thisInt := range intArray {
		freqs[total] = true
		total += thisInt
		if _, ok := freqs[total]; ok {
			fmt.Printf("%d\n", total)
			break
		}
	}
}
