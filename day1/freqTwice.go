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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		freqs[total] = true
		curString := scanner.Text()
		curInt, _ := strconv.Atoi(curString)
		total += curInt
		if _, ok := freqs[total]; ok {
			fmt.Printf("%d\n", total)
		}
	}
	//test
	fmt.Printf("total: %d\n", total)
}
