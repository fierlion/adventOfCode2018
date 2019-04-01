package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

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
	pointSet := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		parsedPoint := parsePoint(curString, currentId)
		points = append(points, parsedPoint)
		pointStr := pointToHashableString(parsedPoint)
		pointSet[pointStr] = currentId
		currentId += 1
	}

	// find boundaries of x and y
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
	minX := maxX
	minY := maxY
	for _, point := range points {
		if point.X <= minX {
			minX = point.X
		}
		if point.Y <= minY {
			minY = point.Y
		}
	}

	result := map[int]int{}

	// get owned
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			thisResult := bfsFindOwner(&Point2d{-1, x, y}, pointSet, &Point2d{-1, minX, minY}, &Point2d{-1, maxX, maxY})
			if count, ok := result[thisResult]; ok {
				result[thisResult] = count + 1
			} else {
				result[thisResult] = 1
			}
			fmt.Printf("%v\n", result)
		}
	}
	fmt.Printf("counts: %v\n", result)
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

func bfsFindOwner(start *Point2d, goalPoints map[string]int, minPoint *Point2d, maxPoint *Point2d) int {
	queue := list.New()
	visited := map[string]bool{}
	queue.PushBack(start)

	for queue.Len() > 0 {
		// pop front
		curPoint := queue.Front()
		queue.Remove(curPoint)
		curPointValue := curPoint.Value.(*Point2d)
		curPointStr := pointToHashableString(curPointValue)
		if _, ok := visited[curPointStr]; ok {
			continue
		}
		if val, ok := goalPoints[curPointStr]; ok {
			return val
		}
		visited[curPointStr] = true
		neighbors := getNeighbors(curPointValue, minPoint, maxPoint)
		for _, neighbor := range neighbors {
			queue.PushBack(neighbor)
		}
	}
	return -1 // shouldn't reach this.
}

func pointToHashableString(point *Point2d) string {
	parsedX := strconv.Itoa(point.X)
	parsedY := strconv.Itoa(point.Y)
	return parsedX + "|" + parsedY
}

func hashStringToPoint(pointStr string) *Point2d {
	splitStr := strings.Split(pointStr, "|")
	splitX, _ := strconv.Atoi(splitStr[0])
	splitY, _ := strconv.Atoi(splitStr[1])
	return &Point2d{
		Id: -1,
		X:  splitX,
		Y:  splitY,
	}
}

func getNeighbors(start *Point2d, minPoint *Point2d, maxPoint *Point2d) []*Point2d {
	var neighbors []*Point2d
	if start.X-1 >= minPoint.X {
		neighbors = append(neighbors, &Point2d{
			Id: -1,
			X:  (start.X - 1),
			Y:  start.Y,
		})
	}
	if start.X+1 <= maxPoint.X {
		neighbors = append(neighbors, &Point2d{
			Id: -1,
			X:  (start.X + 1),
			Y:  start.Y,
		})
	}
	if start.Y-1 >= minPoint.Y {
		neighbors = append(neighbors, &Point2d{
			Id: -1,
			X:  start.X,
			Y:  (start.Y - 1),
		})
	}
	if start.Y+1 <= maxPoint.Y {
		neighbors = append(neighbors, &Point2d{
			Id: -1,
			X:  start.X,
			Y:  (start.Y + 1),
		})
	}
	return neighbors
}
