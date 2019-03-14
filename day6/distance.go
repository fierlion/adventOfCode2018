package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Point2d struct {
	X int
	Y int
}

func (p Point2d) String() string {
	return fmt.Sprintf("{%d, %d}", p.X, p.Y)
}

// assuming that Int16 is large enough for our input set
const (
	MinInt = -1 << 15
	MaxInt = 1<<15 - 1
)

func main() {
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var points []*Point2d
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		parsedPoint := parsePoint(curString)
		points = append(points, parsedPoint)
	}

	minX := MaxInt
	minY := MaxInt
	maxX := MinInt
	maxY := MinInt
	// find minX, minY, maxX, maxY
	for _, point := range points {
		if point.X <= minX {
			minX = point.X
		} else if point.X >= maxX {
			maxX = point.X
		}
		if point.Y <= minY {
			minY = point.Y
		} else if point.Y >= maxY {
			maxY = point.Y
		}
	}
	fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", minX, minY, maxX, maxY)
}

func parsePoint(rawPoint string) *Point2d {
	rawStrings := strings.Split(rawPoint, ", ")
	convX, _ := strconv.Atoi(rawStrings[0])
	convY, _ := strconv.Atoi(rawStrings[1])
	return &Point2d{
		X: convX,
		Y: convY,
	}
}
