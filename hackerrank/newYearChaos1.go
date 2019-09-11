package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
    swapTotal := int32(0)
    // initialize highest to len array
    highest := int32(len(q))
    for highest > 0 {
        // search for highest value in array
        highestPos := findPosition(highest, q)
        // if not in correct position:
        onlyTwoSwaps := 0
        for ((highest - 1) != highestPos) {
            if onlyTwoSwaps > 1 {
                fmt.Printf("Too chaotic\n")
                return
            }
            q[highestPos+1], q[highestPos] = q[highestPos], q[highestPos+1]
            highestPos += 1
            swapTotal += 1
            onlyTwoSwaps += 1
        }
        onlyTwoSwaps = 0
        //fmt.Printf("%v\n", q)
        highest -= 1
    }
    fmt.Printf("%v\n", swapTotal)
}

func swapToPosition(start, end int32, []int32) []int32 {
    
}

func findPosition(next int32, q []int32) int32 {
    for pos, val := range q {
        if next == val {
            return int32(pos)
        }
    }
    return int32(-1)
}



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(readLine(reader), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

