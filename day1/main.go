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
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		curInt, _ := strconv.Atoi(curString)
		total += curInt
	}
	fmt.Printf("total: %d\n", total)
}
