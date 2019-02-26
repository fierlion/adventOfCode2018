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
	month       int
	day         int
	hour        int
	minute      int
	guardId     int
	asleepAwake bool
}

func main() {
	inputFile, _ := filepath.Abs("./testInputSorted")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var entries []*guardEntry

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

}

// entryString is format: "[1518-03-05 23:57] Guard #2963 begins shift"
func parseEntry(entryString string) *guardEntry {
	splitString := strings.Split(entryString, " ")
	rawDate, rawTime, rawConscious, rawId := splitString[0][1:], splitString[1], splitString[2], splitString[3]
	thisId := 0
	thisAsleepAwake := true
	if rawConscious == "Guard" {
		newId, _ := strconv.Atoi(rawId[1:])
		thisId = newId
	} else {
		thisAsleepAwake = (rawConscious == "wakes")
	}

	// get date and time
	splitDate := strings.Split(rawDate, "-")
	thisMonth, _ := strconv.Atoi(splitDate[1])
	thisDay, _ := strconv.Atoi(splitDate[2])
	splitTime := strings.Split(rawTime, ":")
	thisHour, _ := strconv.Atoi(splitTime[0])
	thisMinute, _ := strconv.Atoi(splitTime[1][:2])

	return &guardEntry{
		month:       thisMonth,
		day:         thisDay,
		hour:        thisHour,
		minute:      thisMinute,
		guardId:     thisId,
		asleepAwake: thisAsleepAwake,
	}
}
