package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type elfClaim struct {
	claimId  int
	fromLeft int
	fromTop  int
	width    int
	height   int
}

const CLOTH_SIZE = 1000

func main() {
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var claims []*elfClaim

	// initialized with default int zero value
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		claim := parseClaim(curString)
		claims = append(claims, claim)
	}
	cloth := make([][]int, CLOTH_SIZE)
	for i := range cloth {
		cloth[i] = make([]int, CLOTH_SIZE)
	}
	for _, claim := range claims {
		markClaim(&cloth, claim)
	}
	result := countDuplicateClaims(&cloth)
	fmt.Printf("%d\n", result)
}

// claimString is format: "#1 @ 704,926: 5x4"
func parseClaim(claimString string) *elfClaim {
	splitString := strings.Split(claimString, " ")
	rawId, rawPad, rawSize := splitString[0][1:], splitString[2], splitString[3]
	thisId, _ := strconv.Atoi(rawId)
	// get padding
	splitPad := strings.Split(rawPad[:len(rawPad)-1], ",")
	thisLeft, _ := strconv.Atoi(splitPad[0])
	thisTop, _ := strconv.Atoi(splitPad[1])
	// get dimensions
	splitDim := strings.Split(rawSize, "x")
	thisWidth, _ := strconv.Atoi(splitDim[0])
	thisHeight, _ := strconv.Atoi(splitDim[1])
	return &elfClaim{
		claimId:  thisId,
		fromLeft: thisLeft,
		fromTop:  thisTop,
		width:    thisWidth,
		height:   thisHeight,
	}
}

// this might be nice if we parallelized it with some goroutines
func markClaim(cloth *[][]int, claim *elfClaim) {
	// check that fromLeft + width not > 1000
	// check that fromTop + height not > 1000
	horizEnd := claim.fromLeft + claim.width
	vertEnd := claim.fromTop + claim.height
	for horiz := claim.fromLeft; horiz < horizEnd; horiz++ {
		for vert := claim.fromTop; vert < vertEnd; vert++ {
			(*cloth)[horiz][vert] += 1
		}
	}
}

// we can probably wrap this up in the markClaim step.
// keeping it simple for now.
func countDuplicateClaims(cloth *[][]int) int {
	result := 0
	for horiz := 0; horiz < CLOTH_SIZE; horiz++ {
		for vert := 0; vert < CLOTH_SIZE; vert++ {
			if (*cloth)[horiz][vert] > 1 {
				result += 1
			}
		}
	}
	return result
}
