package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Point2d struct {
	Id int
	X  int
	Y  int
}

func (p Point2d) String() string {
	return fmt.Sprintf("{%d, %d, %d}", p.Id, p.X, p.Y)
}

func main() {
	// parse input into array of Point2d
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	currentId := 0
	var points []*Point2d
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		parsedPoint := parsePoint(curString, currentId)
		points = append(points, parsedPoint)
		currentId += 1
	}
	maxX := 0
	maxY := 0
	for _, point := range points {
		if point.X >= maxX {
			maxX = point.X
		}
		if point.Y >= maxY {
			maxY = point.Y
		}
	}

	resultRegion := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			resultSum := getDistanceFromAllPoints(Point2d{-1, x, y}, points)
			if resultSum < 10000 {
				resultRegion += 1
			}
		}
	}
	fmt.Printf("result region size: %d\n", resultRegion)
}

func getDistanceFromAllPoints(start Point2d, points []*Point2d) int {
	resultSum := 0
	for _, point := range points {
		resultSum += point.taxiDistanceFrom(start)
	}
	return resultSum
}

func (p Point2d) taxiDistanceFrom(q Point2d) int {
	first := p.X - q.X
	if first < 0 {
		first = first * -1
	}
	second := p.Y - q.Y
	if second < 0 {
		second = second * -1
	}
	return first + second
}

func parsePoint(rawPoint string, currentId int) *Point2d {
	rawStrings := strings.Split(rawPoint, ", ")
	convX, _ := strconv.Atoi(rawStrings[0])
	convY, _ := strconv.Atoi(rawStrings[1])
	return &Point2d{
		Id: currentId,
		X:  convX,
		Y:  convY,
	}
}

func isInfinite(point *Point2d, minX int, maxX int, minY int, maxY int) bool {
	return (point.X == minX || point.X == maxX || point.Y == minY || point.Y == maxY)
}
