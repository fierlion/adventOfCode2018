package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type guardEntry struct {
	month   int
	day     int
	hour    int
	minute  int
	guardId int
	isAwake bool
}

const SHIFT_SIZE = 24 * 60

func main() {
	inputFile, _ := filepath.Abs("./testInputSorted")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var entries []*guardEntry
	var sleepResults = make(map[int][SHIFT_SIZE]int)

	scanner := bufio.NewScanner(file)
	currentGuard := 0
	for scanner.Scan() {
		curString := scanner.Text()
		entry := parseEntry(curString)
		if entry.guardId == 0 {
			entry.guardId = currentGuard
		} else {
			currentGuard = entry.guardId
		}
		entries = append(entries, entry)
	}

	// set sleep bits for each guard.
	// assumes guards will only have shifts between 00:00 and 23:59
	// -1 means guard starts shift
	currentGuardFallsAsleep := -1
	for _, entry := range entries {
		currentTime := (entry.hour * 60) + entry.minute
		if !entry.isAwake {
			currentGuardFallsAsleep = currentTime
		} else if currentGuardFallsAsleep >= 0 {
			if sleepResult, ok := sleepResults[entry.guardId]; ok {
				for i := currentGuardFallsAsleep; i < currentTime; i++ {
					sleepResult[i] += 1
				}
				sleepResults[entry.guardId] = sleepResult
			} else {
				var thisResult [SHIFT_SIZE]int
				for i := currentGuardFallsAsleep; i < currentTime; i++ {
					thisResult[i] += 1
				}
				sleepResults[entry.guardId] = thisResult
			}
			currentGuardFallsAsleep = -1
		}
	}
	for id, result := range sleepResults {
		thisResult, thisIndex := findMaxInt(result)
		fmt.Printf("guard %d: %d, %d\n", id, thisResult, thisIndex)
		fmt.Printf("array %v\n", result)
	}
}

// entryString is format: "[1518-03-05 23:57] Guard #2963 begins shift"
func parseEntry(entryString string) *guardEntry {
	splitString := strings.Split(entryString, " ")
	rawDate, rawTime, rawConscious, rawId := splitString[0][1:], splitString[1], splitString[2], splitString[3]
	thisId := 0
	thisIsAwake := true
	if rawConscious == "Guard" {
		newId, _ := strconv.Atoi(rawId[1:])
		thisId = newId
	} else {
		thisIsAwake = (rawConscious == "wakes")
	}

	// get date and time
	splitDate := strings.Split(rawDate, "-")
	thisMonth, _ := strconv.Atoi(splitDate[1])
	thisDay, _ := strconv.Atoi(splitDate[2])
	splitTime := strings.Split(rawTime, ":")
	thisHour, _ := strconv.Atoi(splitTime[0])
	thisMinute, _ := strconv.Atoi(splitTime[1][:2])

	return &guardEntry{
		month:   thisMonth,
		day:     thisDay,
		hour:    thisHour,
		minute:  thisMinute,
		guardId: thisId,
		isAwake: thisIsAwake,
	}
}

func findMaxInt(sleepResult [SHIFT_SIZE]int) (int, int) {
	maxValue := 0
	maxIndex := -1
	for idx, value := range sleepResult {
		if value >= maxValue {
			maxValue = value
			maxIndex = idx
		}
	}
	return maxValue, maxIndex
}
